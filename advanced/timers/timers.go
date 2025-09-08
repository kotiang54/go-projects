package main

import (
	"fmt"
	"time"
)

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
	timer.Reset(time.Second * 1)
	<-timer.C // blocking in nature. Blocks until the timer expires
	fmt.Println("Timer expired")
}
