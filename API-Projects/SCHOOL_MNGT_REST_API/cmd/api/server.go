package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Main entry of the api

	port := 3000

	// Create a routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello Root Route")
		w.Write([]byte("Hello Root Route"))
		fmt.Println("Hello Root Route")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello GET method on Teachers Route"))
			fmt.Println("Hello GET method on Teachers Route")
			return

		case http.MethodPost:
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

		w.Write([]byte("Hello Teachers Route"))
		fmt.Println("Hello Teachers Route")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Students Route"))
		fmt.Println("Hello Students Route")
	})

	http.HandleFunc("/executives", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Executives Route"))
		fmt.Println("Hello Executives Route")
	})

	fmt.Println("Server is running on port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
