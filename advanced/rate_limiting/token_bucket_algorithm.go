package main

import "time"

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

// NewRateLimiter creates a new RateLimiter with the specified rate limit and refill time.
//   - rateLimit: maximum number of requests allowed in the given time frame
//   - refillTime: duration to wait before refilling the token bucket
//
// Example: NewRateLimiter(5, time.Second) allows 5 requests per second.
func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}

	// Fill the token bucket initially
	for range rateLimit {
		rl.tokens <- struct{}{}
	}

	// Start the token refill goroutine
	go rl.startRefill()

	return rl
}

// startRefill refills the token bucket at the specified interval.
func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()

	// Refill tokens at each tick
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default: // Channel is full, do nothing
			}
		}
	}
}

// allow checks if a request can be processed based on the available tokens.
func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	// Simulate 10 requests
	for range 10 {
		if rateLimiter.allow() {
			println("Request allowed")
		} else {
			println("Request denied")
		}
		time.Sleep(200 * time.Millisecond) // Simulate time between requests
	}
}
