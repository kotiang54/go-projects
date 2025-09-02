package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed!")
				// clean up resources if needed
				return
			}
			fmt.Println("Received:", msg)
		}
	}
}
