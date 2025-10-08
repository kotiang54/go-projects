package main

import (
	"fmt"
	"sync"
	"time"
)

// What are deadlocks?
// A deadlock is a situation in concurrent programming where two or more goroutines
// are unable to proceed because each is waiting for the other to release a resource.
// This results in a standstill where none of the goroutines can make progress.
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

	// Deadlock occurs here because:
	// 1. Goroutine 1 locks mu1 and waits for mu2.
	// 2. Goroutine 2 locks mu2 and waits for mu1.

	// This creates a circular wait condition, leading to a deadlock.

	go func() {
		mu2.Lock()
		defer mu2.Unlock() // LIFO unlocking

		fmt.Println("Goroutine 2 locked mu2")
		time.Sleep(time.Second)

		mu1.Lock()
		defer mu1.Unlock()
		fmt.Println("Goroutine 2 locked mu1")
		fmt.Println("Goroutine 2 completed")
	}()

	// To remove the deadlock, ensure both goroutines lock the mutexes in the same order.
	// For example, both should lock mu1 first, then mu2.
	// Alternatively, use a timeout or try-lock mechanism to avoid waiting indefinitely.
	// Uncommenting the following code will prevent the deadlock by locking in the same order.

	// go func() {
	// 	mu1.Lock()
	// 	defer mu1.Unlock()

	// 	fmt.Println("Goroutine 2 locked mu1")
	// 	time.Sleep(time.Second)

	// 	mu2.Lock()
	// 	defer mu2.Unlock()

	// 	fmt.Println("Goroutine 2 locked mu2")
	// 	fmt.Println("Goroutine 2 completed")
	// }()

	// time.Sleep(3 * time.Second)
	// fmt.Println("Main function completed")

	// Prevent the main function from exiting immediately
	// In a real application, use sync.WaitGroup or other synchronization methods
	// to wait for goroutines to finish.
	select {
	// Block forever
	}
}
