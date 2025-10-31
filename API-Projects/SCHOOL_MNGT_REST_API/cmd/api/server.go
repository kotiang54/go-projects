package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "school_management_api/internal/api/middlewares"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

// in-memory slice to hold teachers data
var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// Initialize dummy data
func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Mathematics",
	}
	nextID++

	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10B",
		Subject:   "Science",
	}
	nextID++

	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "8C",
		Subject:   "English",
	}
	// nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	// Path parameters can be handled here if needed
	// e.g. teacherID := chi.URLParam(r, "id")

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	teacherIDStr := strings.TrimSuffix(path, "/")

	if teacherIDStr == "" {
		// Handle query parameters for filtering
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			// Simple filtering logic
			if (firstName == "" || teacher.FirstName == firstName) &&
				(lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameters for specific teacher
	id, err := strconv.Atoi(teacherIDStr)
	if err != nil {
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
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
		// Handle GET request to fetch all teachers
		getTeachersHandler(w, r)

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
		return

	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		return
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Students Route"))
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))
		return
	}
}

func executivesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Executives Route"))
		return

	case http.MethodPost:
		w.Write([]byte("Hello POST method on Executives Route"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Executives Route"))
		return

	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Executives Route"))
		return

	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Executives Route"))
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

	// rate limiting middleware can be added here
	// rl := mw.NewRateLimiter(5, time.Minute)

	// HPP middleware options configuration
	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                 true,
	// 	CheckBody:                  true,
	// 	CheckBodyOnlyForContenType: "application/x-www-form-urlencoded",
	// 	Whitelist:                  []string{"sortBy", "sortOrder", "first_name", "last_name", "class"},
	// }

	// Recommended middleware order (from outermost to innermost)
	// secureMux := mw.Cors( // 1. CORS: Handle cross-origin and preflight requests first
	// 	mw.Hpp(hppOptions)( // 2. HPP: Sanitize query/body params before any logic uses them
	// 		rl.Middleware( // 3. Rate Limiting: Block abusive clients early, before expensive work
	// 			mw.SecurityHeaders( // 4. Security Headers: Set headers for all responses
	// 				mw.ResponseTime( // 5. Response Time: Measure as much as possible
	// 					mw.Compression( // 6. Compression: Compress the final response
	// 						mux, // 7. Your main router/handler
	// 					),
	// 				),
	// 			),
	// 		),
	// 	),
	// )

	// Using helper function to apply middlewares
	secureMux := applyMiddlewares(mux,
		// mw.Compression,     // 6. Compression: Compress the final response
		// mw.ResponseTime,    // 5. Response Time: Measure as much as possible
		mw.SecurityHeaders, // 4. Security Headers: Set headers for all responses
		// rl.Middleware,      // 3. Rate Limiting: Block abusive clients early, before expensive work
		// mw.Hpp(hppOptions), // 2. HPP: Sanitize query/body params before any logic uses them
		// mw.Cors,            // 1. CORS: Handle cross-origin and preflight requests first
	)

	// Create a custom server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      secureMux,
		TLSConfig:    tlsConfig,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}

// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func applyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
