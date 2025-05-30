package main

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

	// Access a slice
	fmt.Println("Element at index 3 is", slice_copy[3])

	// modify a slice value
	// slice_copy[3] = 50
	// fmt.Println("Element at index 3 after update is", slice_copy[3])

	if slices.Equal(slice1, slice_copy) {
		fmt.Println("slice1 is equal to slice_copy!")
	} else {
		fmt.Println("slice1 is not equal slice_copy!")
	}

}
