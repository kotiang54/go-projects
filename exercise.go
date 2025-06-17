package main

import "fmt"

func modifyValue(ptr *int) {
	*ptr = 20
}

func main() {
	value := 10
	ptr := &value

	fmt.Println("Before modification:", value)
	modifyValue(ptr)
	fmt.Println("After modification:", value)
}
