package middleware

import (
	"freesnow/api"
	"net/http"
	"time"
)

// Logging - logging middleware that logs each request method, uri, and execution time
func Logging(app *api.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		app.Logger.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
