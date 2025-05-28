package main

import "fmt"

func main() {
	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 13
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	// array of string values
	fruits := [4]string{"Apples", "Bananas", "Oranges", "Pinapples"}
	fmt.Println("Fruits array: ", fruits)

	// Traversing an array
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index, ", i, ": ", fruits[i])
	}

	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
