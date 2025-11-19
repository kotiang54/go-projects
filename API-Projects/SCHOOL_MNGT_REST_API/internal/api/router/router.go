package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func Router() *http.ServeMux {
	// Multiplexer for http routes
	mux := http.NewServeMux()

	// Create routes
	mux.HandleFunc("GET /", handlers.RootHandler)

	// Teachers route
	mux.HandleFunc("GET /teachers/", handlers.GetTeachersHandler)
	mux.HandleFunc("POST /teachers/", handlers.CreateTeachersHandler)
	mux.HandleFunc("PATCH /teachers/", handlers.PatchTeachersHandler)
	mux.HandleFunc("DELETE /teachers/", handlers.DeleteTeachersHandler)

	// Teachers route with ID
	mux.HandleFunc("GET /teachers/{id}", handlers.GetOneTeacherHandler)
	mux.HandleFunc("PUT /teachers/{id}", handlers.UpdateTeachersHandler)
	mux.HandleFunc("PATCH /teachers/{id}", handlers.PatchOneTeacherHandler)
	mux.HandleFunc("DELETE /teachers/{id}", handlers.DeleteOneTeacherHandler)

	// Students route
	mux.HandleFunc("GET /students/", handlers.GetStudentsHandler)
	mux.HandleFunc("POST /students/", handlers.CreateStudentsHandler)
	mux.HandleFunc("PATCH /students/", handlers.PatchStudentsHandler)
	mux.HandleFunc("DELETE /students/", handlers.DeleteStudentsHandler)

	// Students route with ID
	mux.HandleFunc("GET /students/{id}", handlers.GetOneStudentHandler)
	mux.HandleFunc("PUT /students/{id}", handlers.UpdateStudentsHandler)
	mux.HandleFunc("PATCH /students/{id}", handlers.PatchOneStudentHandler)
	mux.HandleFunc("DELETE /students/{id}", handlers.DeleteOneStudentHandler)

	// Executives route
	mux.HandleFunc("/executives/", handlers.ExecutivesHandler)

	return mux
}
