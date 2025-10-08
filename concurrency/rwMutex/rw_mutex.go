package main

import (
	"fmt"
	"sync"
)

// What are RWMutexes?
// RWMutexes (Read-Write Mutexes) are synchronization primitives that allow
// multiple readers or a single writer to access a shared resource concurrently.
// They are useful in scenarios where read operations are more frequent than write operations,
// as they allow multiple goroutines to read the data simultaneously while ensuring exclusive access for write operations.

var (
	rwmu    sync.RWMutex
	counter int
)

// readCounter reads the value of counter with a read lock
func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Println("Read Counter:", counter)
}

// writeCounter writes a new value to counter with a write lock
func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	defer rwmu.Unlock()
	counter = value
	fmt.Println("Wrote Counter:", counter)
}

func main() {
	// Example usage of RWMutex
	var wg sync.WaitGroup

	// Start multiple readers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readCounter(&wg)
	}

	// Start a writer
	wg.Add(1)
	// time.Sleep(time.Second) // Ensure some reads happen before write
	go writeCounter(&wg, 15)

	wg.Wait()
}
