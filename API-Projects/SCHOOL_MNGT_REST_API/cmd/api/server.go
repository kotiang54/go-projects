package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

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
