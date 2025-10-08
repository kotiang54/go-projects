package main

import (
	"fmt"
	"sync"
)

var (
	rwmu    sync.RWMutex
	counter int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	defer rwmu.RUnlock()
	fmt.Println("Read Counter:", counter)
}

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
	go writeCounter(&wg, 42)

	wg.Wait()
}
