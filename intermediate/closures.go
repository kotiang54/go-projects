package main

import "fmt"

func main() {
	/*
		Closures are function variables that access values outside their body.
		Persistent state in closures
	*/
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
}

// Function returning a function
func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
