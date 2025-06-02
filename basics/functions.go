package main

import "fmt"

func main() {
	// func <name>(parameters list) returnType {
	// Code block
	// return value
	// }

	sum := add(1, 2)
	fmt.Println("Sum total", sum)
}

func add(a, b int) int {
	return a + b
}
