package main

import "fmt"

func main() {
	// Switch statement (switch, case, default) (fallthrough)
	// switch expression {
	// case value1:
	// 	// Code to be executed if expression equals value1
	// case value2:
	// 	// Code to be executed if expresion equals value2
	// case value3:
	// 	// Code to be executed if expression equals value3
	// default:
	// 	// Code to be executed if expression does not match any value
	// }

	fruit := "banana"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	case "orange":
		fmt.Println("It's an orange")
	default:
		fmt.Println("Unknown fruit")
	}
}
