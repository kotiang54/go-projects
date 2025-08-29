package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	go func() {
		fmt.Println("Sending ... ")
		ch <- 9 // Blocking until the value is received
		// time.Sleep(1 * time.Second)
		fmt.Println("Sent value 9")
	}()

	value := <-ch // Blocking until a value is sent
	fmt.Println("Received value:", value)

}
