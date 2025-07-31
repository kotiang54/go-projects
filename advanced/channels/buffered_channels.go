package main

import (
	"fmt"
	"time"
)

// Buffered channels allow sending and receiving values without blocking,
// up to a specified capacity. They are useful for decoupling the sender and receiver.
func main() {

	// channelName := make(chan chanType, capacity) // Create a buffered channel with capacity 3
	bufferedChannel := make(chan int, 3)
	bufferedChannel <- 1
	bufferedChannel <- 2
	bufferedChannel <- 3

	go func() {
		time.Sleep((1 * time.Second))
		// Receiving from the buffered channel
		fmt.Println("Buffered Channel Values:", <-bufferedChannel)
		fmt.Println("Buffered Channel Values:", <-bufferedChannel)
	}()

	bufferedChannel <- 4 // This will block if the channel is full
	fmt.Println("Buffered Channels!")
}
