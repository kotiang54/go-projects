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
	// *** Demonstrating Mutex for Safe Concurrent Access ***
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
