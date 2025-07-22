package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("User env var:", user)
	fmt.Println("Home env var:", home)

	err := os.Setenv("FRUIT", "banana")
	if err != nil {
		fmt.Println("Error setting environment variable FRUIT:", err)
		return
	}
	fmt.Println("FRUIT env var:", os.Getenv("FRUIT"))

	for _, elem := range os.Environ() {
		key_value := strings.SplitN(elem, "=", 2)
		fmt.Println("Key:", key_value[0])
		// fmt.Println("Value:", key_value[1])
	}

	err = os.Unsetenv("FRUIT")
	if err != nil {
		fmt.Println("Error unsetting environment variable FRUIT:", err)
		return
	}
	fmt.Println("FRUIT env var unset successfully")
	// Check if the variable is still set
	fmt.Println("FRUIT env var after unset:", os.Getenv("FRUIT"))
}
