package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	// Why use mutexes?
	// Safety: Mutexes provide a straightforward way to ensure that only one goroutine
	//     can access a piece of code or data at a time, preventing race conditions.
	// Simplicity: Mutexes are easy to understand and implement, making them a good choice
	//     for protecting shared resources in concurrent programming.
	// Flexibility: Mutexes can be used to protect complex data structures and critical sections
	//     of code, allowing for more granular control over concurrency.
	// Deadlock prevention: When used correctly, mutexes can help prevent deadlocks by ensuring
	//     that locks are acquired and released in a consistent order.
	// Read/Write Locks: Mutexes can be extended to read/write locks (RWMutex),
	//     allowing multiple readers but exclusive access for writers, improving performance
	//     in read-heavy scenarios.

	// Use cases:
	//     * Protecting shared resources (e.g., maps, slices)
	//     * Critical sections of code
	//     * Coordinating access to files or databases
	//     * Implementing thread-safe data structures
	//     * Ensuring consistency in complex operations
	//     * Synchronizing state changes in concurrent applications

	// This example shows how to use a mutex to safely increment a counter from multiple goroutines.

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	counter := &Counter{}

	numGoroutines := 10

	// wg.Add(numGoroutines)

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Count: %d\n", counter.getCount())

	// *** Demonstration without a struct (not recommended) ***
	// This example shows how to use a mutex to protect a simple integer variable.

	var (
		count int
		mu    sync.Mutex
	)

	numGoroutines = 10
	wg.Add(numGoroutines)

	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			count++
			mu.Unlock()
		}
	}

	for range numGoroutines {
		go increment()
	}

	wg.Wait()
	fmt.Printf("Final Count without struct: %d\n", count)
}
