package main

import "fmt"

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
	slice := make([]int, 5)
	slice = nums2[1:4]
	fmt.Println(slice)

	slice = append(slice, 6, 0)
	fmt.Println(slice)
}
