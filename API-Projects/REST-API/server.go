package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define route handlers
	http.HandleFunc("/orders", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Handling users")
	})

	// Define the port to listen on
	port := 3000

	// Start the server
	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
