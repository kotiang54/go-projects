package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func Router() *http.ServeMux {
	// Multiplexer for http routes
	mux := http.NewServeMux()

	// Create a routes
	mux.HandleFunc("/", handlers.RootHandler)

	// Teachers route
	mux.HandleFunc("GET /teachers/", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.CreateTeachersHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteTeachersHandler)

	// Teachers route with ID
	mux.HandleFunc("GET /teachers/{id}", handlers.GetTeacherHandler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteTeachersHandler)

	// Students route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// Executives route
	mux.HandleFunc("/executives/", handlers.ExecutivesHandler)

	return mux
}
