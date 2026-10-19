package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

type Middleware struct{}

// Cache response from next handler for x hours
func (h *Middleware) Cache(next http.Handler, hours int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			age := strconv.Itoa(hours * 3600)
			w.Header().Set("Cache-Control", "public, max-age="+age)
		}
		next.ServeHTTP(w, r)
	})
}

// Log all http requests to stdout
func (h *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
