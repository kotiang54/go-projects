package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
)

// GetTeachersHandler handles GET requests to fetch teachers
func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var teachers []models.Teacher
	teachers, err := sqlconnect.GetTeachersCollection(teachers, r)
	if err != nil {
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachers),
		Data:   teachers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetOneTeacherHandler handles GET requests to fetch a specific teacher
func GetOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Handle Path parameters for specific teacher
	teacherIDStr := r.PathValue("id")
	id, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return
	}

	teacher, err := sqlconnect.GetTeacherByID(id)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

// CreateTeachersHandler handles the creation of new teachers
func CreateTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a slice of Teacher structs
	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers, err := sqlconnect.CreateTeachers(newTeachers)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create teachers: %v", err), http.StatusInternalServerError)
		return
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
	result, err := sqlconnect.UpdateTeacherByID(id, updatedTeacher)
	if err != nil {
		return
	}

	// return updated teacher with status field
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string         `json:"status"`
		Data   models.Teacher `json:"data"`
	}{
		Status: "success",
		Data:   result,
	}

	json.NewEncoder(w).Encode(response)
}

// PatchTeachersHandler handles PATCH requests to partially update teachers records
// PATCH /teachers/
func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var updatedFields []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	teachersFromDB, err := sqlconnect.PatchTeachersInDb(updatedFields)
	if err != nil {
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

	teacherToUpdate, err := sqlconnect.PatchTeacherByID(id, updatedFields)
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
			http.Error(w, fmt.Sprintf("Teacher with ID %d not found", id), http.StatusNotFound)
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
