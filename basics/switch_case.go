package basics

import "fmt"

func main() {
	/*  // Switch statement (switch, case, default) (fallthrough)
	switch expression {
	case value1:
		// Code to be executed if expression equals value1
	case value2:
		// Code to be executed if expresion equals value2
	case value3:
		// Code to be executed if expression equals value3
	default:
		// Code to be executed if expression does not match any value
	}
	*/

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

	// Multiple Conditions
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day")
	}

	number := 17

	switch {
	case number < 10:
		fmt.Println("Number is less than 10.")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19.")
	default:
		fmt.Println("Number is 20 or more.")
	}

	// Fallthrough expression
	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not a Two")
	}

	// Switch-case for type assertion
	checkType(10)
	checkType(3.14257)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's a float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
