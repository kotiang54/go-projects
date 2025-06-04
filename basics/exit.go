package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Exit() exits control flow without perfoming any cleanup operations
	// such as defer statements

	defer fmt.Println("Deferred statement")
	fmt.Println("Starting the main function")

	// Exit with status code 1
	os.Exit(1)

	// This line will never be executed
	fmt.Println("End of main function")
}
