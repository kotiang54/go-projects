package basics

import "fmt"

func main() {
	/*	... Ellipsis
		func functionName(param1 type1, param2 type2, param3 ... type3) returnType {
			functionBody
		}
	*/

	// NOTE: Place the variadic parameters at the tail of the function input arguments

	total := sum(2, 3, 4, 5)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4, 5, 9}
	fmt.Println("Sum of slice arguments:", sum(numbers...))
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
