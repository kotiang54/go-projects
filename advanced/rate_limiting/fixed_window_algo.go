package main

import (
	"sync"
	"time"
)

// RateLimiter implements a fixed window rate limiting algorithm.
// It allows a maximum number of requests in a defined time window.
type RateLimiter struct {
	mu        sync.Mutex
	count     int           // Current count of requests in the window
	limit     int           // Maximum allowed requests in the window
	window    time.Duration // Duration of the time window
	resetTime time.Time
}

// NewRateLimiter creates a new RateLimiter with the specified limit and window duration.
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

// Allow checks if a request is allowed under the rate limit.
// It returns true if allowed, false otherwise.
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	// Reset the count if the window has passed
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	// Check if the request can be allowed
	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

// Example usage
func main() {
	var wg sync.WaitGroup
	ratelimiter := NewRateLimiter(5, time.Second)

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(n int) {
			if ratelimiter.Allow() {
				println("Request", i, "allowed")
			} else {
				println("Request", i, "denied")
			}
			wg.Done()
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}
