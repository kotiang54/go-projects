package sqlconnect

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
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

// Helper functions
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

// CreateTeachers adds new teachers to the database.
func CreateTeachers(newTeachers []models.Teacher) ([]models.Teacher, error) {
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

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to prepare SQL statement: %v", err), http.StatusInternalServerError)
		return nil, err
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			// http.Error(w, fmt.Sprintf("Failed to insert teacher: %v", err), http.StatusInternalServerError)
			return nil, err
		}

		// Get the last inserted ID
		lastId, err := res.LastInsertId()
		if err != nil {
			// http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
			return nil, err
		}

		newTeacher.ID = int(lastId)
		addedTeachers[i] = newTeacher
	}

	return addedTeachers, nil
}

// UpdateTeacherByID updates an existing teacher's details by their ID.
func UpdateTeacherByID(id int, updatedTeacher models.Teacher) (models.Teacher, error) {
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

	// get the existing teacher from database
	query := "SELECT * FROM teachers WHERE id = ?"
	var teacherToUpdate models.Teacher
	err = db.QueryRow(query, id).
		Scan(&teacherToUpdate.ID, &teacherToUpdate.FirstName, &teacherToUpdate.LastName, &teacherToUpdate.Email, &teacherToUpdate.Class, &teacherToUpdate.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, err
		}
		// http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return models.Teacher{}, err
	}

	// Check if there are any changes before updating
	if updatedTeacher.FirstName == teacherToUpdate.FirstName &&
		updatedTeacher.LastName == teacherToUpdate.LastName &&
		updatedTeacher.Email == teacherToUpdate.Email &&
		updatedTeacher.Class == teacherToUpdate.Class &&
		updatedTeacher.Subject == teacherToUpdate.Subject {

		// http.Error(w, "No changes detected in the teacher's details", http.StatusBadRequest)
		return models.Teacher{}, err
	}

	const updateTeacherQuery = `
		UPDATE teachers
		SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ?
		WHERE id = ?`

	updatedTeacher.ID = teacherToUpdate.ID
	_, err = db.Exec(updateTeacherQuery, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	return updatedTeacher, nil
}

// PatchTeachersInDb performs partial updates on multiple teachers in the database.
func PatchTeachersInDb(updatedFields []map[string]interface{}) ([]models.Teacher, error) {

	var teachersFromDB []models.Teacher
	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return teachersFromDB, err
	}
	defer db.Close()

	// Validate all fields before starting the transaction
	for _, teacherUpdate := range updatedFields {
		idFloat, ok := teacherUpdate["id"].(float64)
		if !ok {
			// http.Error(w, "Each update must include a valid 'id' field", http.StatusBadRequest)
			return teachersFromDB, err
		}
		id := int(idFloat)

		teacherToUpdate, err := getTeacherByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				// http.Error(w, fmt.Sprintf("Teacher not found (id: %d)", id), http.StatusNotFound)
				return teachersFromDB, err
			}
			// http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
			return teachersFromDB, err
		}

		validFields := buildValidFieldsMap(teacherToUpdate)
		if err := validateUpdateFields(validFields, teacherUpdate); err != nil {
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return teachersFromDB, err
		}
	}

	tx, err := db.Begin()
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return teachersFromDB, err
	}

	for _, teacherUpdate := range updatedFields {
		id := int(teacherUpdate["id"].(float64))
		teacherToUpdate, err := getTeacherByID(tx, id)
		if err != nil {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
			return teachersFromDB, err
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
			// http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
			return teachersFromDB, err
		}
		updateArgs = append(updateArgs, teacherToUpdate.ID)
		updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

		_, err = tx.Exec(updateTeacherQuery, updateArgs...)
		if err != nil {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
			return teachersFromDB, err
		}

		teachersFromDB = append(teachersFromDB, teacherToUpdate)
	}

	if err := tx.Commit(); err != nil {
		// http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return teachersFromDB, err
	}

	return teachersFromDB, nil
}

// PatchTeacherByID performs a partial update on a single teacher by their ID.
func PatchTeacherByID(id int, updatedFields map[string]interface{}) (models.Teacher, error) {

	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return models.Teacher{}, err
	}
	defer db.Close()

	// Get existing teacher by id using helper
	teacherToUpdate, err := getTeacherByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "Teacher not found", http.StatusNotFound)
			return models.Teacher{}, err
		}
		// http.Error(w, fmt.Sprintf("Database query error: %v", err), http.StatusInternalServerError)
		return models.Teacher{}, err
	}

	// Build valid fields map using helper
	validFields := buildValidFieldsMap(teacherToUpdate)

	// Validate fields to update using helper
	if err := validateUpdateFields(validFields, updatedFields); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return models.Teacher{}, err
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
		// http.Error(w, "No valid fields provided for update", http.StatusBadRequest)
		return models.Teacher{}, fmt.Errorf("no valid fields provided for update")
	}

	updateArgs = append(updateArgs, teacherToUpdate.ID)
	updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateTeacherQuery, updateArgs...)
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to update teacher: %v", err), http.StatusInternalServerError)
		return models.Teacher{}, err
	}

	return teacherToUpdate, nil
}

func DeleteTeacherByID(id int) error {
	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	// Delete the teacher
	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to delete teacher: %v", err), http.StatusInternalServerError)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to retrieve affected rows: %v", err), http.StatusInternalServerError)
		return err
	}

	if rowsAffected == 0 {
		// http.Error(w, "Teacher not found", http.StatusNotFound)
		return fmt.Errorf("teacher with ID %d not found", id)
	}
	return nil
}

// DeleteTeachersInDB deletes multiple teachers by their IDs and
// returns the list of deleted IDs.
func DeleteTeachersInDB(IDs []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {
		// http.Error(w, "Database connection error", http.StatusInternalServerError)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to begin transaction: %v", err), http.StatusInternalServerError)
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		tx.Rollback()
		// http.Error(w, fmt.Sprintf("Failed to prepare statement: %v", err), http.StatusInternalServerError)
		return nil, err
	}
	defer stmt.Close()

	// type myInt int
	deletedIDs := []int{}

	for _, id := range IDs {
		// Delete the teacher
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Failed to delete teacher with ID %d: %v", id, err), http.StatusInternalServerError)
			return nil, err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Failed to retrieve affected rows: %v", err), http.StatusInternalServerError)
			return nil, err
		}

		if rowsAffected == 0 {
			tx.Rollback()
			// http.Error(w, fmt.Sprintf("Teacher with ID %d not found", id), http.StatusNotFound)
			return nil, fmt.Errorf("teacher with ID %d not found", id)
		}

		deletedIDs = append(deletedIDs, id)
	}

	err = tx.Commit()
	if err != nil {
		// http.Error(w, fmt.Sprintf("Failed to commit transaction: %v", err), http.StatusInternalServerError)
		return nil, err
	}

	if len(deletedIDs) == 0 {
		// http.Error(w, "No teachers found", http.StatusNotFound)
		return nil, fmt.Errorf("no teachers found")
	}
	return deletedIDs, nil
}
