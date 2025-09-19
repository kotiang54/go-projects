package main

import (
	"sync"
	"time"
)

// RateLimiter implements a token bucket algorithm for rate limiting.
// It allows a certain number of requests in a given time frame.
// It uses a buffered channel to represent available tokens, and refills tokens at a specified interval.
// The struct also includes synchronization primitives to manage the lifecycle of internal goroutines.
type RateLimiterTBA struct {
	tokens     chan struct{}
	refillTime time.Duration

	// Optional: to manage goroutine lifecycle
	wg          sync.WaitGroup
	stopChannel chan struct{}
}

// NewRateLimiter creates a new RateLimiter with the specified rate limit and refill time.
//   - rateLimit: maximum number of requests allowed in the given time frame
//   - refillTime: duration to wait before refilling the token bucket
//
// Example: NewRateLimiter(5, time.Second) allows 5 requests per second.
func NewRateLimiterTBA(rateLimit int, refillTime time.Duration) *RateLimiterTBA {
	rl := &RateLimiterTBA{
		tokens:      make(chan struct{}, rateLimit),
		refillTime:  refillTime,
		stopChannel: make(chan struct{}),
	}

	// Fill the token bucket initially
	for i := 0; i < rateLimit; i++ {
		rl.tokens <- struct{}{}
	}

	// Start the refill goroutine
	rl.wg.Add(1)

	// Start the token refill goroutine
	go rl.startRefill(rateLimit)

	return rl
}

// startRefill refills the token bucket at the specified interval.
func (rl *RateLimiterTBA) startRefill(rateLimit int) {
	defer rl.wg.Done()

	// Calculate the interval for refilling tokens
	interval := rl.refillTime / time.Duration(rateLimit)
	if interval <= 0 {
		interval = time.Nanosecond // Prevent division by zero or negative intervals
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Refill tokens at each tick
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default: // Token bucket is full, do nothing
			}
		case <-rl.stopChannel:
			return
		}
	}
}

// allow checks if a request can be processed based on the available tokens.
func (rl *RateLimiterTBA) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

// Stop stops the rate limiter and cleans up resources.
func (rl *RateLimiterTBA) Stop() {
	close(rl.stopChannel)
	rl.wg.Wait()
}

// Example usage of the RateLimiter
func main() {
	rateLimiter := NewRateLimiterTBA(5, time.Second)
	defer rateLimiter.Stop()

	// Simulate 20 requests
	for i := 0; i < 20; i++ {
		if rateLimiter.Allow() {
			println("Request allowed")
		} else {
			println("Request denied")
		}
		time.Sleep(50 * time.Millisecond) // Simulate time between requests
	}
}
