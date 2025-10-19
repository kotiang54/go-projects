package middlewares

import (
	"fmt"
	"net/http"
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
