package middleware

import (
	"freesnow/api"
	"net/http"
	"runtime/debug"
)

// PanicRecovery - Send back 500 to client in case of server side Panic and log the call stack
func PanicRecovery(app *api.Application, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(app *api.Application) {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				app.Logger.Println(string(debug.Stack()))
			}
		}(app)
		next.ServeHTTP(w, r)
	})
}
