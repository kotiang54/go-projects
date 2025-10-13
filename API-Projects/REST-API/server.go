package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// Define route handlers

	// The api does not respond to a specific http method
	// It responds to all http methods (GET, POST, PUT, DELETE, etc.)
	// You can use curl or Postman to test the endpoints with different methods

	// Example: curl -X GET http://localhost:3000/orders
	// Example: curl -X POST http://localhost:3000/users

	http.HandleFunc("/orders", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, "Handling users")
	})

	// Define the port to listen on
	port := 3000

	// Modify the code to handle TLS/SSL for secure communication (HTTPS)
	// You will need to generate a self-signed certificate for testing purposes
	// In a production environment, use a valid SSL certificate from a trusted CA

	// For HTTPS, use ListenAndServeTLS method instead
	// Make sure to generate cert.pem and key.pem files using OpenSSL
	// You can use the following command to generate a self-signed certificate:
	// 	- openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// configure the TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
		Handler:   nil, // Use default mux
	}

	// Enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)
	// Start the server with HTTP 1.1 - Server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	// Start the server with HTTP 2.0 - Server with TLS
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
