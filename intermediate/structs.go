package main

import "fmt"

func main() {

	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	// Initializing a struct
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	p2 := Person{
		firstName: "Kelvin",
		age:       25,
	}

	// Accessing the struct fields
	fmt.Printf(" Person #1 firstname: %s, lastname: %s, age: %d\n", p1.firstName, p1.lastName, p1.age)
	fmt.Printf(" Person #2 firstname: %s, lastname: %s, age: %d\n", p2.firstName, p2.lastName, p2.age)
}
