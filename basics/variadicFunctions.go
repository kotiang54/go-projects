package main

import "fmt"

func main() {
	// ... Ellipsis
	// func functionName(param1 type1, param2 type2, param3 ... type3) returnType {
	// 	functionBody
	// }

	total := sum(2, 3, 4, 5)
	fmt.Println(total)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
