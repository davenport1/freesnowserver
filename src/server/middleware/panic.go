package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// PanicRecovery - Send back 500 to client in case of server side Panic and log the call stack
func PanicRecovery(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				logger.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
