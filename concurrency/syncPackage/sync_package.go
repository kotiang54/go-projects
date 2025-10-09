package main

import (
	"fmt"
	"sync"
)

// What are condition variables?
// Condition variables are synchronization primitives that allow goroutines
// to wait for certain conditions to be met. They are often used in conjunction
// with mutexes to avoid busy waiting and to signal other goroutines when a condition has changed.

const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{items: make([]int, 0, size)}
	b.cond = sync.NewCond(&b.mu) // Initialize the condition variable with the buffer's mutex
	return b
}

// Produce adds an item to the buffer if there is space
// Otherwise, it waits until space is available
func (b *buffer) produce(item int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == bufferSize {
		b.cond.Wait() // Wait until there is space in the buffer
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal() // Signal that an item has been added
}

// Consume removes and returns an item from the buffer
// If the buffer is empty, it waits until an item is available
func (b *buffer) consume() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	for len(b.items) == 0 {
		b.cond.Wait() // Wait until there is an item to consume
	}

	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("Consumed:", item)
	b.cond.Signal() // Signal that an item has been removed
	return item
}

func main() {
	// Example usage of condition variables
}
