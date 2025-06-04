package main

import "fmt"

func main() {
	process()
}

func process() {
	/* The deferred function is will be deferred till the end of the surrounding
	   functions have executed.
	*/
	defer fmt.Println("Deferred statement executed")
	fmt.Println("Normal execution statement")
}
