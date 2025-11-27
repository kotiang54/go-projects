package sqlconnect

import (
	"crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"school_management_api/internal/models"
	"school_management_api/pkg/utils"
	"strings"

	"golang.org/x/crypto/argon2"
)

// =========== Helper functions ===================

// addExecutivesFilter adds filtering conditions to the SQL query based on URL query parameters.
func addExecutivesFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	// Handle Query parameters for filtering
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"username":   "username",
		"role":       "role",
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

// getExecutiveByID retrieves a executive by ID from the database
func getExecutiveByID(db queryer, id int) (models.Executive, error) {
	var executive models.Executive
	query := "SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE id = ?"
	err := db.QueryRow(query, id).
		Scan(&executive.ID, &executive.FirstName, &executive.LastName, &executive.Email, &executive.Username, &executive.UserCreatedAt, &executive.InactiveStatus, &executive.Role)
	return executive, err
}

// ================ Database Operations ===================

// GetExecutivesInDb retrieves a collection of executives from the database
// with optional filtering and sorting.
func GetExecutivesInDb(executives []models.Executive, r *http.Request) ([]models.Executive, error) {
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
	query := "SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE 1=1" // * id, first_name, last_name, email
	var args []interface{}

	// Add filters based on query parameters
	query, args = addExecutivesFilter(r, query, args)

	// Example: /executives/?sortby=last_name:asc&sortby=username:desc
	query += utils.BuildOrderByClause(r)

	// Execute the query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error retrieving executives from database")
	}

	defer rows.Close()

	for rows.Next() {
		var executive models.Executive
		if err := rows.Scan(&executive.ID, &executive.FirstName, &executive.LastName, &executive.Email, &executive.Username, &executive.UserCreatedAt, &executive.InactiveStatus, &executive.Role); err != nil {
			return nil, utils.ErrorHandler(err, "Error retrieving executives from database")
		}
		executives = append(executives, executive)
	}
	return executives, nil
}

// GetExecutiveByID retrieves a single executive by their ID.
func GetExecutiveByID(id int) (models.Executive, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Executive{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	var executive models.Executive
	query := "SELECT id, first_name, last_name, email, username, user_created_at, inactive_status, role FROM execs WHERE id = ?" // id, first_name, last_name, email, username, user_created_at, inactive_status, role
	err = db.QueryRow(query, id).
		Scan(&executive.ID, &executive.FirstName, &executive.LastName, &executive.Email, &executive.Username, &executive.UserCreatedAt, &executive.InactiveStatus, &executive.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Executive{}, utils.ErrorHandler(err, fmt.Sprintf("Executive with ID: %d not found in database", id))
		}
		return models.Executive{}, utils.ErrorHandler(err, "Error retrieving executive by ID from database")
	}
	return executive, nil
}

// CreateExecutives adds new executives to the database.
func CreateExecutives(newExecutives []models.Executive) ([]models.Executive, error) {
	db, err := ConnectDb()
	if err != nil {
		return []models.Executive{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	stmt, err := db.Prepare(utils.GenerateInsertQuery(models.Executive{}, "execs"))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Error inserting executive data into database")
	}
	defer stmt.Close()

	addedExecutives := make([]models.Executive, len(newExecutives))

	for i, newExecutive := range newExecutives {
		if newExecutive.Password == "" {
			return nil, utils.ErrorHandler(fmt.Errorf("password is required"), "Error inserting executive data into database")
		}

		// Hash the password using Argon2id with a random salt
		salt := make([]byte, 16)
		if _, err := rand.Read(salt); err != nil {
			return nil, utils.ErrorHandler(errors.New("failed to generate salt"), "Error inserting executive data into database")
		}

		hash := argon2.IDKey([]byte(newExecutive.Password), salt, 1, 64*1024, 4, 32)
		saltBase64 := base64.StdEncoding.EncodeToString(salt)
		hashBase64 := base64.StdEncoding.EncodeToString(hash)
		encodedHash := fmt.Sprintf("%s.%s", saltBase64, hashBase64)

		newExecutive.Password = encodedHash

		values := utils.GetStructValues(newExecutive)
		res, err := stmt.Exec(values...)
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting executive data into database")
		}

		// Get the last inserted ID
		lastId, err := res.LastInsertId()
		if err != nil {
			return nil, utils.ErrorHandler(err, "Error inserting executive data into database")
		}

		newExecutive.ID = int(lastId)
		addedExecutives[i] = newExecutive
	}

	return addedExecutives, nil
}

