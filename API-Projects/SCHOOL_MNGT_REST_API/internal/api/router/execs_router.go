package router

import (
	"net/http"
	"school_management_api/internal/api/handlers"
)

func execsRouter() *http.ServeMux {
	// Define the router for executive-related routes
	mux := http.NewServeMux()

	// Executives route
	mux.HandleFunc("GET /executives", handlers.GetExecutivesHandler)
	mux.HandleFunc("POST /executives", handlers.CreateExecutivesHandler)
	mux.HandleFunc("PATCH /executives", handlers.PatchExecutivesHandler)

	// Executives route with ID
	mux.HandleFunc("GET /executives/{id}", handlers.GetOneExecutiveHandler)
	mux.HandleFunc("PATCH /executives/{id}", handlers.PatchOneExecutiveHandler)
	mux.HandleFunc("DELETE /executives/{id}", handlers.DeleteOneExecutiveHandler)
	mux.HandleFunc("POST /executives/{id}/updatepassword", handlers.UpdatePasswordHandler)

	mux.HandleFunc("POST /executives/login", handlers.LoginHandler)
	mux.HandleFunc("POST /executives/logout", handlers.LogoutHandler)
	mux.HandleFunc("POST /executives/forgotpassword", handlers.ForgotPasswordHandler)
	mux.HandleFunc("POST /executives/resetpassword/reset/{resetcode}", handlers.ResetPasswordHandler)

	return mux
}
