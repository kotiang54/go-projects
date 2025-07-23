package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"`
	Email   string  `json:"email"`
	Address Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
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
}
