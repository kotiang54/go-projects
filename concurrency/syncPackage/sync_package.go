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

// sync.Pool
// sync.Pool is a concurrency-safe pool of temporary objects that can be reused
// to reduce the overhead of allocating and deallocating memory frequently.
// It is particularly useful in high-performance applications where objects are
// created and destroyed frequently, as it helps to minimize garbage collection
// overhead and improve performance by reusing objects instead of creating new ones.

// Example usage of sync.Once and sync.Pool
type person struct {
	name string
	age  int
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

	fmt.Println("")

	// Example usage of sync.Pool
	personPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new person object.")
			return &person{}
		},
	}

	// Get a person object from the pool: Creates a new one since the pool is empty
	p1 := personPool.Get().(*person)
	p1.name = "Alice"
	p1.age = 24
	fmt.Println("Person 1:", p1)
	fmt.Printf("Person1 - Name: %s, Age: %d\n", p1.name, p1.age)

	// Return the person object to the pool
	personPool.Put(p1)
	fmt.Println("Returned Person 1 to the pool.")

	// Get another person object from the pool
	p2 := personPool.Get().(*person)
	fmt.Println("Person 2:", p2)                                 // This will reuse the object returned to the pool
	fmt.Printf("Person2 - Name: %s, Age: %d\n", p2.name, p2.age) // Should print Alice, 24

	p3 := personPool.Get().(*person)
	fmt.Println("Got another person:", p3) // The pool is empty, so it creates a new one with default values
}
