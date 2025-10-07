package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var mu1, mu2 sync.Mutex

	// This example demonstrates a deadlock situation.
	// Two goroutines attempt to lock two mutexes in opposite order,
	// leading to a deadlock where neither can proceed.
	go func() {
		mu1.Lock()
		defer mu1.Unlock() // LIFO unlocking
		defer mu2.Unlock()

		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()
		defer mu1.Unlock()

		fmt.Println("Goroutine 2 locked mu2")
		time.Sleep(time.Second)
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu1")
	}()

	// time.Sleep(3 * time.Second)
	// fmt.Println("Main function completed")

	select {
	// Block forever
	}
}
