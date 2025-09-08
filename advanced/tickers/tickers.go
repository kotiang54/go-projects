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
	defer ticker.Stop()

	i := 1
	for range 5 {
		i *= 2
		fmt.Println(i)
	}

	ticker = time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at", tick)
		case <-stop:
			fmt.Println("Ticker stopped.")
			return
		}
	}

	// Using ticker to perform a task every second
	ticker = time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			periodicTask()
		}
	}
}
