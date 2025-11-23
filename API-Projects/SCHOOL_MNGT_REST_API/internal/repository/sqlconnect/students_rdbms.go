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

// addStudentsFilter adds filtering conditions to the SQL query based on URL query parameters.
func addStudentsFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	// Handle Query parameters for filtering
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
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

// getStudentByID retrieves a student by ID from the database
func getStudentByID(db queryer, id int) (models.Student, error) {
	var student models.Student
	query := "SELECT * FROM students WHERE id = ?"
	err := db.QueryRow(query, id).
		Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)
	return student, err
}

// ================ Database Operations ===================

// GetStudentsInDb retrieves a collection of students from the database
// with optional filtering and sorting.
func GetStudentsInDb(students []models.Student, r *http.Request) ([]models.Student, error) {
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
	query := "SELECT * FROM students WHERE 1=1" // * id, first_name, last_name, email, class
	var args []interface{}

	// Add filters based on query parameters
	query, args = addStudentsFilter(r, query, args)

	// Example: /students/?sortby=last_name:asc&sortby=class:desc
	query += utils.BuildOrderByClause(r)

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving students from database")
	}

	defer rows.Close()

	for rows.Next() {
		var student models.Student
		if err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class); err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving students from database")
		}
		students = append(students, student)
	}
	return students, nil
}

// GetStudentByID retrieves a single student by their ID.
func GetStudentByID(id int) (models.Student, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	var student models.Student
	query := "SELECT * FROM students WHERE id = ?" // id, first_name, last_name, email, class
	err = db.QueryRow(query, id).
		Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Class)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, utils.ErrorHandler(err, fmt.Sprintf("Student with ID: %d not found in database", id))
		}
		return models.Student{}, utils.ErrorHandler(err, "Error retrieving student by ID from database")
	}
	return student, nil
}

// CreateStudents adds new students to the database.
func CreateStudents(newStudents []models.Student) ([]models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return []models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// stmt, err := db.Prepare("INSERT INTO students (first_name, last_name, email, class) VALUES (?, ?, ?, ?)")
	stmt, err := db.Prepare(utils.GenerateInsertQuery(models.Student{}, "students"))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error inserting student data into database")
	}
	defer stmt.Close()

	addedStudents := make([]models.Student, len(newStudents))

	for i, newStudent := range newStudents {
		values := utils.GetStructValues(newStudent)
		res, err := stmt.Exec(values...)
		if err != nil {
			if strings.Contains(err.Error(),
				"a foreign key constraint fails (`school_management`.`students`, CONSTRAINT `students_ibfk_1` FOREIGN KEY (`class`) REFERENCES `teachers` (`class`))") {
				return nil, utils.ErrorHandler(err, "class/class teacher does not exist!")
			}
			return nil, utils.ErrorHandler(err, "Error inserting student data into database")
		}

		// Get the last inserted ID
		lastId, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting student data into database")
		}

		newStudent.ID = int(lastId)
		addedStudents[i] = newStudent
	}

	return addedStudents, nil
}

// UpdateStudentByID updates an existing student's details by their ID.
func UpdateStudentByID(id int, updatedStudent models.Student) (models.Student, error) {
	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	// get the existing student from database
	query := "SELECT * FROM students WHERE id = ?"
	var studentToUpdate models.Student
	err = db.QueryRow(query, id).
		Scan(&studentToUpdate.ID, &studentToUpdate.FirstName, &studentToUpdate.LastName, &studentToUpdate.Email, &studentToUpdate.Class)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, utils.ErrorHandler(err, fmt.Sprintf("student with ID: %d not found in database", id))
		}
		return models.Student{}, utils.ErrorHandler(err, "Error updating student in the database")
	}

	// Check if there are any changes before updating
	if updatedStudent.FirstName == studentToUpdate.FirstName &&
		updatedStudent.LastName == studentToUpdate.LastName &&
		updatedStudent.Email == studentToUpdate.Email &&
		updatedStudent.Class == studentToUpdate.Class {

		return models.Student{}, fmt.Errorf("no changes detected in the student's details")
	}

	const updateStudentQuery = `
		UPDATE students
		SET first_name = ?, last_name = ?, email = ?, class = ?
		WHERE id = ?`

	updatedStudent.ID = studentToUpdate.ID
	values := append(utils.GetStructValues(updatedStudent), studentToUpdate.ID)
	_, err = db.Exec(updateStudentQuery, values...)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating student in the database")
	}
	return updatedStudent, nil
}

