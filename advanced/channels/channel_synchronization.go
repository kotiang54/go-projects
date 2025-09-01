package main

import (
	"fmt"
	"strconv"
	"time"
)

func synchronization() {

	ch := make(chan int)

	// Producer goroutine sending data to dataCh
	go func() {
		fmt.Println("Sending ... ")
		ch <- 9 // Blocking until the value is received
		// time.Sleep(1 * time.Second)
		fmt.Println("Sent value 9")
	}()

	value := <-ch // Blocking until a value is sent
	fmt.Println("Received value:", value)

	// ***** Synchronizing multiple goroutines using channels *****
	// This example launches multiple goroutines that perform work and signal completion via a channel.
	numGoroutines := 3
	done := make(chan int, 3)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d working...\n", id)
			time.Sleep(time.Second)
			done <- id
		}(i)
	}

	for i := 0; i < numGoroutines; i++ {
		<-done // Wait for each Goroutine to finish
	}

	// All goroutines have completed
	fmt.Println("All goroutines completed")

	// ***** Synchronize Data Exchange Between Goroutines *****
	dataCh := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			dataCh <- "Hello " + strconv.Itoa(i)
			time.Sleep(time.Millisecond * 100)
		}
		close(dataCh) // Close the channel after sending all data
	}()

	for value := range dataCh {
		// Continuously receive data simulating a blocking operation
		// Loops until the channel is closed
		fmt.Println("Received:", value, ":", time.Now())
	}
}
