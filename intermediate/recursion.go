package intermediate

import (
	"fmt"
)

func main() {
	fmt.Println("Factorial(5):", factorial(5))
	fmt.Println("Factorial(10):", factorial(10))

	fmt.Println("Sum of digit in 12:", sumOfDigits(12))
	fmt.Println("Sum of digit in 387:", sumOfDigits(387))

	fmt.Println("The 5th element in fibonacci:", fibonacci(5))
	fmt.Println("The 8th element in fibonacci:", fibonacci(8))
}

func factorial(n int) int {
	// This function does not consider many edge cases

	if n < 0 {
		// os.Exit(1)
		panic("Factorial is undefined for negative numbers")
	}
	// Base case:
	if n < 2 {
		return 1
	}

	// Recursive case: factorial n = n * factorial(n-1)
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// This function does not consider edge cases
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	// Base case:
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
