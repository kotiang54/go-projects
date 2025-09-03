package main

import (
	"context"
	"fmt"
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

	ctx1 := context.WithValue(contextBkg, contextKey("city"), "New York")
	fmt.Println(ctx1)
	fmt.Println(ctx1.Value(contextKey("city")))
}