// PatchExecutivesInDb performs partial updates on multiple executives in the database.
func PatchExecutivesInDb(updatedFields []map[string]interface{}) ([]models.Executive, error) {

	var executivesFromDB []models.Executive
	db, err := ConnectDb()
	if err != nil {
		return executivesFromDB, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Validate all fields before starting the transaction
	for _, executiveUpdate := range updatedFields {
		id, err := utils.GetIDFromMap(executiveUpdate)
		if err != nil {
			return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
		}

		executiveToUpdate, err := getExecutiveByID(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return executivesFromDB, utils.ErrorHandler(err, fmt.Sprintf("executive with ID: %d not found in database", id))
			}
			return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
		}

		validFields := utils.BuildValidFieldsMap(executiveToUpdate)
		if err := utils.ValidateUpdateFields(models.Executive{}, validFields, executiveUpdate); err != nil {
			return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
	}

	for _, executiveUpdate := range updatedFields {
		id, err := utils.GetIDFromMap(executiveUpdate)
		if err != nil {
			tx.Rollback()
			return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
		}
		executiveToUpdate, err := getExecutiveByID(tx, id)
		if err != nil {
			tx.Rollback()
			return executivesFromDB, utils.ErrorHandler(err, fmt.Sprintf("executive with ID: %d not found in database", id))
		}

		validFields := utils.BuildValidFieldsMap(executiveToUpdate)
		utils.ApplyUpdateToStruct(&executiveToUpdate, validFields, executiveUpdate)

		var updateFields []string
		var updateArgs []interface{}
		for key, value := range executiveUpdate {
			if key == "id" {
				continue
			}
			updateFields = append(updateFields, fmt.Sprintf("%s = ?", key))
			updateArgs = append(updateArgs, value)
		}
		if len(updateFields) == 0 {
			tx.Rollback()
			return executivesFromDB, fmt.Errorf("no valid fields provided for update")
		}
		updateArgs = append(updateArgs, executiveToUpdate.ID)
		updateExecutiveQuery := fmt.Sprintf("UPDATE execs SET %s WHERE id = ?", strings.Join(updateFields, ", "))

		_, err = tx.Exec(updateExecutiveQuery, updateArgs...)
		if err != nil {
			tx.Rollback()
			return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
		}

		executivesFromDB = append(executivesFromDB, executiveToUpdate)
	}

	if err := tx.Commit(); err != nil {
		return executivesFromDB, utils.ErrorHandler(err, "Error updating executive data into database")
	}

	return executivesFromDB, nil
}

// PatchexecutiveByID performs a partial update on a single executive by their ID.
func PatchExecutiveByID(id int, updatedFields map[string]interface{}) (models.Executive, error) {

	db, err := ConnectDb()
	if err != nil {
		return models.Executive{}, utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Get existing executive by id using helper
	executiveToUpdate, err := getExecutiveByID(db, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Executive{}, utils.ErrorHandler(err, fmt.Sprintf("Executive with ID: %d not found in database", id))
		}
		return models.Executive{}, utils.ErrorHandler(err, "Error updating executive data into database")
	}

	// Build valid fields map using helper
	validFields := utils.BuildValidFieldsMap(executiveToUpdate)

	// Validate fields to update using helper
	if err := utils.ValidateUpdateFields(models.Executive{}, validFields, updatedFields); err != nil {
		return models.Executive{}, utils.ErrorHandler(err, "Error updating executive data into database")
	}

	// Apply updates to struct using helper
	utils.ApplyUpdateToStruct(&executiveToUpdate, validFields, updatedFields)

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
		return models.Executive{}, fmt.Errorf("no valid fields provided for update")
	}

	updateArgs = append(updateArgs, executiveToUpdate.ID)
	updateExecutiveQuery := fmt.Sprintf("UPDATE execs SET %s WHERE id = ?", strings.Join(updateFields, ", "))

	_, err = db.Exec(updateExecutiveQuery, updateArgs...)
	if err != nil {
		return models.Executive{}, utils.ErrorHandler(err, "Error updating executive data into database")
	}

	return executiveToUpdate, nil
}

// DeleteExecutiveByID deletes a single executive by their ID.
func DeleteExecutiveByID(id int) error {
	db, err := ConnectDb()
	if err != nil {
		return utils.ErrorHandler(err, "Error connecting to database")
	}
	defer db.Close()

	// Delete the executive
	result, err := db.Exec("DELETE FROM execs WHERE id = ?", id)
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting executive from database")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.ErrorHandler(err, "Error deleting executive from database")
	}

	if rowsAffected == 0 {
		return fmt.Errorf("executive with ID %d not found", id)
	}
	return nil
}

// ExecLogin verifies executive login credentials.
func ExecLogin(username string, password string) (string, error) {
	db, err := ConnectDb()
	if err != nil {
		return "", utils.ErrorHandler(err, "Database connection error")
	}
	defer db.Close()

	// search for user if exists
	var userExec models.Executive
	query := "SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?"
	err = db.QueryRow(query, username).
		Scan(&userExec.ID, &userExec.FirstName, &userExec.LastName, &userExec.Email, &userExec.Username, &userExec.Password, &userExec.InactiveStatus, &userExec.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", utils.ErrorHandler(err, "executive not found in database")
		}
		return "", utils.ErrorHandler(err, "database query error")
	}

	// is user active
	if userExec.InactiveStatus {
		return "", utils.ErrorHandler(errors.New("executive account is inactive"), "executive account is inactive")
	}

	// verify password
	// split stored password into salt and hash
	parts := strings.Split(userExec.Password, ".")
	if len(parts) != 2 {
		return "", utils.ErrorHandler(errors.New("invalid stored password format"), "invalid encoded hash format")
	}

	saltBase64, hashBase64 := parts[0], parts[1]
	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return "", utils.ErrorHandler(err, "failed to decode the salt")
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		return "", utils.ErrorHandler(err, "failed to decode the hashed password")
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		return "", utils.ErrorHandler(errors.New("incorrect password"), "password verification failed")
	}

	// constant time comparison to prevent timing attacks
	if subtle.ConstantTimeCompare(hash, hashedPassword) != 1 {
		return "", utils.ErrorHandler(errors.New("incorrect password"), "password verification failed")
	}

	// generate token
	token, err := utils.SignToken(userExec.ID, userExec.Username, userExec.Role)
	if err != nil {
		return "", utils.ErrorHandler(err, "failed to generate token")
	}

	return token, nil
}
