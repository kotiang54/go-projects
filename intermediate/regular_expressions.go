package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\".")
	fmt.Println(`He said, "I an great".`)

	// Compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test strings
	email1 := "user@email.com"
	email2 := "invalid_email"
	email3 := "jdoe2025@shockers.wichita.edu"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))
	fmt.Println("Email3:", re.MatchString(email3))
}
