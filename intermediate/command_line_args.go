package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Args and flag packages

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument", i, ":", arg)
	}

	// flag
	var (
		name string
		age  int
		male bool
	)

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)
}
