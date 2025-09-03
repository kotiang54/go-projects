package main

import (
	"context"
	"fmt"
	"log"
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

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func logWithContext(ctx context.Context, message string) {
	if requestID, ok := ctx.Value(contextKey("requestID")).(string); ok {
		log.Printf("Request ID: %s - %s\n", requestID, message)
	} else {
		log.Println("No Request ID found in context")
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

	// Using context with a
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

	fmt.Println("")

	rootCtx := context.Background()
	// Create a context with a timeout of 2 seconds
	// This context will be automatically cancelled after 2 seconds
	// Uncomment the following lines to use a timeout instead of manual cancellation

	// ctx, cancel = context.WithTimeout(rootCtx, 2*time.Second)
	// defer cancel()

	// Manually cancel the context (if not already cancelled)
	ctx, cancel = context.WithCancel(rootCtx)

	go func() {
		time.Sleep(2 * time.Second) // simulate heavy time consuming operation
		cancel()                    // manually cancel the context
	}()

	ctx = context.WithValue(ctx, contextKey("requestID"), "Jhtgfsr7353425232")

	// Start a goroutine that does some work
	go doWork(ctx)
	time.Sleep(3 * time.Second) // Let the work run for a while

	// After 2 seconds, the context will be cancelled due to the timeout
	// The doWork function should print "Work cancelled: context deadline exceeded"
	// The requestID should still be accessible
	requestID := ctx.Value(contextKey("requestID"))
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No Request ID found")
	}

	// Logging with context
	logWithContext(ctx, "This is a log message with context")
}
