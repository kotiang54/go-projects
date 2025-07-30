package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They allow concurrent execution of functions.

func main() {
	go sayHello()
	// Wait for the goroutine to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Main function finished.")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine!")
}
