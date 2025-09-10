package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, tasks <-chan int, result chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d procesing task %d\n", id, task)
		// Simulate work
		time.Sleep(time.Second)
		result <- task * 2 // Example processing
	}
}

func main() {
	// Tasks, Workers, Task Queue
	numWorkers := 3
	numJobs := 10
	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create workers
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, results)
	}

	// Send values to the task channel
	for j := 0; j < numJobs; j++ {
		tasks <- j
	}
	close(tasks)

}
