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

	// Capturing Groups
	// Compile regex to capture data components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date := "2024-07-30"

	// Find all submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])
}
