package sqlconnect

import (
	"database/sql"
	"fmt"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
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
		return nil, utils.ErrorHandler(err, "Error connecting to database")
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
		return nil, utils.ErrorHandler(err, "Error retrieving teachers from database")
	}

	defer rows.Close()

	// teachers := make([]models.Teacher, 0)
	for rows.Next() {
		var teacher models.Teacher
		if err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject); err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving teachers from database")
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

// GetTeacherByID retrieves a single teacher by their ID.
func GetTeacherByID(id int) (models.Teacher, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
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
			return models.Teacher{}, utils.ErrorHandler(err, fmt.Sprintf("Teacher with ID: %d not found in database", id))
		}
		return models.Teacher{}, utils.ErrorHandler(err, "Error retrieving teacher by ID from database")
	}
	return teacher, nil
}

// CreateTeachers adds new teachers to the database.
func CreateTeachers(newTeachers []models.Teacher) ([]models.Teacher, error) {
	db, err := ConnectDb()
	if err != nil {
		return []models.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error inserting teacher data into database")
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting teacher data into database")
		}

		// Get the last inserted ID
		lastId, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting teacher data into database")
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
		return models.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
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
			return models.Teacher{}, utils.ErrorHandler(err, fmt.Sprintf("Teacher with ID: %d not found in database", id))
		}
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher in the database")
	}

	// Check if there are any changes before updating
	if updatedTeacher.FirstName == teacherToUpdate.FirstName &&
		updatedTeacher.LastName == teacherToUpdate.LastName &&
		updatedTeacher.Email == teacherToUpdate.Email &&
		updatedTeacher.Class == teacherToUpdate.Class &&
		updatedTeacher.Subject == teacherToUpdate.Subject {

		return models.Teacher{}, fmt.Errorf("no changes detected in the teacher's details")
	}

	const updateTeacherQuery = `
		UPDATE teachers
		SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ?
		WHERE id = ?`

	updatedTeacher.ID = teacherToUpdate.ID
	_, err = db.Exec(updateTeacherQuery, updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher in the database")
	}
	return updatedTeacher, nil
}

// PatchTeachersInDb performs partial updates on multiple teachers in the database.
func PatchTeachersInDb(updatedFields []map[string]interface{}) ([]models.Teacher, error) {

	var teachersFromDB []models.Teacher
	db, err := ConnectDb()
	if err != nil {
		return teachersFromDB, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Validate all fields before starting the transaction
	for _, teacherUpdate := range updatedFields {
		idFloat, ok := teacherUpdate["id"].(float64)
		if !ok {
			return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
		}
		id := int(idFloat)

		teacherToUpdate, err := getTeacherByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return teachersFromDB, utils.ErrorHandler(err, fmt.Sprintf("Teacher with ID: %d not found in database", id))
			}
			return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
		}

		validFields := buildValidFieldsMap(teacherToUpdate)
		if err := validateUpdateFields(validFields, teacherUpdate); err != nil {
			return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
	}

	for _, teacherUpdate := range updatedFields {
		id := int(teacherUpdate["id"].(float64))
		teacherToUpdate, err := getTeacherByID(tx, id)
		if err != nil {
			tx.Rollback()
			return teachersFromDB, utils.ErrorHandler(err, fmt.Sprintf("Teacher with ID: %d not found in database", id))
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
			return teachersFromDB, fmt.Errorf("no valid fields provided for update")
		}
		updateArgs = append(updateArgs, teacherToUpdate.ID)
		updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

		_, err = tx.Exec(updateTeacherQuery, updateArgs...)
		if err != nil {
			tx.Rollback()
			return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
		}

		teachersFromDB = append(teachersFromDB, teacherToUpdate)
	}

	if err := tx.Commit(); err != nil {
		return teachersFromDB, utils.ErrorHandler(err, "Error updating teacher data into database")
	}

	return teachersFromDB, nil
}

// PatchTeacherByID performs a partial update on a single teacher by their ID.
func PatchTeacherByID(id int, updatedFields map[string]interface{}) (models.Teacher, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Get existing teacher by id using helper
	teacherToUpdate, err := getTeacherByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Teacher{}, utils.ErrorHandler(err, fmt.Sprintf("Teacher with ID: %d not found in database", id))
		}
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher data into database")
	}

	// Build valid fields map using helper
	validFields := buildValidFieldsMap(teacherToUpdate)

	// Validate fields to update using helper
	if err := validateUpdateFields(validFields, updatedFields); err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher data into database")
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
		return models.Teacher{}, fmt.Errorf("no valid fields provided for update")
	}

	updateArgs = append(updateArgs, teacherToUpdate.ID)
	updateTeacherQuery := fmt.Sprintf("UPDATE teachers SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateTeacherQuery, updateArgs...)
	if err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher data into database")
	}

	return teacherToUpdate, nil
}

// DeleteTeacherByID deletes a single teacher by their ID.
func DeleteTeacherByID(id int) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Delete the teacher
	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting teacher from database")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting teacher from database")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("teacher with ID %d not found", id)
	}
	return nil
}

// DeleteTeachersInDB deletes multiple teachers by their IDs and
// returns the list of deleted IDs.
func DeleteTeachersInDB(IDs []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error deleting teachers from database")
	}

	stmt, err := tx.Prepare("DELETE FROM teachers WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "Error deleting teachers from database")
	}
	defer stmt.Close()

	// type myInt int
	deletedIDs := []int{}

	for _, id := range IDs {
		// Delete the teacher
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("Failed to delete teacher with ID %d: %v", id, err))
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "Error deleting teachers from database")
		}

		if rowsAffected == 0 {
			tx.Rollback()
			return nil, fmt.Errorf("teacher with ID %d not found", id)
		}

		deletedIDs = append(deletedIDs, id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error deleting teachers from database")
	}

	if len(deletedIDs) == 0 {
		return nil, fmt.Errorf("no teachers found")
	}
	return deletedIDs, nil
}
