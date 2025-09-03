package main

import (
	"context"
	"fmt"
	"time"
)

type contextKey string

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}

}

func main() {
	// Using context.TODO and context.Background
	// Both are empty contexts, but TODO is used when you are unsure which to use
	// Background is typically used at the top level of a program
	// and is never canceled, has no values, and has no deadline.
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, contextKey("fullname"), "John Doe")
	fmt.Println(ctx)
	fmt.Println(ctx.Value(contextKey("fullname")))

	ctx = context.WithValue(contextBkg, contextKey("city"), "New York")
	fmt.Println(ctx)
	fmt.Println(ctx.Value(contextKey("city")))

	fmt.Println("")

	// Using context with a function
	ctx = context.TODO()
	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result with context.WithTimeout():", result)

	time.Sleep(3 * time.Second) // Simulate a delay to trigger the timeout
	// Check the result after the timeout
	// The context will be done, and the function should return "Operation cancelled"
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result after timeout:", result)

}
