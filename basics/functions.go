package basics

import (
	"errors"
	"fmt"
)

func main() {
	/* func functionName(parameters list) returnType {
		  // Code block
		  return value
	   }

		// functions with multiple return types
		func functionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...) { .
			//Code block
			return value1, value2
		}
	*/

	sum := add(1, 2)
	fmt.Println("Sum total", sum)
	fmt.Println("Sum total", add(3, 4))

	// Anonymous functions / closures / function literals:
	// functions defined without a name directly inline where they are used
	greet := func() {
		fmt.Println("Hello Anonymous Function")
	}
	greet()

	// function as types
	operation := add
	result := operation(3, 5)
	fmt.Println(result)

	// first-class citizen/object
	// 1: pass a function as an argument
	result = applyOperation(5, 5, add)
	fmt.Println("5 + 5 =", result)

	// 2: return and use a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 x 2 =", multiplyBy2(6))

	// Example multiple returnType function
	q, rem := divide(10, 3)
	fmt.Printf("Quotient: %d. Remainder: %d\n", q, rem)

	result1, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result1)
	}
}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Multiple returnType function
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater!")
	}
}
