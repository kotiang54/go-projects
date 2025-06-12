package main

import "fmt"

func main() {
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

	// Accessing the struct fields using dot notation
	fmt.Printf(" Person #1 firstname: %s, lastname: %s, age: %d\n", p1.firstName, p1.lastName, p1.age)
	fmt.Printf(" Person #2 firstname: %s, lastname: %s, age: %d\n", p2.firstName, p2.lastName, p2.age)
	fmt.Println(p1.fullName())

	// Anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "userjDoe",
		email:    "johndoe89@examples.org",
	}

	fmt.Println(user.email)
}

// Structs are usually defined outside of main function
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Methods: functions associated with specific type
// defined with a receiver - the struct type on which the method operates
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}
