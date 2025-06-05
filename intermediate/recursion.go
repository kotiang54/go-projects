package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	// Base case:
	if n < 2 {
		return 1
	}

	// Recursive case: factorial n = n * factorial(n-1)
	return n * factorial(n-1)
}
