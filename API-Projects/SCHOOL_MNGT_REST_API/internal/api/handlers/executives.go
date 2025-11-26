package handlers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/internal/repository/sqlconnect"
	"school_management_api/pkg/utils"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
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

func UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an executive's password
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for executive login
	var req models.Executive

	// data validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// search for user if exists
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var userExec models.Executive
	query := "SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?"
	err = db.QueryRow(query, req.Username).
		Scan(&userExec.ID, &userExec.FirstName, &userExec.LastName, &userExec.Email, &userExec.Username, &userExec.Password, &userExec.InactiveStatus, &userExec.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Executive not found", http.StatusNotFound)
			return
		}
		http.Error(w, "database query error", http.StatusNotFound)
		return
	}

	// is user active
	if userExec.InactiveStatus {
		http.Error(w, "Executive account is inactive", http.StatusForbidden)
		return
	}

	// verify password
	// split stored password into salt and hash
	parts := strings.Split(userExec.Password, ".")
	if len(parts) != 2 {
		utils.ErrorHandler(errors.New("invalid stored password format"), "invalid encoded hash format")
		http.Error(w, "Invalid stored password format", http.StatusForbidden)
		return
	}

	saltBase64, hashBase64 := parts[0], parts[1]
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the salt")
		http.Error(w, "Failed to decode the salt", http.StatusForbidden)
		return
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the hashed password")
		http.Error(w, "Failed to decode the hashed password", http.StatusForbidden)
		return
	}

	hash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		utils.ErrorHandler(errors.New("incorrect password"), "password verification failed")
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	// constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(hash, hashedPassword) != 1 {
		utils.ErrorHandler(errors.New("incorrect password"), "password verification failed")
		http.Error(w, "Incorrect password", http.StatusForbidden)
		return
	}

	// generate token
	token := "abc"

	// send token as response or as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "Bearer", //"exec_auth_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	json.NewEncoder(w).Encode(response)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for executive logout
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for forgot password functionality
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for resetting password functionality
}
