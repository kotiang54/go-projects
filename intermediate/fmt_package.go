package main

import "fmt"

func main() {
	// fmt.Print
	fmt.Print("Hello ")
	fmt.Print("World! ")
	fmt.Print("13, 1999")

	// fmt.Println
	fmt.Println("")
	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println("13, 1999")

	// fmt.Printf
	name := "John"
	age := 34
	fmt.Println("")
	fmt.Printf("Name: %s, age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %#X\n", age, age)

	// Formating functions
	fmt.Println("")
	s := fmt.Sprint("Hello ", "World! ", 123, 456)
	fmt.Println(s)
}
