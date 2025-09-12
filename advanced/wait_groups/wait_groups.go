package main

import (
	"fmt"
	"sync"
	"time"
)

// ***** Worker function without using channels
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// ***** Worker function using channels to return results
func workerChannels(id int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	// Send result back via channel to main function
	time.Sleep(time.Second)
	results <- id * 2 // Example result
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create a worker group and add workers
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers)

	// *** Start workers without channels ***
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("All workers completed.")

	fmt.Println("")

	// *** Start workers with channels ***
	results := make(chan int, numWorkers)
	wg.Add(numWorkers)

	// Start workers that send results back via channel
	for i := 1; i <= numWorkers; i++ {
		go workerChannels(i, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}

	fmt.Println("All workers with channels completed.")
}
