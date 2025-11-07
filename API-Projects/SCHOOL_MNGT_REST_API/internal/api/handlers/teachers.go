package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Path parameters e.g. /teachers/{id}
	// Query parameters e.g. /teachers/?key=value&query=value2&sortBy=email&sortOrder=ASC

	switch r.Method {
	case http.MethodGet:
		// Handle GET request to fetch all teachers
		getTeachersHandler(w, r)

	case http.MethodPost:
		// Handle POST request to create a new teacher
		createTeachersHandler(w, r)

	case http.MethodPut:
		// Handle PUT request to update an existing teacher
		updateTeachersHandler(w, r)

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		return
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Path parameters can be handled here if needed
	// e.g. teacherID := chi.URLParam(r, "id")

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	teacherIDStr := strings.TrimSuffix(path, "/")

	if teacherIDStr == "" {

		// Build the SQL query with filters
		query := "SELECT * FROM teachers WHERE 1=1" // * id, first_name, last_name, email, class, subject
		var args []interface{}

		// Add filters based on query parameters
		query, args = addFilters(r, query, args)

		// Example: /teachers/?sortby=last_name:asc&sortby=class:desc
		query += buildOrderByClause(r)

		// Execute the query
		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, fmt.Sprintf("Database Query Error: %v", err), http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		teacherList := make([]models.Teacher, 0)
		for rows.Next() {
			var teacher models.Teacher
			if err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject); err != nil {
				http.Error(w, fmt.Sprintf("Database scan error: %v", err), http.StatusInternalServerError)
				return
			}
			teacherList = append(teacherList, teacher)
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameters for specific teacher
	id, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return
	}

	var teacher models.Teacher
	query := "SELECT * FROM teachers WHERE id = ?" // id, first_name, last_name, email, class, subject
	err = db.QueryRow(query, id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func createTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to prepare SQL statement: %v", err), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert teacher: %v", err), http.StatusInternalServerError)
			return
		}

		// Get the last inserted ID
		lastId, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return
		}

		newTeacher.ID = int(lastId)
		addedTeachers[i] = newTeacher
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Response structure with status, count, and data
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}

// addFilters adds filtering conditions to the SQL query based on URL query parameters.
func addFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	// Handle Query parameters for filtering
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += fmt.Sprintf(" AND %s = ?", dbField)
			args = append(args, value)
		}
	}
	return query, args
}

// Extracted function for building ORDER BY clause from sortby query parameters
func buildOrderByClause(r *http.Request) string {
	sortParams := r.URL.Query()["sortby"]
	if len(sortParams) == 0 {
		return ""
	}
	orderBy := " ORDER BY "
	for i, param := range sortParams {
		parts := strings.Split(param, ":")
		if len(parts) == 2 {
			field := parts[0]
			order := strings.ToUpper(parts[1])
			if order != "ASC" && order != "DESC" {
				order = "ASC"
			}
			if i > 0 {
				orderBy += ", "
			}
			orderBy += fmt.Sprintf("%s %s", field, order)
		}
	}
	return orderBy
}

func updateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating a teacher
}
