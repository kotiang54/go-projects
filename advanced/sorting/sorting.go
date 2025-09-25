package main

import (
	"fmt"
	"sort"
)

// Demonstrates sorting in Go using the sort package
type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	// Sorting integers
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Sorting strings
	stringSlice := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(stringSlice)
	fmt.Println("Sorted strings:", stringSlice)

	// sorting by functions
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	sort.Sort(ByAge(people))
	fmt.Println("People sorted by age:", people)
}
