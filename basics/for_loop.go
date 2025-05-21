package main

import "fmt"

func main() {

	// Simple iteration over range
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	// Iterate over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // continue the loop but skip the rest of the line/statements
		}
		fmt.Println("Odd number: ", i)
		if i == 5 {
			break // break out of the loop
		}
	}

	rows := 5

	// Outer loop
	for i := 1; i <= rows; i++ {
		// inner llop for spaces
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		// inner loop for stars
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
