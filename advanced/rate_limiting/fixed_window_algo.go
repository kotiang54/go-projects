package main

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main() {
	ratelimiter := NewRateLimiter(5, time.Second)

	for i := 1; i <= 10; i++ {
		if ratelimiter.Allow() {
			println("Request", i, "allowed")
		} else {
			println("Request", i, "denied")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
