package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {

	tRouter := teachersRouter()
	sRouter := studentsRouter()

	tRouter.Handle("/", sRouter)
	return tRouter
	// Multiplexer for http routes
	// mux := http.NewServeMux()

	// // Create routes
	// mux.HandleFunc("GET /", handlers.RootHandler)

	// // Teachers route

	// // Students route

	// // Executives route
	// mux.HandleFunc("GET /executives/", handlers.ExecutivesHandler)

	// return mux
}
