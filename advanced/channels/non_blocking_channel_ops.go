package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	// ***** Non-blocking receive *****
	select {
	case msg := <-ch:
		fmt.Println("Received message:", msg)
	default:
		fmt.Println("No message received")
	}

	// ***** Non-blocking send *****
	select {
	case ch <- 1:
		fmt.Print("Sent message: 1\n")
	default:
		fmt.Println("Channel is not ready to receive")
	}

	// ***** Non-blocking operation in real time systems *****

	dataCh := make(chan int)
	quitCh := make(chan bool)

	go func() {
		for {
			select {
			case d := <-dataCh:
				fmt.Println("Processing data:", d)
			case <-quitCh:
				fmt.Println("Quitting data processing")
				return
			default:
				fmt.Println("Waiting for data ...")
				time.Sleep(500 * time.Millisecond) // Simulate doing other work
			}
		}
	}()

	for i := range 5 {
		dataCh <- i
		time.Sleep(1 * time.Second)
	}

	quitCh <- true
}
