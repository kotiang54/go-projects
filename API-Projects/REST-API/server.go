package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

// logRequestDetails logs the HTTP version and TLS version (if applicable) of the incoming request
func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received request with HTTP version:", httpVersion)

	// Check if the request is over TLS
	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version:", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

// getTLSVersionName returns the human-readable name of the TLS version
func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS version"
	}
}

// loadClientCAs loads a PEM-encoded CA certificate from the file "cert.pem",
// appends it to a new x509.CertPool, and returns the pool.
// This CertPool can be used to verify client certificates in mutual TLS setups.
func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load client CA:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {
	// Define route handlers

	// The api does not respond to a specific http method
	// It responds to all http methods (GET, POST, PUT, DELETE, etc.)
	// You can use curl or Postman to test the endpoints with different methods

	// Example: curl -X GET http://localhost:3000/orders
	// Example: curl -X POST http://localhost:3000/users

	http.HandleFunc("/orders", func(resp http.ResponseWriter, req *http.Request) {
		logRequestDetails(req)
		fmt.Fprintf(resp, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(resp http.ResponseWriter, req *http.Request) {
		logRequestDetails(req)
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
	// 	- or use a config file for more options:
	// Create a file named openssl.cnf with the following content:
	/*
		[req]
		default_bits       = 2048
		prompt             = no
		default_md         = sha256
		distinguished_name = dn
		x509_extensions    = v3_req

		[dn]
		C = US
		ST = California
		L = San Francisco
		O = My Company
		CN = localhost

		[v3_req]
		subjectAltName = @alt_names

		[alt_names]
		DNS.1 = localhost
		DNS.2 = example.com
	*/
	// Then run the following command to generate the cert and key:
	// openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// configure the TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,

		// enforce mTLS example
		// To disable mutual TLS (mTLS), comment out the following two lines.
		// With these lines enabled, the server requires clients to present a valid certificate signed by the CA in cert.pem.
		// This enforces mTLS, meaning both server and client must authenticate each other.
		// ClientAuth: tls.RequireAndVerifyClientCert,
		// ClientCAs:  loadClientCAs(),
	}

	// Create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		TLSConfig: tlsConfig,
		Handler:   nil, // Use default mux
		// Disable http2 protocol from the API - comment out for http2 --> http 1.1
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	// Uncomment to enable http2
	// http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)
	// Start the server with HTTP 1.1 - Server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	// Start the server with HTTP 1.1 / 2.0 - Server with TLS
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
