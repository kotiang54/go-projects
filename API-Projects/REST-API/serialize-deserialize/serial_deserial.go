package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// User defines a simple struct with JSON tags for serialization/deserialization
type User struct {
	Name  string `json:"name"`  // Name will be mapped to "name" in JSON
	Email string `json:"email"` // Email will be mapped to "email" in JSON
}

func main() {
	// Create a User instance
	user := User{Name: "Alice", Email: "alice09@example.com"}
	// fmt.Println(user)

	// Marshal the Go object to JSON (serialize)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalln("Error marshalling Go object:", err)
	}

	fmt.Println("Marshalled json string:", string(jsonData))

	// Using json.Unmarshal with a byte slice (in-memory approach) - deserialize
	var user1 User
	if err = json.Unmarshal(jsonData, &user1); err != nil {
		log.Fatalln("Error unmarshalling the json object:", err)
	}

	fmt.Println("Unmarshalled data from json data:", user1)

	// Encoding json objects
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Encoded json string:", buf.String())

	// Example JSON string to Decode
	jsonData1 := `{"name": "John", "email": "john25@example.com"}`

	// Using json.NewDecoder with an io.Reader (streaming approach)
	reader := strings.NewReader(jsonData1) // Create a reader from the JSON string
	decoder := json.NewDecoder(reader)     // Create a new JSON decoder

	var user2 User
	err = decoder.Decode(&user2) // Decode JSON from the reader into user2
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Decoded struct:", user2) // Output the decoded struct

	// ----------------------------------------------------------------------
	/// Conclusion:
	// - Use json.Marshal (and json.Unmarshal) when you have small, in-memory Go data structures
	//   and want to quickly convert them to/from JSON []byte or string.
	//   This is simple and efficient for small payloads or when you already have all the data in memory.
	//
	// - Use json.NewEncoder and json.NewDecoder when working with streams (like HTTP request/response bodies, files, or large data).
	//   This approach is more memory-efficient and scalable, as it processes data directly
	//   from/to an io.Reader or io.Writer without loading everything into memory.
	//
	// For most production environments—especially APIs, file processing, or any scenario
	// with large or streaming data—prefer the streaming approach with json.NewEncoder and json.NewDecoder.
}
