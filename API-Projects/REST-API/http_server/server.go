package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// Handle requests to the root URL
		fmt.Fprintln(resp, "Hello, World! This is a simple HTTP server.")
	})

	// Create a simple HTTP server that listens on port 3000
	const serverAddr string = "127.0.0.1:3000"

	fmt.Println("Server is listening on http://" + serverAddr)
	// Handle requests to the root URL
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("Error starting server: " + err.Error())
	}
}
