package sqlconnect

import (
	"database/sql"
	"fmt"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
	"strings"
)

// =========== Helper functions ===================

// addTeachersFilter adds filtering conditions to the SQL query based on URL query parameters.
func addTeachersFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
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

// getTeacherByID retrieves a teacher by ID from the database
func getTeacherByID(db queryer, id int) (models.Teacher, error) {
	var teacher models.Teacher
	query := "SELECT * FROM teachers WHERE id = ?"
	err := db.QueryRow(query, id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	return teacher, err
}

// ================ Database Operations ===================

// GetTeachersCollection retrieves a collection of teachers from the database
// with optional filtering and sorting.
func GetTeachersInDb(teachers []models.Teacher, r *http.Request) ([]models.Teacher, error) {
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
	query, args = addTeachersFilter(r, query, args)

	// Example: /teachers/?sortby=last_name:asc&sortby=class:desc
	query += utils.BuildOrderByClause(r)

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

	// stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?, ?, ?, ?, ?)")
	stmt, err := db.Prepare(utils.GenerateInsertQuery(models.Teacher{}, "teachers"))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error inserting teacher data into database")
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		// res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		values := utils.GetStructValues(newTeacher)
		res, err := stmt.Exec(values...)
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
	values := append(utils.GetStructValues(updatedTeacher), teacherToUpdate.ID)
	_, err = db.Exec(updateTeacherQuery, values...)
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

		validFields := utils.BuildValidFieldsMap(teacherToUpdate)
		if err := utils.ValidateUpdateFields(models.Teacher{}, validFields, teacherUpdate); err != nil {
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

		validFields := utils.BuildValidFieldsMap(teacherToUpdate)
		utils.ApplyUpdateToStruct(&teacherToUpdate, validFields, teacherUpdate)

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
	validFields := utils.BuildValidFieldsMap(teacherToUpdate)

	// Validate fields to update using helper
	if err := utils.ValidateUpdateFields(models.Teacher{}, validFields, updatedFields); err != nil {
		return models.Teacher{}, utils.ErrorHandler(err, "Error updating teacher data into database")
	}

	// Apply updates to struct using helper
	utils.ApplyUpdateToStruct(&teacherToUpdate, validFields, updatedFields)

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

func GetStudentsByTeacherID(teacherId string) ([]models.Student, error) {
	var students []models.Student

	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	query := `SELECT * FROM	students WHERE class = (SELECT class FROM teachers WHERE id = ?)`
	rows, err := db.Query(query, teacherId)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving data from database")
	}
	defer rows.Close()

	for rows.Next() {
		var student models.Student
		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving data from database")
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving data from database")
	}
	return students, err
}
