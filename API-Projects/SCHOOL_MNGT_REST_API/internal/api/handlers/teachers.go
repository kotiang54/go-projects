package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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
		// Handle PATCH request to partially update teacher records
		patchTeachersHandler(w, r)

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
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

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

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, fmt.Sprintf("Teacher not found with ID: %d", id), http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
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
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

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

// updateTeachersHandler handles updating an existing teacher
func updateTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// get teachers id from path
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Teacher ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// create updated teacher variable from request body
	var updatedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the updatedTeacher fields
	if updatedTeacher.FirstName == "" || updatedTeacher.LastName == "" || updatedTeacher.Email == "" || updatedTeacher.Class == "" || updatedTeacher.Subject == "" {
		http.Error(w, "All fields (first_name, last_name, email, class, subject) are required", http.StatusBadRequest)
		return
	}

	// update teacher in database
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// get the existing teacher from database
	query := "SELECT * FROM teachers WHERE id = ?"
	var existingTeacher models.Teacher
	err = db.QueryRow(query, id).
		Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if there are any changes before updating
	if updatedTeacher.FirstName == existingTeacher.FirstName &&
		updatedTeacher.LastName == existingTeacher.LastName &&
		updatedTeacher.Email == existingTeacher.Email &&
		updatedTeacher.Class == existingTeacher.Class &&
		updatedTeacher.Subject == existingTeacher.Subject {

		http.Error(w, "No changes detected in the teacher's details", http.StatusBadRequest)
		return
	}

	const updateTeacherQuery = `
		UPDATE teachers
		SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ?
		WHERE id = ?`

	updatedTeacher.ID = existingTeacher.ID
	_, err = db.Exec(updateTeacherQuery, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
		return
	}

	// return updated teacher with status field
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string         `json:"status"`
		Data   models.Teacher `json:"data"`
	}{
		Status: "success",
		Data:   updatedTeacher,
	}

	json.NewEncoder(w).Encode(response)
}

// patchTeachersHandler handles PATCH requests to partially update teacher records
func patchTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Get the teacher id from the path
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Teacher ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Decode fields to update from request body
	var updatedFields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// Connect to database
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Get existing teacher by id
	var existingTeacher models.Teacher
	query := "SELECT * FROM teachers WHERE id = ?"
	err = db.QueryRow(query, id).
		Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Validate fields to update
	teacherType := reflect.TypeOf(existingTeacher)
	validFields := make(map[string]int)
	for i := 0; i < teacherType.NumField(); i++ {
		jsonTag := strings.Split(teacherType.Field(i).Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			validFields[jsonTag] = i
		}
	}

	// Build update query and args
	var updateFields []string
	var updateArgs []interface{}
	for key, value := range updatedFields {
		fieldIdx, ok := validFields[key]
		if !ok {
			http.Error(w, fmt.Sprintf("Invalid field: %s", key), http.StatusBadRequest)
			return
		}
		updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
		updateArgs = append(updateArgs, value)

		// Update the struct field using reflection
		fieldVal := reflect.ValueOf(&existingTeacher).Elem().Field(fieldIdx)
		val := reflect.ValueOf(value)
		if val.Type().ConvertibleTo(fieldVal.Type()) {
			fieldVal.Set(val.Convert(fieldVal.Type()))
		} else {
			http.Error(w, fmt.Sprintf("Type mismatch for field: %s", key), http.StatusBadRequest)
			return
		}
	}

	if len(updateFields) == 0 {
		http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
		return
	}

	updateArgs = append(updateArgs, existingTeacher.ID)
	updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateTeacherQuery, updateArgs...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the updated teacher
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}
