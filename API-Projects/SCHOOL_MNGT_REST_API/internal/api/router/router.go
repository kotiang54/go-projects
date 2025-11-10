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
	mux.HandleFunc("GET /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("GET /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PUT /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.TeachersHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.TeachersHandler)

	// Students route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// Executives route
	mux.HandleFunc("/executives/", handlers.ExecutivesHandler)

	return mux
}
