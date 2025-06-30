package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convert string numbers to integers of type int
	numStr := "1234"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", num)
	fmt.Println("Parsed Integer plus 1:", num+1)

	// convert string to interger of type int32 or int64
	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", numistr)

	// parse floats
	floatstr := "3.142"
	floatVal, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Println("Parsed Float:", floatVal)
}
