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

	// Compare arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Array 1 is equal to array 2: ", array1 == array2)

	// Multidimensional array e.g. matrix
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	arr := [3]int{10, 20, 30}
	// var copied_arr *[3]int
	copied_arr := &arr // copied_arr is a pointer variable and not an array variable in this case
	copied_arr[0] = 15
	fmt.Println(arr)
	fmt.Println(copied_arr)

	fmt.Println("Arr is equal copied_arr: ", arr == *copied_arr)
}
