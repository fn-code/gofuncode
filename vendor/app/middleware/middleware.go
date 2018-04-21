package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware is type form declaration http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging for logging your acces the url
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "Desired Content Type; charset=utf-8")
			w.Header().Set("Cache-Control", "no-cache, no-store")
			w.Header().Set("Pragma", "no-cache")
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f.ServeHTTP(w, r)
		}
	}
}

// CSRFError is gorilla csrf handling error
func CSRFError() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}

// Chain for wrap all middleware, we use
func Chain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	for _, mn := range m {
		f = mn(f)
	}
	return f
}
