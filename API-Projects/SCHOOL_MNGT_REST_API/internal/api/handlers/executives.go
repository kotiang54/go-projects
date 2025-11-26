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

func GetExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	var executives []models.Executive
	executives, err := sqlconnect.GetExecutivesInDb(executives, r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string             `json:"status"`
		Count  int                `json:"count"`
		Data   []models.Executive `json:"data"`
	}{
		Status: "success",
		Count:  len(executives),
		Data:   executives,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	// Variable validations
	var newExecutives []models.Executive
	var rawExecutives []map[string]any

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &rawExecutives)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the fields in the incoming request.
	validFields := GetFieldNames(models.Executive{})

	// Validate each executive object in the incoming request
	for _, executive := range rawExecutives {
		for key := range executive {
			if _, ok := validFields[key]; !ok {
				http.Error(w, fmt.Sprintf("Unacceptable field: %s, found in request.", key), http.StatusBadRequest)
				return
			}
		}
	}

	// Decode the request body into a slice of Executive structs
	err = json.Unmarshal(body, &newExecutives)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Validate the newExecutives fields
	for _, executive := range newExecutives {
		err = CheckBlankFields(executive)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecutives, err := sqlconnect.CreateExecutives(newExecutives)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Response structure with status, count, and data
	response := struct {
		Status string             `json:"status"`
		Count  int                `json:"count"`
		Data   []models.Executive `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecutives),
		Data:   addedExecutives,
	}

	json.NewEncoder(w).Encode(response)
}

func PatchExecutivesHandler(w http.ResponseWriter, r *http.Request) {
	var updatedFields []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	executivesFromDB, err := sqlconnect.PatchExecutivesInDb(updatedFields)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(executivesFromDB)
}

func GetOneExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Handle Path parameters for specific executive
	executiveIDStr := r.PathValue("id")
	id, err := strconv.Atoi(executiveIDStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Executive ID: %s", executiveIDStr), http.StatusBadRequest)
		return
	}

	executive, err := sqlconnect.GetExecutiveByID(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(executive)
}

func PatchOneExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	/// Get the executive id from the path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Executive ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Decode fields to update from request body
	var updatedFields map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updatedFields)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}

	executiveToUpdate, err := sqlconnect.PatchExecutiveByID(id, updatedFields)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated executive
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(executiveToUpdate)
}

func DeleteOneExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Get the executives Id from the path
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid Executive ID: %s", idStr), http.StatusBadRequest)
		return
	}

	// Connect to database
	err = sqlconnect.DeleteExecutiveByID(id)
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
		Message: fmt.Sprintf("Executive with ID %d deleted successfully", id),
	}

	json.NewEncoder(w).Encode(response)
}

func UpdatePasswordExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an executive's password
}

func LoginExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for executive login
}

func LogoutExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for executive logout
}

func ForgotPasswordExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for forgot password functionality
}

func ResetPasswordExecutiveHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for resetting password functionality
}
