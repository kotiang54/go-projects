package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Allow multiple origins
var allowedOrigins = []string{
	"https://localhost:3000",
	"https://my-origin-url.com",
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		// Network / transport
		h.Set("X-DNS-Prefetch-Control", "off")
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		// Framing & clickjacking
		h.Set("X-Frame-Options", "DENY")

		// Content Security Policy (tune for your app)
		h.Set("Content-Security-Policy", "default-src 'self'; base-uri 'self'; frame-ancestors 'self'")

		// Referrer policy
		h.Set("Referred-Policy", "no-referrer") // or "strict-origin-when-cross-origin"

		// Legacy (optional; modern Chrome/Edge ignore it)
		// h.Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}

// CORS - Cross-Origin resource sharing
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)

		if !isOriginAllowed(origin) {
			http.Error(w, "Not allowed by CORS", http.StatusForbidden)
			return
		}

		h := w.Header()
		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		h.Set("Access-Control-Allow-Credentials", "true")
		h.Set("Access-Control-Expose-Headers", "Authorization")
		h.Set("Access-Control-Max-Age", "3600")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// responseWriter wraps http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

// Override the WriteHeader method to capture the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

// Track the performance for the API
func ResponseTime(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a custom ResponseWriter to capture the status code
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(wrappedWriter, r)

		// Calculate the duration
		duration := time.Since(start)

		// Log the request details
		fmt.Printf("Method: %s, URL: %s, Status: %d, Duration: %v\n", r.Method, r.URL, wrappedWriter.status, duration.String())
		fmt.Println("Sent Response from Response Writer!")
	})
}

// gzipResponseWriter wraps http.ResponseWriter to provide gzip compression
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

// Override the Write method to use gzip.Writer
func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

// Compression middlewares
// Compress the response using gzip encoding if the client supports it
func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Compression logic goes here
		// Check if client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		// Set the response header to indicate gzip encoding
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		// Wrap the original ResponseWriter with gzipResponseWriter
		gzWriter := &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		next.ServeHTTP(gzWriter, r)
		fmt.Println("Sent Compressed Response from Gzip Writer!")
	})
}

// Rate limiting middleware
type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int // Map to track request counts per user/IP
	limit     int
	resetTime time.Duration
}

// NewRateLimiter creates a new rateLimiter instance
func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		visitors:  make(map[string]int),
		limit:     limit,
		resetTime: resetTime,
	}
	// Start the reset visitor counts periodically
	rl.resetVisitorCount()
	return rl
}

// resetVisitorCount resets the visitor count for all users
func (rl *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}
