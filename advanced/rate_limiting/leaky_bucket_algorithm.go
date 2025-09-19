package main

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate)
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func main() {
	leakyBucket := NewLeakyBucket(5, 500*time.Millisecond)

	for i := 1; i <= 10; i++ {
		if leakyBucket.Allow() {
			fmt.Println("Request accepted.")
		} else {
			fmt.Println("Request denied!")
		}
		time.Sleep(200 * time.Millisecond)
	}
}
