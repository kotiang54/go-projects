package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex

	// This example demonstrates a deadlock situation.
	// Two goroutines attempt to lock two mutexes in opposite order,
	// leading to a deadlock where neither can proceed.
	go func() {
		mu1.Lock()
		defer mu1.Unlock() // LIFO unlocking

		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)

		mu2.Lock()
		defer mu2.Unlock()
		fmt.Println("Goroutine 1 locked mu2")
		fmt.Println("Goroutine 1 completed")
	}()

	// To remove the deadlock, ensure both goroutines lock the mutexes in the same order.
	// For example, both should lock mu1 first, then mu2.
	// Alternatively, use a timeout or try-lock mechanism to avoid waiting indefinitely.
	go func() {
		mu1.Lock()
		defer mu1.Unlock()

		fmt.Println("Goroutine 2 locked mu1")
		time.Sleep(time.Second)

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("Goroutine 2 locked mu2")
		fmt.Println("Goroutine 2 completed")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")

	// select {
	// // Block forever
	// }
}
