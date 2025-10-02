package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Why use signals?
// Signals are a way to notify a process that a specific event has occurred.
// They are used for inter-process communication and can be sent by the operating system or other processes.
// Common use cases include:
// 	- Graceful shutdown: Catching termination signals (like SIGINT or SIGTERM) to perform cleanup tasks before exiting.
// 	- Reloading configuration: Catching SIGHUP to reload configuration files without restarting the process.
// 	- Custom actions: Defining custom signal handlers for specific application needs.

func main() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)

	// Create a channel to receive OS signals
	signals := make(chan os.Signal, 1)

	// Notify channel on interrupt or terminal signals
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		fmt.Println("Received signal:", sig)

		// Perform cleanup tasks here
		fmt.Println("Performing cleanup tasks...")

		// Exit the program
		os.Exit(0)
	}()

	// Simulate a long-running process
	fmt.Println("Process is running.....")
	fmt.Println("Press Ctrl+C to exit.")

	for {
		time.Sleep(time.Second)
	}
}
