package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	x, y := 1, 2
	fmt.Printf("Before swap: x is %v and y is %v\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("After swap: x is %v and y is %v\n", x, y)
}
