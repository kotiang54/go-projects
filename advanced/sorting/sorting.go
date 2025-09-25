package main

import (
	"fmt"
	"sort"
)

// Demonstrates sorting in Go using the sort package
// Custom sorting by implementing sort.Interface
type Person struct {
	Name string
	Age  int
}

// By is the type of a "less" function that defines
// the ordering of its Person arguments.
type By func(p1, p2 *Person) bool

// personSorter joins a By function and a slice of People to be sorted.
type personSorter struct {
	people []Person
	by     By // func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

// Sort sorts the argument slice according to the function by.
func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	// sort.Sort(ps)
	sort.Stable(ps) // Use stable sort to maintain order of equal elements
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person
type ByName []Person

// Len, Less, and Swap methods are required by sort.Interface
func (a ByAge) Len() int {
	return len(a)
}

func (a ByName) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	// Sorting integers
	numbers := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Sorting strings by last character using custom sort.Interface
	stringSlice := []string{"banana", "apple", "cherry", "date"}
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})

	fmt.Println("Sorted strings by last character:", stringSlice)

	// sorting by functions
	people := []Person{
		{"Alice", 30},
		{"Bob", 27},
		{"Bob", 25},
		{"Charlie", 35},
		{"Bobby", 22},
		{"David", 30},
	}

	// Sort people by age using custom sort
	// sort.Sort(ByAge(people)) // Uncomment this line to sort people by age using ByAge
	// fmt.Println("People sorted by age:", people)

	// Sort people by name using custom sort
	// Uncommenting this section to demonstrate sorting by name using ByName
	sort.Sort(ByName(people))
	fmt.Println("People sorted by name:", people)

	// Sort people by age using a custom function
	// Define the comparison function
	compareByName := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	// The By type is a function type that implements sort.Interface by defining
	// a custom Less method. Here, we use it to sort the people slice by name.
	By(compareByName).Sort(people)
	fmt.Println("People sorted by name using custom function:", people)
}
