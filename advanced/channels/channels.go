package main

import "fmt"

func main() {
	// Correct syntax: channelName := make(chan valueType)
	greeting := make(chan string)
	greetString := "Hello, Channel!"

	go func() {
		greeting <- greetString
	}()

	receiver := <-greeting // the receiver is a channel inside the main goroutine
	fmt.Println(receiver)

}
