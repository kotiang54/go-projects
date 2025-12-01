package middlewares

import (
	"net/http"
	"strings"
)

// MiddlewaresExcludePath wraps given middlewares to exclude specified paths
func MiddlewaresExcludePath(middlewares func(http.Handler) http.Handler, excludePaths ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request path is in the exclude list}
			for _, path := range excludePaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}
			middlewares(next).ServeHTTP(w, r)
		})
	}
}
