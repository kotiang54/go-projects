package main

import (
	"fmt"
	"sync"
	"time"
)

// Struct to hold worker results
type WorkerResult struct {
	WorkerID int
	Task     int
	Result   int
}

// ***** Worker function without using channels
// wg is used to signal when the worker is done
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Simulate work
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

// ***** Worker function using channels to return results
// results channel is used to send data back to main function
func workerChannels(id int, tasks <-chan int, results chan<- WorkerResult, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// Send result back via channel to main function
	for task := range tasks {
		// Simulate work per task
		time.Sleep(time.Second)
		results <- WorkerResult{WorkerID: id, Task: task, Result: task * 2} // Example result
	}
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// Create a worker group and add workers
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	// *** Start workers without channels ***
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("All workers completed.")

	fmt.Println("")

	// *** Start workers with channels ***
	results := make(chan WorkerResult, numWorkers)
	wg.Add(numWorkers)

	// Send tasks to workers via channel
	for i := 1; i <= numJobs; i++ {
		tasks <- i
	}
	close(tasks) // Close tasks channel after sending all tasks

	// Start workers that send results back via channel
	for i := 1; i <= numWorkers; i++ {
		go workerChannels(i, tasks, results, &wg)
	}

	// Close results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results from channel
	for result := range results {
		fmt.Printf("Worker %d processed task %d with result %d\n", result.WorkerID, result.Task, result.Result)
	}

	fmt.Println("All workers with channels completed.")
}
