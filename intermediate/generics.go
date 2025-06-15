package main

import "fmt"

/*
	func swap[T any](a, b T) (T, T) {
		return b, a
	}
*/

type Stack[T any] struct {
	elements []T // slice of any type, but all elements must be of same type
}

func (s *Stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if s.isEmpty() {
		var zero T
		return zero, false
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func (s Stack[T]) printAll() {
	if s.isEmpty() {
		fmt.Println("The stack is empty!")
		return
	}
	fmt.Print("Stack elements:")
	for _, element := range s.elements {
		fmt.Print(" ", element)
	}
	fmt.Println()
}

func main() {
	/*
		x, y := 1, 2
		fmt.Printf("Before swap: x is %v and y is %v\n", x, y)
		x, y = swap(x, y)
		fmt.Printf("After swap: x is %v and y is %v\n", x, y)
	*/

	// Create an int stack
	arr := []int{1, 2, 3, 4, 5}
	intStack := Stack[int]{}

	for _, num := range arr {
		intStack.push(num)
	}

	intStack.printAll()
	if item, ok := intStack.pop(); ok {
		fmt.Printf("The popped item is %v\n", item)
	} else {
		fmt.Println("Something wrong happened!")
	}
	intStack.printAll()

	// Create a string stack
	fruits := []string{"Apple", "Orange", "Banana", "Mango", "Peach"}
	strStack := Stack[string]{}

	for _, fruit := range fruits {
		strStack.push(fruit)
	}

	strStack.printAll()
	if fruit, ok := strStack.pop(); ok {
		fmt.Printf("The popped item is %s\n", fruit)
	} else {
		fmt.Println("Something wrong happened!")
	}
	strStack.printAll()
}
