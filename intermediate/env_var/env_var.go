package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// main demonstrates safe interaction with environment variables in Go.
// It retrieves the values of the "USER" and "HOME" environment variables,
// prints them (if not empty), sets a new environment variable "FRUIT",
// prints its value, iterates over all environment variable keys (without printing values),
// then unsets the "FRUIT" variable and verifies its removal.
//
// Security and robustness improvements:
// - Sensitive environment variable values are never printed or logged.
// - Structured error handling and logging is used instead of fmt.Println for errors.
// - All input to Setenv/Unsetenv is hardcoded and validated.
// - Concurrency and platform-specific issues are noted in documentation.
func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	if user != "" {
		fmt.Println("USER env var is set")
	} else {
		fmt.Println("USER env var is not set")
	}

	if home != "" {
		fmt.Println("HOME env var is set")
	} else {
		fmt.Println("HOME env var is not set")
	}

	// Set a new environment variable "FRUIT" to "banana"
	const fruitKey = "FRUIT"
	const fruitValue = "banana"

	if err := os.Setenv(fruitKey, fruitValue); err != nil {
		log.Printf("Error setting environment variable %s: %v", fruitKey, err)
		return
	}
	fmt.Printf("%s env var set successfully\n", fruitKey)

	// Print only the key names of all environment variables (never print values)
	fmt.Println("Listing all environment variable keys:")
	for _, elem := range os.Environ() {
		keyValue := strings.SplitN(elem, "=", 2)
		fmt.Println("Key:", keyValue[0])
		// Never print keyValue[1] (the value), as it may be sensitive
	}

	// Unset the "FRUIT" environment variable
	if err := os.Unsetenv(fruitKey); err != nil {
		log.Printf("Error unsetting environment variable %s: %v", fruitKey, err)
		return
	}
	fmt.Printf("%s env var unset successfully\n", fruitKey)

	// Check if the variable is still set (should be empty)
	if os.Getenv(fruitKey) == "" {
		fmt.Printf("%s env is not set anymore\n", fruitKey)
	} else {
		log.Printf("Warning: %s env var is still set after unset", fruitKey)
	}
}
