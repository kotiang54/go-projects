package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

// GetTeachersHandler handles GET requests to fetch teachers
func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var teachers []models.Teacher
	teachers, err := sqlconnect.GetTeachersInDb(teachers, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, fmt.Sprintf("Invalid Teacher ID: %s", teacherIDStr), http.StatusBadRequest)
		return
	}

	teacher, err := sqlconnect.GetTeacherByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		log.Println(err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	val := reflect.TypeOf(models.Teacher{})
	validFields := make(map[string]int)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" {
			validFields[jsonTag] = i
		}
	}

	// Validate the newTeachers fields
	for _, teacher := range newTeachers {
		val := reflect.ValueOf(teacher)
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fmt.Println(field)
			if field.Kind() == reflect.String && field.Len() == 0 {
				http.Error(w, "All fields (first_name, last_name, email, class, subject) are required", http.StatusBadRequest)
				return
			}
		}
	}

	addedTeachers, err := sqlconnect.CreateTeachers(newTeachers)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	err = sqlconnect.DeleteTeacherByID(id)
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
		Message: fmt.Sprintf("Teacher with ID %d deleted successfully", id),
	}

	json.NewEncoder(w).Encode(response)
}

// DeleteTeachersHandler handles DELETE requests to remove teachers record
func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var IDs []int
	err := json.NewDecoder(r.Body).Decode(&IDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	deletedIDs, err := sqlconnect.DeleteTeachersInDB(IDs)
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
		Status:     "Teachers successfully deleted",
		DeletedIDs: deletedIDs,
	}

	json.NewEncoder(w).Encode(response)
}
