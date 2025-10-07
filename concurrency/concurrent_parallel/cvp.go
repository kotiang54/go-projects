package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Concurrency versus parallelism

// Concurrency is about managing multiple tasks at once and not necessarily doing them at the same time
// Parallelism is about executing multiple tasks simultaneously

// Function to print numbers with a delay
func printNumbers() {
	for i := range 5 {
		fmt.Println(time.Now())
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Function to print letters with a delay
func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(time.Now())
		fmt.Println(string(letter))
		time.Sleep(500 * time.Millisecond)
	}
}

func heavyTask(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d starting\n", id)

	for range 100_000_000 {
		// Simulate heavy computation
	}
	fmt.Println(time.Now())
	fmt.Printf("Task %d completed\n", id)
}

func main() {

	// Example of parallel execution using goroutines
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second) // wait for goroutine to return

	// Example of concurrent parallel execution
	numThreads := 4

	runtime.GOMAXPROCS(numThreads) // Set the number of OS threads to use
	var wg sync.WaitGroup

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go heavyTask(i, &wg)
	}
	wg.Wait()
}
