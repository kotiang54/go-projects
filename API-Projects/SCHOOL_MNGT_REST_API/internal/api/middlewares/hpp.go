package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

// HPPOptions defines options for the HPP middleware
type HPPOptions struct {
	CheckQuery                 bool
	CheckBody                  bool
	CheckBodyOnlyForContenType string
	Whitelist                  []string
}

// HPP - HTTP Parameter Pollution middleware to protect against HPP attacks
func Hpp(options HPPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fmt.Println("HPP Middleware...")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("HPP Middleware being returned...")
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyForContenType) {
				// filter the body params
				filterBodyParams(r, options.Whitelist)
			}

			// Check for query parameters
			if options.CheckQuery && r.URL.Query() != nil {
				// filter the query params
				filterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
			fmt.Println("HPP Middleware ends...")
		})
	}
}

// isCorrectContentType checks if the request's Content-Type matches the specified content type
// e.g., "application/x-www-form-urlencoded"
func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

// filterBodyParams removes duplicate parameters from the request body based on the whitelist
// and keeps only the first occurrence of each parameter.
func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0]) // first value
			// r.Form.Set(k, v[len(v) - 1]) // last value
		}
		if !isWhiteListed(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

// filterQueryParams removes duplicate parameters from the query string based on the whitelist
// and keeps only the first occurrence of each parameter.
func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[0]) // first value
			// query.Set(k, v[len(v) - 1]) // last value
		}
		if !isWhiteListed(k, whitelist) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}

// isWhiteListed checks if a parameter is in the whitelist
func isWhiteListed(param string, whitelist []string) bool {
	for _, v := range whitelist {
		if param == v {
			return true
		}
	}
	return false
}
