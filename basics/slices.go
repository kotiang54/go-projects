package basics

import (
	"fmt"
	"slices"
)

func main() {
	// Unlike arrays slices have no fixed length
	// var sliceName []elementType

	var nums []int
	fmt.Println(nums)

	var nums1 = []int{1, 2, 3}
	fmt.Println(nums1)

	nums2 := []int{7, 9, 8, 2, 5, 3, 1}
	fmt.Println(nums2)

	// Use make function to create a slice
	slice1 := make([]int, 5)
	slice1 = nums2[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 0)
	fmt.Println("Slice: ", slice1)

	// make copy of a slice
	slice_copy := make([]int, len(slice1))
	copy(slice_copy, slice1)

	fmt.Println("Copy slice: ", slice_copy)

	for i, v := range slice_copy {
		fmt.Println(i, ":", v)
	}

	// Access a slice element
	fmt.Println("Element at index 3 is", slice_copy[3])

	// modify a slice value
	// slice_copy[3] = 50
	// fmt.Println("Element at index 3 after update is", slice_copy[3])

	if slices.Equal(slice1, slice_copy) {
		fmt.Println("slice1 is equal to slice_copy!")
	} else {
		fmt.Println("slice1 is not equal slice_copy!")
	}

	// Multidimensional slice
	two_D_slice := make([][]int, 3)
	counter := 0

	for i := 0; i < 3; i++ {
		two_D_slice[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			two_D_slice[i][j] = counter
			counter++
		}
	}

	fmt.Println(two_D_slice)

	// Length and capacity of a slice
	slice2 := slice1[2:4]
	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice2: ", slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The length of slice2 is", len(slice2))
}
