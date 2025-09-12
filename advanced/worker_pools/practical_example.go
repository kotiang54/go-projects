package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

// simulate ticket processing requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost $%d\n", req.numTickets, req.personID, req.cost)
		// Simulate processing time
		time.Sleep(time.Second)
		results <- req.personID // Send back the personID as confirmation
	}
}

func main() {
	numRequest := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequest)
	ticketResults := make(chan int)

	// Start ticket processors/workers
	for i := 0; i < numRequest; i++ {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// Simulate sending ticket requests - JobQueue
	for i := 1; i <= numRequest; i++ {
		req := ticketRequest{
			personID:   i,
			numTickets: i * 2, // Each person requests double their ID in tickets
			cost:       i * 2 * price,
		}
		ticketRequests <- req
	}
	close(ticketRequests) // Close the requests channel

	// Collect results
	for i := 1; i <= numRequest; i++ {
		result := <-ticketResults
		fmt.Printf("Ticket processing confirmed for personID %d\n", result)
	}
}
