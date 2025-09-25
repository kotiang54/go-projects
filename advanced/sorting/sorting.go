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

type By func(p1, p2 *Person) bool

type personSorter struct {
	people []Person
	by     func(p1, p2 *Person) bool
}

func (s *personSorter) Len() int {
	return len(s.people)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.people[i], &s.people[j])
}

func (s *personSorter) Swap(i, j int) {
	s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By) Sort(people []Person) {
	ps := &personSorter{
		people: people,
		by:     by,
	}
	sort.Sort(ps)
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

	// Sorting strings
	stringSlice := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(stringSlice)
	fmt.Println("Sorted strings:", stringSlice)

	// sorting by functions
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"Bobby", 22},
		{"David", 30},
	}

	// Sort people by age using custom sort
	// sort.Sort(ByAge(people))
	// fmt.Println("People sorted by age:", people)

	// Sort people by name using custom sort
	// sort.Sort(ByName(people))
	// fmt.Println("People sorted by name:", people)

	// Sort people by age using a custom function
	// Define the comparison function
	name := func(p1, p2 *Person) bool {
		return p1.Name < p2.Name
	}

	By(name).Sort(people)
	fmt.Println("People sorted by name using custom function:", people)
}
