package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// Simulate work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create a worker group
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed.")
}
