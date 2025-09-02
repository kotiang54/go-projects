package main

import "fmt"

func chan_directions() {

	ch := make(chan int)

	// producer only channel
	producer(ch)

	// consumer only channel
	consumer(ch)
}

// producer sends data to the channel
func producer(ch chan<- int) {
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf("Sending %v\n", i)
		}
		close(ch)
	}()
}

// consumer receives data from the channel
func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("Received %v\n", i)
	}
}