// PatchStudentsInDb performs partial updates on multiple students in the database.
func PatchStudentsInDb(updatedFields []map[string]interface{}) ([]models.Student, error) {

	var studentsFromDB []models.Student
	db, err := ConnectDb()
	if err != nil {
		return studentsFromDB, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Validate all fields before starting the transaction
	for _, studentUpdate := range updatedFields {
		id, err := utils.GetIDFromMap(studentUpdate)
		if err != nil {
			return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
		}

		studentToUpdate, err := getStudentByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return studentsFromDB, utils.ErrorHandler(err, fmt.Sprintf("student with ID: %d not found in database", id))
			}
			return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
		}

		validFields := utils.BuildValidFieldsMap(studentToUpdate)
		if err := utils.ValidateUpdateFields(models.Student{}, validFields, studentUpdate); err != nil {
			return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
	}

	for _, studentUpdate := range updatedFields {
		id, err := utils.GetIDFromMap(studentUpdate)
		if err != nil {
			tx.Rollback()
			return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
		}
		studentToUpdate, err := getStudentByID(tx, id)
		if err != nil {
			tx.Rollback()
			return studentsFromDB, utils.ErrorHandler(err, fmt.Sprintf("student with ID: %d not found in database", id))
		}

		validFields := utils.BuildValidFieldsMap(studentToUpdate)
		utils.ApplyUpdateToStruct(&studentToUpdate, validFields, studentUpdate)

		var updateFields []string
		var updateArgs []interface{}
		for key, value := range studentUpdate {
			if key == "id" {
				continue
			}
			updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
			updateArgs = append(updateArgs, value)
		}
		if len(updateFields) == 0 {
			tx.Rollback()
			return studentsFromDB, fmt.Errorf("no valid fields provided for update")
		}
		updateArgs = append(updateArgs, studentToUpdate.ID)
		updateStudentQuery := fmt.Sprintf("UPDATE students SET %s WHERE id = ?", strings.Join(updateFields, ", "))

		_, err = tx.Exec(updateStudentQuery, updateArgs...)
		if err != nil {
			tx.Rollback()
			return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
		}

		studentsFromDB = append(studentsFromDB, studentToUpdate)
	}

	if err := tx.Commit(); err != nil {
		return studentsFromDB, utils.ErrorHandler(err, "Error updating student data into database")
	}

	return studentsFromDB, nil
}

// PatchstudentByID performs a partial update on a single student by their ID.
func PatchStudentByID(id int, updatedFields map[string]interface{}) (models.Student, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Get existing student by id using helper
	studentToUpdate, err := getStudentByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Student{}, utils.ErrorHandler(err, fmt.Sprintf("Student with ID: %d not found in database", id))
		}
		return models.Student{}, utils.ErrorHandler(err, "Error updating student data into database")
	}

	// Build valid fields map using helper
	validFields := utils.BuildValidFieldsMap(studentToUpdate)

	// Validate fields to update using helper
	if err := utils.ValidateUpdateFields(models.Student{}, validFields, updatedFields); err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating student data into database")
	}

	// Apply updates to struct using helper
	utils.ApplyUpdateToStruct(&studentToUpdate, validFields, updatedFields)

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
		return models.Student{}, fmt.Errorf("no valid fields provided for update")
	}

	updateArgs = append(updateArgs, studentToUpdate.ID)
	updateStudentQuery := fmt.Sprintf("UPDATE students SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateStudentQuery, updateArgs...)
	if err != nil {
		return models.Student{}, utils.ErrorHandler(err, "Error updating student data into database")
	}

	return studentToUpdate, nil
}

// DeleteStudentByID deletes a single student by their ID.
func DeleteStudentByID(id int) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Delete the student
	result, err := db.Exec("DELETE FROM students WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting student from database")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting student from database")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student with ID %d not found", id)
	}
	return nil
}

// DeleteStudentsInDB deletes multiple students by their IDs and
// returns the list of deleted IDs.
func DeleteStudentsInDB(IDs []int) ([]int, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error deleting students from database")
	}

	stmt, err := tx.Prepare("DELETE FROM students WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return nil, utils.ErrorHandler(err, "Error deleting students from database")
	}
	defer stmt.Close()

	// type myInt int
	deletedIDs := []int{}

	for _, id := range IDs {
		// Delete the student
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, fmt.Sprintf("Failed to delete student with ID %d: %v", id, err))
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return nil, utils.ErrorHandler(err, "Error deleting students from database")
		}

		if rowsAffected == 0 {
			tx.Rollback()
			return nil, fmt.Errorf("student with ID %d not found", id)
		}

		deletedIDs = append(deletedIDs, id)
	}

	err = tx.Commit()
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error deleting students from database")
	}

	if len(deletedIDs) == 0 {
		return nil, fmt.Errorf("no students found")
	}
	return deletedIDs, nil
}
