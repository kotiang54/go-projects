package sqlconnect

import (
	"database/sql"
	"fmt"
	"net/http"
	"school_management_api/internal/models"
	"strings"
)

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

// GetTeachersCollection retrieves a collection of teachers from the database
// with optional filtering and sorting.
func GetTeachersCollection(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return nil, err
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

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
		// http.Error(w, fmt.Sprintf("Database Query Error: %v", err), http.StatusInternalServerError)
		return nil, err
	}

	defer rows.Close()

	// teachers := make([]models.Teacher, 0)
	for rows.Next() {
		var teacher models.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject); err != nil {
			// http.Error(w, fmt.Sprintf("Database scan error: %v", err), http.StatusInternalServerError)
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

// GetTeacherByID retrieves a single teacher by their ID.
func GetTeacherByID(id int) (models.Teacher, error) {

	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	var teacher models.Teacher
	query := "SELECT * FROM teachers WHERE id = ?" // id, first_name, last_name, email, class, subject
	err = db.QueryRow(query, id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, fmt.Sprintf("Teacher not found with ID: %d", id), http.StatusNotFound)
			return models.Teacher{}, err
		}
		// http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return teacher, nil
}
