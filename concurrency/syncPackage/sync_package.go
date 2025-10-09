package main

import (
	"fmt"
	"sync"
	"time"
)

// What are condition variables?
// Condition variables are synchronization primitives that allow goroutines
// to wait for certain conditions to be met. They are often used in conjunction
// with mutexes to avoid busy waiting and to signal other goroutines when a condition has changed.

// sync.NewCond example:
const bufferSize = 5

type buffer struct {
	items []int
	mu    sync.Mutex
	cond  *sync.Cond
}

// newBuffer initializes a new buffer with a given size
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
		b.cond.Wait() // Wait until there is space in the buffe
	}

	b.items = append(b.items, item)
	fmt.Println("Produced:", item)
	b.cond.Signal() // Signal that an item has been added
}

// Consume removes and returns the first item from the buffer in a thread-safe manner.
// If the buffer is empty, it waits until an item becomes available.
// After consuming an item, it signals any waiting goroutines that space is available.
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

// Example producer and consumer functions
func producer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		b.produce(i + 100)
		time.Sleep(100 * time.Millisecond) // Simulate time taken to produce an item
	}
}

func consumer(b *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		b.consume()
		time.Sleep(200 * time.Millisecond) // Simulate time taken to consume an item
	}
}

// sync.Once example:

var once sync.Once

func initialize() {
	fmt.Println("This function will be called only once even if called multiple times.")
}

func main() {
	// Example usage of condition variables
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
	fmt.Println("All producers and consumers have finished.")
	fmt.Println("")

	// Example usage of sync.Once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Goroutine # ", i)
			once.Do(initialize) // initialize() will be called only once
		}(i)
	}
	wg.Wait()
}
