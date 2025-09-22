// Package main demonstrates a thread-safe token-bucket style rate limiter.
// NOTE: Although the type is named LeakyBucket, this implementation behaves
// like a *token bucket*: tokens accumulate over time (up to capacity) and each
// request consumes one token. The sustained rate is ~1 token per `leakRate`.
// A classic "leaky bucket (as queue)" drains at a constant rate and smooths
// bursts by queuing work; this code does not queue—requests are allowed/denied.
package main

import (
	"fmt"
	"sync"
	"time"
)

// LeakyBucket is a simple, thread-safe rate limiter.
//
// Semantics (token bucket):
//   - capacity: maximum tokens the bucket can hold (also the maximum burst).
//   - leakRate: time between token accruals (e.g., 500ms => 2 tokens/sec).
//   - tokens: current number of available tokens (permits).
//   - lastLeak: last time we accounted for token accrual.
//   - mu: protects tokens and lastLeak.
//
// At most one token is added every `leakRate` units of elapsed time,
// up to `capacity`. Each allowed request consumes one token.
type LeakyBucket struct {
	capacity int           // Maximum number of tokens (burst size)
	leakRate time.Duration // Interval between token accruals (1/leakRate ~= tokens per second)

	tokens   int       // Current available tokens
	lastLeak time.Time // Last time we updated (accrued) tokens

	mu sync.Mutex // Guards tokens and lastLeak
}

// NewLeakyBucket constructs a limiter with a starting bucket that is "full"
// (i.e., it begins with `capacity` tokens, allowing an immediate burst).
//
// Example:
//
//	// Allow bursts of 5, refill 1 token every 500ms (~2 RPS sustained).
//	lb := NewLeakyBucket(5, 500*time.Millisecond)
func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,   // start full for burst friendliness
		lastLeak: time.Now(), // base time for accrual calculations (monotonic component included)
	}
}

// Allow returns true if a request is permitted *now* and false otherwise.
// It performs O(1) work and is safe for concurrent callers.
//
// How it works:
//  1. Compute how much time has passed since we last accrued tokens.
//  2. Convert elapsed time into whole tokens (elapsed / leakRate).
//  3. Add those tokens to the bucket, capping at capacity.
//  4. Advance lastLeak by the amount of time equal to the tokens we added.
//     (This preserves any "fractional" remainder for next time.)
//  5. If we have at least one token, consume it and allow; else deny.
//
// NOTE: This method is *non-blocking*: it never waits for tokens to accrue.
// If you need blocking behavior, wrap this with a retry/sleep or add a
// separate Acquire(context) method that waits until a token is available.
func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// 1) How much time has passed since we last updated?
	elapsed := time.Since(lb.lastLeak)
	var tokensToAdd int
	if elapsed > 0 && lb.leakRate > 0 {
		// 2) Convert elapsed time into whole tokens to add
		tokensToAdd = int(elapsed / lb.leakRate)
		if tokensToAdd > 0 {
			// 3) Add tokens, saturating at capacity
			lb.tokens += tokensToAdd
			if lb.tokens > lb.capacity {
				lb.tokens = lb.capacity
			}

			// 4) Advance lastLeak by exactly the time that accounts for the tokens we added.
			//    Any leftover fractional time remains, so we don't "lose" accrual precision.
			lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
		}
	}

	fmt.Printf("Tokens add %d, Tokens substracted %d, Total tokens: %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time: %v\n", lb.lastLeak)

	// 5) Spend a token if available
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

func main() {
	// capacity=5, leakRate=500ms => burst up to 5; sustained ≈ 2 requests/sec.
	leakyBucket := NewLeakyBucket(5, 500*time.Millisecond)
	var wg sync.WaitGroup
	start := make(chan struct{})

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Try to take a token (non-blocking).
			<-start
			allowed := leakyBucket.Allow()
			now := time.Now().Format("15:04:05.0000")
			status := "DENIED"

			if allowed {
				status = "ALLOWED"
			}

			fmt.Printf("[%02d] %s -> %s\n", id, now, status)

			// if leakyBucket.Allow() {
			// 	fmt.Println("Current time: ", time.Now())
			// 	fmt.Println("Request accepted.")
			// } else {
			// 	fmt.Println("Current time: ", time.Now())
			// 	fmt.Println("Request denied!")
			// }
			// time.Sleep(200 * time.Millisecond) // simulate time between requests
		}(i)
	}
	// Wait for all goroutines to finish.
	close(start)
	wg.Wait()
}
