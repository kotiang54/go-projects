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
	fmt.Println("After sayHello()")

	go printNumbers()
	go printLetters()

	// Wait for a while to let goroutines finish
	time.Sleep(2 * time.Second)
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
