package main

import "fmt"

// Structs are usually defined outside of main function - global scoping
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address // embedding concept
	PhoneHomeCell         // anonymous embedded field
}

type PhoneHomeCell struct {
	cell string
	home string
}

// Embedding - a struct within another struct
type Address struct {
	city    string
	country string
}

/*
Methods: functions associated with specific type
defined with a receiver - the struct type on which the method operates
  - value receivers
  - pointer receivers
*/
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// pointer receivers
func (p *Person) increamentAgeByOne() {
	p.age++
}

func main() {
	// Initializing a struct
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "UK",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "2349873543",
			cell: "3168738270",
		},
	}

	p2 := Person{
		firstName: "Kelvin",
		age:       25,
	}
	p2.address.city = "New York"
	p2.address.country = "USA"

	// Accessing the struct fields using dot notation
	fmt.Printf("Person #1 firstname: %s, lastname: %s, age: %d\n", p1.firstName, p1.lastName, p1.age)
	fmt.Printf("Person #1 home city: %s\n", p1.address.city)

	// Direct access to anonymous fields of struct
	fmt.Printf("Person #1 cell phone: %s\n", p1.cell)

	fmt.Printf("Person #2 firstname: %s, lastname: %s, age: %d\n", p2.firstName, p2.lastName, p2.age)
	fmt.Printf("Person #2 home country: %s\n", p2.address.country)

	// Method calls
	fmt.Println(p1.fullName())
	p2.increamentAgeByOne()
	fmt.Println(p2.age)

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
