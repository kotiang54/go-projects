package middlewares

import (
	"fmt"
	"net/http"
)

func SecurityHeaders(next http.Handler) http.Handler {
	fmt.Println("Security Headers Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Security Headers Middleware being returned...")
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
		fmt.Println("Security Headers Middleware ends...")
	})
}
