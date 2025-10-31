package handlers

import (
	"encoding/json"
	"net/http"
	"school_management_api/internal/models"
	"strconv"
	"strings"
	"sync"
)

// in-memory slice to hold teachers data
var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// Initialize dummy data
func init() {
	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Mathematics",
	}
	nextID++

	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10B",
		Subject:   "Science",
	}
	nextID++

	teachers[nextID] = models.Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "8C",
		Subject:   "English",
	}
	nextID++
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Path parameters e.g. /teachers/{id}
	// Query parameters e.g. /teachers/?key=value&query=value2&sortBy=email&sortOrder=ASC

	switch r.Method {
	case http.MethodGet:
		// Handle GET request to fetch all teachers
		getTeachersHandler(w, r)

	case http.MethodPost:
		// Handle POST request to create a new teacher
		createTeachersHandler(w, r)

	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		return
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Path parameters can be handled here if needed
	// e.g. teacherID := chi.URLParam(r, "id")

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	teacherIDStr := strings.TrimSuffix(path, "/")

	if teacherIDStr == "" {
		// Handle query parameters for filtering
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]models.Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			// Simple filtering logic
			if (firstName == "" || teacher.FirstName == firstName) &&
				(lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string           `json:"status"`
			Count  int              `json:"count"`
			Data   []models.Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameters for specific teacher
	id, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func createTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a new teacher
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []models.Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid input body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
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
