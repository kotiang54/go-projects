package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
)

// Rate limiting middleware
type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int // Map to track request counts per user/IP
	limit     int            // Maximum allowed requests
	resetTime time.Duration  // Duration to reset the counts
}

// NewRateLimiter creates a new rateLimiter instance
func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	// Start a goroutine to reset visitor counts periodically
	go rl.resetVisitorCount()
	return rl
}

// resetVisitorCount resets the visitor count for all users
func (rl *rateLimiter) resetVisitorCount() {
	ticker := time.NewTicker(rl.resetTime)
	defer ticker.Stop()
	for range ticker.C {
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func clientIP(r *http.Request) string {
	// minimal: strip port from RemoteAddr
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return host
	}
	return r.RemoteAddr
}

// Middleware function to enforce rate limiting based on user/IP addresses
func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	fmt.Println("Rate Limiter Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Rate Limiter Middleware being returned...")
		// Identify the user/IP (for simplicity, using RemoteAddr here)
		user := clientIP(r)

		rl.mu.Lock()
		rl.visitors[user]++
		fmt.Println("Visitor:", user, "Count:", rl.visitors[user])
		exceeded := rl.visitors[user] > rl.limit
		rl.mu.Unlock()

		if exceeded {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("Rate Limiter Middleware ends...")
	})
}
