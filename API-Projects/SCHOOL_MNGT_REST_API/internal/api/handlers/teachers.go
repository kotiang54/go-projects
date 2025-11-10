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

// Add this interface for query compatibility
type queryer interface {
	QueryRow(query string, args ...interface{}) *sql.Row
}

// type dbHandler struct {
// 	db *sql.DB
// }

// func (h *dbHandler) QueryRow(query string, args ...interface{}) *sql.Row {
// 	return h.db.QueryRow(query, args...)
// }

// GetTeachersHandler handles GET requests to fetch teachers
func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

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

// GetOneTeacherHandler handles GET requests to fetch a specific teacher
func GetOneTeacherHandler(w http.ResponseWriter, r *http.Request) {

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

	teacherIDStr := r.PathValue("id")

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

// CreateTeachersHandler handles the creation of new teachers
func CreateTeachersHandler(w http.ResponseWriter, r *http.Request) {

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

// UpdateTeachersHandler handles updating an existing teacher
func UpdateTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// get teachers id from path
	idStr := r.PathValue("id")
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
	var teacherToUpdate models.Teacher
	err = db.QueryRow(query, id).
		Scan(&teacherToUpdate.ID, &teacherToUpdate.FirstName, &teacherToUpdate.LastName, &teacherToUpdate.Email, &teacherToUpdate.Class, &teacherToUpdate.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Check if there are any changes before updating
	if updatedTeacher.FirstName == teacherToUpdate.FirstName &&
		updatedTeacher.LastName == teacherToUpdate.LastName &&
		updatedTeacher.Email == teacherToUpdate.Email &&
		updatedTeacher.Class == teacherToUpdate.Class &&
		updatedTeacher.Subject == teacherToUpdate.Subject {

		http.Error(w, "No changes detected in the teacher's details", http.StatusBadRequest)
		return
	}

	const updateTeacherQuery = `
		UPDATE teachers
		SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ?
		WHERE id = ?`

	updatedTeacher.ID = teacherToUpdate.ID
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

// PatchTeachersHandler handles PATCH requests to partially update teachers records
// PATCH /teachers/
func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var updatedFields []map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate all fields before starting the transaction
	for _, teacherUpdate := range updatedFields {
		idFloat, ok := teacherUpdate["id"].(float64)
		if !ok {
			http.Error(w, "Each update must include a valid 'id' field", http.StatusBadRequest)
			return
		}
		id := int(idFloat)

		teacherToUpdate, err := getTeacherByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, fmt.Sprintf("Teacher not found (id: %d)", id), http.StatusNotFound)
				return
			}
			http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
			return
		}

		validFields := buildValidFieldsMap(teacherToUpdate)
		if err := validateUpdateFields(validFields, teacherUpdate); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return
	}

	var teachersFromDB []models.Teacher

	for _, teacherUpdate := range updatedFields {
		id := int(teacherUpdate["id"].(float64))
		teacherToUpdate, err := getTeacherByID(tx, id)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
			return
		}

		validFields := buildValidFieldsMap(teacherToUpdate)
		applyUpdateToStruct(&teacherToUpdate, validFields, teacherUpdate)

		var updateFields []string
		var updateArgs []interface{}
		for key, value := range teacherUpdate {
			if key == "id" {
				continue
			}
			updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
			updateArgs = append(updateArgs, value)
		}
		if len(updateFields) == 0 {
			tx.Rollback()
			http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
			return
		}
		updateArgs = append(updateArgs, teacherToUpdate.ID)
		updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

		_, err = tx.Exec(updateTeacherQuery, updateArgs...)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
			return
		}

		teachersFromDB = append(teachersFromDB, teacherToUpdate)
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachersFromDB)
}

