package main

import (
	"fmt"
	"time"
)

// periodic task execution
// schedule logging, periodic tasks, polling for updates
func periodicTask() {
	fmt.Println("Performing periodic task at:", time.Now())
}

// Tickers are used to perform an action repeatedly at regular intervals.
// They are similar to timers, but instead of firing once, they fire repeatedly.
// A ticker holds a channel that delivers 'ticks' of a clock at intervals.
func main() {
	ticker := time.NewTicker(2 * time.Second)

	i := 1
	for j := 0; j < 5; j++ {
		i *= 2
		fmt.Println(i)
	}

	// Using ticker to perform a task every second
	ticker.Stop() // Stop the previous ticker to prevent resource leaks
	ticker = time.NewTicker(time.Second)
	defer ticker.Stop()

	// Using a ticker to perform a task periodically
	stop := time.After(5 * time.Second) // Define a stop channel to terminate the loop after 5 seconds

	for {
		select {
		case tick := <-ticker.C:
			// This case executes when the ticker sends a tick
			fmt.Println("Tick at", tick)
			periodicTask() // Perform the periodic task
		case <-stop:
			// This case executes when the stop channel sends a signal
			fmt.Println("Ticker stopped.")
			return
		}
	}
}
