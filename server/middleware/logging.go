package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logging - logging middleware that logs each request method, uri, and execution time
func Logging(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
