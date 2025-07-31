package main

import (
	"fmt"
	"time"
)

func main() {
	// Correct syntax: channelName := make(chan valueType)
	greeting := make(chan string)
	greetString := "Hello, Channel!"

	go func() {
		greeting <- greetString
		for _, ch := range "abcde" {
			greeting <- "Alphabet: " + string(ch)
		}
	}()

	receiver := <-greeting // the receiver is a channel inside the main goroutine
	fmt.Println(receiver)

	for range 5 {
		// Receiving from the channel
		msg := <-greeting
		fmt.Println(msg)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished.")
}
