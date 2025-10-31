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
