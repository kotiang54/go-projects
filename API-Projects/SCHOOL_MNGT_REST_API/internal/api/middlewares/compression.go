package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

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
	fmt.Println("Compression Middleware...")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Compression Middleware being returned...")
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
