package main

import (
	"fmt"
	"time"
)

func longRunningOperation() {
	for i := range 20 {
		fmt.Println("Iteration:", i)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("Starting the app...")
	// create timers using time package
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting for timer.C")

	// Stopping the timer before it expires
	if ok := timer.Stop(); ok {
		fmt.Println("Timer stopped")
	}

	// Reset timer: for only stopped or expired timers
	timer.Reset(time.Second)
	fmt.Println("Waiting for timer.C again: Timer reset")
	<-timer.C // blocking in nature. Blocks until the timer expires
	fmt.Println("Timer expired")

	fmt.Println("")

	// implement timeout using select statement
	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Timed out")
	case <-done:
		fmt.Println("Operation completed")
	}

	fmt.Println("")

	// Schedule delayed operation
	timer = time.NewTimer(2 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("2 seconds passed. Delayed operation executed")
	}()

	fmt.Println("Waiting ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Exiting app...")
	timer.Stop()
}
