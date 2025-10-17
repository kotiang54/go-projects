package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"school_management_api/internal/api/middlewares"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	// Path parameters e.g. /teachers/{id}
	// Query parameters e.g. /teachers/?key=value&query=value2&sortBy=email&sortOrder=ASC

	switch r.Method {
	case http.MethodGet:
		// Path
		fmt.Println("Path:", r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("The teacher ID is:", userID)

		// Query
		fmt.Println("Query:", r.URL.Query())
		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortBy")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortOrder")
		if sortorder == "" {
			sortorder = "DESC"
		}

		// Print values
		fmt.Printf("Sortby: %v, Sort order: %v, Key: %v\n", sortby, key, sortorder)

		w.Write([]byte("Hello GET method on Teachers Route"))
		fmt.Println("Hello GET method on Teachers Route")
		return

	case http.MethodPost:
		// Parse RAW Body data
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()

		fmt.Println("RAW Body:", string(body))

		// If you expect a json data, then unmarshall
		var userInstance User
		err = json.Unmarshal(body, &userInstance)
		if err != nil {
			log.Fatalln("Unmarshall error:", err)
			return
		}

		fmt.Println(userInstance)
		fmt.Println("Receved user name as:", userInstance.Name)

		// Access the request details:
		fmt.Println("Body:", r.Body)
		fmt.Println("Form:", r.Form)
		fmt.Println("Header:", r.Header)
		fmt.Println("Context:", r.Context())
		fmt.Println("Host:", r.Host)
		fmt.Println("Method:", r.Method)
		fmt.Println("Protocol:", r.Proto)
		fmt.Println("Remote Address:", r.RemoteAddr)
		fmt.Println("Request URI:", r.RequestURI)
		fmt.Println("URL:", r.URL)
		fmt.Println("Port:", r.URL.Port())
		fmt.Println("TLS:", r.TLS)
		fmt.Println("Trailer:", r.Trailer)
		fmt.Println("User Agent:", r.UserAgent())
		fmt.Println("Transfer Encoding:", r.TransferEncoding)

		w.Write([]byte("Hello POST method on Teachers Route"))
		fmt.Println("Hello POST method on Teachers Route")
		return

	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
		fmt.Println("Hello PUT method on Teachers Route")
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		fmt.Println("Hello PATCH method on Teachers Route")
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		fmt.Println("Hello DELETE method on Teachers Route")
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students Route"))
		fmt.Println("Hello GET method on Students Route")
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students Route"))
		fmt.Println("Hello POST method on Students Route")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
		fmt.Println("Hello PUT method on Students Route")
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
		fmt.Println("Hello PATCH method on Students Route")
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))
		fmt.Println("Hello DELETE method on Students Route")
		return
	}
}

func executivesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Executives Route"))
		fmt.Println("Hello GET method on Executives Route")
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Executives Route"))
		fmt.Println("Hello POST method on Executives Route")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Executives Route"))
		fmt.Println("Hello PUT method on Executives Route")
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Executives Route"))
		fmt.Println("Hello PATCH method on Executives Route")
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Executives Route"))
		fmt.Println("Hello DELETE method on Executives Route")
		return
	}
}

func main() {
	// Main entry of the api

	port := 3000
	cert := "cert.pem"
	key := "key.pem"

	// Multiplexer for http routes
	mux := http.NewServeMux()

	// Create a routes
	mux.HandleFunc("/", rootHandler)

	// Teachers route
	mux.HandleFunc("/teachers/", teachersHandler)

	// Students route
	mux.HandleFunc("/students/", studentsHandler)

	// Executives route
	mux.HandleFunc("/executives/", executivesHandler)

	fmt.Println("Server is running on port:", port)

	// Make HTTP 1.1 with TLS server
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// Create a custom server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      middlewares.SecurityHeaders(mux),
		TLSConfig:    tlsConfig,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