// PatchOneTeacherHandler handles PATCH requests to partially update a teacher records
// PATCH /teachers/{id}
func PatchOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Get the teacher id from the path
	idStr := r.PathValue("id")
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

	// Get existing teacher by id using helper
	teacherToUpdate, err := getTeacherByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Build valid fields map using helper
	validFields := buildValidFieldsMap(teacherToUpdate)

	// Validate fields to update using helper
	if err := validateUpdateFields(validFields, updatedFields); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Apply updates to struct using helper
	applyUpdateToStruct(&teacherToUpdate, validFields, updatedFields)

	// Build update query and args
	var updateFields []string
	var updateArgs []interface{}
	for key, value := range updatedFields {
		if key == "id" {
			continue
		}
		updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
		updateArgs = append(updateArgs, value)
	}
	if len(updateFields) == 0 {
		http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
		return
	}

	updateArgs = append(updateArgs, teacherToUpdate.ID)
	updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateTeacherQuery, updateArgs...)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the updated teacher
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacherToUpdate)
}

// DeleteOneTeacherHandler handles DELETE requests to remove a teacher record
func DeleteOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Get the teachers Id from the path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Teacher ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Connect to database
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Delete the teacher
	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete teacher: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to retrieve affected rows: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	// w.WriteHeader(http.StatusNoContent)
	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: fmt.Sprintf("Teacher with ID %d deleted successfully", id),
	}

	json.NewEncoder(w).Encode(response)
}

// DeleteTeachersHandler handles DELETE requests to remove teachers record
func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var IDs []int
	err = json.NewDecoder(r.Body).Decode(&IDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		tx.Rollback()
		http.Error(w, fmt.Sprintf("Failed to prepare statement: %v", err), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// type myInt int
	deletedIDs := []int{}

	for _, id := range IDs {
		// Delete the teacher
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Failed to delete teacher with ID %d: %v", id, err), http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Failed to retrieve affected rows: %v", err), http.StatusInternalServerError)
			return
		}

		if rowsAffected == 0 {
			tx.Rollback()
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}

		deletedIDs = append(deletedIDs, id)
	}

	err = tx.Commit()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return
	}

	if len(deletedIDs) == 0 {
		http.Error(w, "No teachers found", http.StatusNotFound)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "Teachers successfully deleted",
		DeletedIDs: deletedIDs,
	}

	json.NewEncoder(w).Encode(response)
}

// Helper functions
// getTeacherByID retrieves a teacher by ID from the database
func getTeacherByID(db queryer, id int) (models.Teacher, error) {
	var teacher models.Teacher
	query := "SELECT * FROM teachers WHERE id = ?"
	err := db.QueryRow(query, id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	return teacher, err
}

// buildValidFieldsMap builds a map of valid JSON field names to struct field indices
func buildValidFieldsMap(teacher models.Teacher) map[string]int {
	teacherType := reflect.TypeOf(teacher)
	validFields := make(map[string]int)
	for i := 0; i < teacherType.NumField(); i++ {
		jsonTag := strings.Split(teacherType.Field(i).Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			validFields[jsonTag] = i
		}
	}
	return validFields
}

// validateUpdateFields checks if the fields in the update map are valid and of correct type
func validateUpdateFields(validFields map[string]int, update map[string]interface{}) error {
	for key, value := range update {
		if key == "id" {
			continue
		}
		fieldIdx, ok := validFields[key]
		if !ok {
			return fmt.Errorf("invalid field: %s", key)
		}
		fieldType := reflect.TypeOf(models.Teacher{}).Field(fieldIdx).Type
		val := reflect.ValueOf(value)
		if !val.Type().ConvertibleTo(fieldType) {
			return fmt.Errorf("type mismatch for field: %s", key)
		}
	}
	return nil
}

func applyUpdateToStruct(teacher *models.Teacher, validFields map[string]int, update map[string]interface{}) {
	for key, value := range update {
		if key == "id" {
			continue
		}
		fieldIdx := validFields[key]
		fieldVal := reflect.ValueOf(teacher).Elem().Field(fieldIdx)
		val := reflect.ValueOf(value)
		fieldVal.Set(val.Convert(fieldVal.Type()))
	}
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
