package main

import (
	"fmt"
	"time"
)

// Buffered channels allow sending and receiving values without blocking,
// up to a specified capacity. They are useful for decoupling the sender and receiver.

// func main() {
// 	// ************* Blocking on RECEIVE Only if buffer is empty **********
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 	}()

// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of program!")
// }

func main() {
	// ************* Blocking on SEND Only if buffer is FULL ***************

	// channelName := make(chan chanType, capacity) // Create a buffered channel with capacity 3
	bufferedChannel := make(chan int, 3)
	bufferedChannel <- 1
	bufferedChannel <- 2
	// bufferedChannel <- 3
	fmt.Println("Receiving from buffer!")

	go func() {
		fmt.Println("Goroutine 2 seconds timer started!")
		time.Sleep((4 * time.Second))
		// Receiving from the buffered channel
		fmt.Println("Buffered Channel Values:", <-bufferedChannel)
		fmt.Println("Buffered Channel Values:", <-bufferedChannel)
	}()

	fmt.Println("Blocking starts!")
	bufferedChannel <- 4 // This will block if the buffer is full
	// bufferedChannel <- 5
	// bufferedChannel <- 6
	fmt.Println("Buffered Channels!")
}
