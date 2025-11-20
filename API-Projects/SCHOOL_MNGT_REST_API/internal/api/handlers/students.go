package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
)

// GetStudentsHandler handles GET requests to fetch students
func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var students []models.Student
	students, err := sqlconnect.GetStudentsInDb(students, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetOneStudentHandler handles GET requests to fetch a specific student
func GetOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Handle Path parameters for specific student
	studentIDStr := r.PathValue("id")
	id, err := strconv.Atoi(studentIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Student ID: %s", studentIDStr), http.StatusBadRequest)
		return
	}

	student, err := sqlconnect.GetStudentByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// CreateStudentsHandler handles the creation of new students
func CreateStudentsHandler(w http.ResponseWriter, r *http.Request) {

	// Variable validations
	var newStudents []models.Student
	var rawStudents []map[string]any

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &rawStudents)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the fields in the incoming request.
	validFields := GetFieldNames(models.Student{})

	// Validate each student object in the incoming request
	for _, student := range rawStudents {
		for key := range student {
			if _, ok := validFields[key]; !ok {
				http.Error(w, fmt.Sprintf("Unacceptable field: %s, found in request.", key), http.StatusBadRequest)
				return
			}
		}
	}

	// Decode the request body into a slice of Student structs
	err = json.Unmarshal(body, &newStudents)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the newStudents fields
	for _, student := range newStudents {
		err = CheckBlankFields(student)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedStudents, err := sqlconnect.CreateStudents(newStudents)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Response structure with status, count, and data
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(addedStudents),
		Data:   addedStudents,
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateStudentsHandler handles updating an existing student
func UpdateStudentsHandler(w http.ResponseWriter, r *http.Request) {
	// get students id from path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Student ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// create updated student variable from request body
	var updatedStudent models.Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the updatedStudent fields
	if updatedStudent.FirstName == "" || updatedStudent.LastName == "" || updatedStudent.Email == "" || updatedStudent.Class == "" {
		http.Error(w, "All fields (first_name, last_name, email, class) are required", http.StatusBadRequest)
		return
	}

	// update student in database
	result, err := sqlconnect.UpdateStudentByID(id, updatedStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return updated student with status field
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string         `json:"status"`
		Data   models.Student `json:"data"`
	}{
		Status: "success",
		Data:   result,
	}

	json.NewEncoder(w).Encode(response)
}

// PatchStudentsHandler handles PATCH requests to partially update students records
// PATCH /students/
func PatchStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var updatedFields []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	studentsFromDB, err := sqlconnect.PatchStudentsInDb(updatedFields)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentsFromDB)
}

// PatchOneStudentHandler handles PATCH requests to partially update a student records
// PATCH /students/{id}
func PatchOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the student id from the path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Student ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Decode fields to update from request body
	var updatedFields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	studentToUpdate, err := sqlconnect.PatchStudentByID(id, updatedFields)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated student
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentToUpdate)
}

// DeleteOneStudentHandler handles DELETE requests to remove a student record
func DeleteOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the students Id from the path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Student ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Connect to database
	err = sqlconnect.DeleteStudentByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ----- Alternate approach -----
	// w.WriteHeader(http.StatusNoContent)

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  "success",
		Message: fmt.Sprintf("Student with ID %d deleted successfully", id),
	}

	json.NewEncoder(w).Encode(response)
}

// DeleteStudentsHandler handles DELETE requests to remove students record
func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var IDs []int
	err := json.NewDecoder(r.Body).Decode(&IDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	deletedIDs, err := sqlconnect.DeleteStudentsInDB(IDs)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "Students successfully deleted",
		DeletedIDs: deletedIDs,
	}

	json.NewEncoder(w).Encode(response)
}
