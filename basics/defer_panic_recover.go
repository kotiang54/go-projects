package basics

import "fmt"

func main() {
	// defer statements
	process_defer(10)
	fmt.Println("")

	// panic(interface{})
	// panic stops ordinary flow of control
	// process_panic(-3)

	// recover is used to regain control of goroutine
	// it stops propagation of panic and regain control of flow
	process_recover()
	fmt.Println("Returned from recover process!")
}

func process_defer(i int) {
	/* The deferred function will be deferred till the end of the surrounding
	   functions have executed.
	   Multiple defer statements are executed in LIFO manner
	*/
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}

func process_panic(input int) {
	defer fmt.Println("Deferred statement after panic!")
	if input < 0 {
		panic("Input must be a non-negative number!")
	}
	fmt.Println("Processing input:", input)
}

func process_recover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start process!")
	panic("Something went wrong!")
}
