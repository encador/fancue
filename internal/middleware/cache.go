package middleware

import (
	"net/http"
	"strconv"
)

func Cache(next http.Handler, hours int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet || r.Method == http.MethodHead {
			age := strconv.Itoa(hours * 3600)
			w.Header().Set("Cache-Control", "public, max-age="+age)
		}
		next.ServeHTTP(w, r)
	})
}
