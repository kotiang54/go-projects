// Package main demonstrates JSON encoding and decoding in Go using structs with named and anonymous embedding.
// It defines Person, Address, and Employee types with appropriate JSON struct tags, and shows how to marshal and unmarshal
// these types to and from JSON. The program also illustrates handling arrays of structs and decoding unknown JSON structures
// into a map[string]interface{}. Key features include:
//   - Named struct embedding in Person for Address.
//   - Anonymous struct embedding in Employee for Address.
//   - Use of `omitempty` to skip zero-value fields during marshaling.
//   - Marshaling structs and slices to JSON.
//   - Unmarshaling JSON strings into structs and generic maps.
//   - Accessing embedded struct fields directly from the parent struct.
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"address"` // named struct embedding
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Employee struct {
	FullName string           `json:"full_name"`
	Age      int              `json:"age,omitempty"`
	EmpID    string           `json:"emp_id"`
	Email    string           `json:"email"`
	Address  `json:"address"` // anonymous struct embedding
}

func main() {

	person := Person{
		Name: "John Doe",
		// Age:   30,
		Email: "john.doe@example.com",
	}

	// The person variable can now be used to marshal to JSON or perform other operations
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Example of another person with all fields filled
	person1 := Person{
		Name:    "Jane Smith",
		Age:     25,
		Email:   "jane.smith@example.com",
		Address: Address{City: "New York", State: "NY"},
	}

	// Marshal the second person to JSON
	jsonData1, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling person1 to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData1))

	// Decode or unMarshal JSON back to a Person struct
	jsonData2 := `{"full_name": "Alice Johnson", "age": 28, "emp_id": "0007", "email": "alice.johnson@company.com", "address": {"city": "Los Angeles", "state": "CA"}}`

	var employeeFromJSON Employee
	// Unmarshal the JSON string into the Employee struct
	err = json.Unmarshal([]byte(jsonData2), &employeeFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON to Employee struct:", err)
		return
	}
	fmt.Println("Employee struct from JSON:", employeeFromJSON)

	fmt.Println("Jane's age increased by 5 years:", employeeFromJSON.Age+5)
	fmt.Println("Jane's city:", employeeFromJSON.City)

	// Decoding arrays
	cityStateList := []Address{
		{City: "San Francisco", State: "CA"},
		{City: "Austin", State: "TX"},
		{City: "Seattle", State: "WA"},
		{City: "Miami", State: "FL"},
	}

	fmt.Println("List of City and State:", cityStateList)
	jsonList, err := json.Marshal(cityStateList)
	if err != nil {
		fmt.Println("Error marshalling list of cities to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonList))

	// Handling unknown json structures
	jsonData3 := `{"name": "Bob", "age": 35, "email": "bob@example.com", "address": {"city": "Chicago", "state": "IL"}}`
	var data map[string]interface{}

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Decoded JSON data:", data)
}
