package main

import "fmt"

func main() {
	/*
		A closure in Go is a function value that references variables from outside its body.
		Useful in Go to maintain state across function calls
	*/
	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	// Anonymous function
	subtracter := func() func(int) int {
		count_down := 99
		return func(x int) int {
			count_down -= x
			return count_down
		}
	}()

	// Using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(5))
	fmt.Println(subtracter(7))
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
