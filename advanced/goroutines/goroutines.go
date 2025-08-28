package main

import (
	"fmt"
	"time"
)

// Goroutines are lightweight threads managed by the Go runtime.
// They allow concurrent execution of functions and are non-blocking.

func main() {
	var err error

	go sayHello()
	// Wait for the goroutine to finish
	fmt.Println("After sayHello()")

	// Handle error from doSomething function
	// Note: In a real-world scenario, you would use channels or sync.WaitGroup to
	// synchronize goroutines and handle errors properly.
	// Here, we are simulating an error for demonstration purposes.
	go func() {
		err = doSomething()
	}()

	// err = go doSomething() This is not acceptable syntax in Go.
	go printNumbers()
	go printLetters()

	// Wait for a while to let goroutines finish
	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("doSomething() completed successfully.")
	}
	fmt.Println("Main function finished.")
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine!")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'a'; ch < 'e'; ch++ {
		fmt.Println(string(ch), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

// Error handling
func doSomething() error {
	// Simulate an error
	time.Sleep(3 * time.Second)
	return fmt.Errorf("an error occurred while doing something")
}
