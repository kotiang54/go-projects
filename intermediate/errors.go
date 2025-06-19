package intermediate

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

// Custom error with struct
type customError struct {
	code    int
	message string
	er      error // wrapped error
}

// Error returns the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v\n", e.code, e.message, e.er)
}

// Function that returns a custom error
func doSomething() error {
	err := doSomethingElse()
	if err != nil {
		return &customError{
			code:    500,
			message: "Something went wrong!",
			er:      err,
		}
	}
	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
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

	err := doSomething()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Operation completed successfuly!")

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
