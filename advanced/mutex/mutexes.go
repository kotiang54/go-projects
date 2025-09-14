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
	var wg sync.WaitGroup
	counter := &Counter{}

	numGoroutines := 10

	wg.Add(numGoroutines)

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
}
