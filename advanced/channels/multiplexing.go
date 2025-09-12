package main

import (
	"fmt"
	"time"
)

// Multiplexing with select
func multiplexing() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
			// default:
			// 	fmt.Println("No messages received")
		}
	}

	fmt.Println("End of program")

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Received from channel:", msg)
	}
	fmt.Println("Channel closed!")
	// clean up resources if needed
	// No additional cleanup required as the channel is already closed
}
