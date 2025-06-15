package main

import (
	"errors"
	"fmt"
)

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("mathError: square root of a negative number")
	}
	// commputer the square root
	return 1, nil
}

func main() {

	// positive number
	result, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

	// negative number
	result, err = sqrt(-16)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
