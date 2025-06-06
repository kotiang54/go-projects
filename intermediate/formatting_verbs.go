package main

import "fmt"

func main() {
	/*
		--- General formating verbs ---
		%v  Prints the value in the default format
		%#v Prints the value in Go-syntax format
		%T  Prints the type of the value
		%%  Prints the % sign
	*/
	i := 15 // float64 15.5

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i)

	// Strings
	str := "Hello World!"
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)

}
