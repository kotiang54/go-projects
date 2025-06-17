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

func process(data []byte) error {
	if len(data) == 0 {
		return errors.New("Empty data!")
	}
	// Process data
	return nil
}

type myError struct {
	message string
}

// Custom error
func (m *myError) Error() string {
	return fmt.Sprintf("Error: %s", m.message)
}

func eprocess() error {
	return &myError{"Custom error message"}
}

// error handling
func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("config error")
}

func main() {
	/*
		// positive number
		result, err := sqrt(16)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)

		// negative number
		result, err = sqrt(-16)
		if err != nil {Error() string
		fmt.Println(result)
	*/

	data := []byte{}
	if err := process(data); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data processed succesfully")

	// error interface of built-in package
	if err1 := eprocess(); err1 != nil {
		fmt.Println(err1)
		return
	}

	if err := readData(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data read succesfully")
}
