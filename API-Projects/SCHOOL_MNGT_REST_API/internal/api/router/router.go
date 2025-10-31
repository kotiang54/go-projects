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
	mux.HandleFunc("/teachers/", handlers.TeachersHandler)

	// Students route
	mux.HandleFunc("/students/", handlers.StudentsHandler)

	// Executives route
	mux.HandleFunc("/executives/", handlers.ExecutivesHandler)

	return mux
}
