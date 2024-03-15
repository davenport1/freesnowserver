package server

import (
	"freesnow/server/api"
	"freesnow/server/middleware"
	"net/http"
)

func (app *Application) router() *http.ServeMux {
	// Instantiate our ServeMux and add health check
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", api.HealthCheck)

	// Configure routes and handlers
	//mux.HandleFunc("/v1/books", nil)
	//mux.HandleFunc("/v1/books/", nil)

	// Setup the middleware pipeline
	handler := LoggingMiddleware(mux, app.Logger)
	handler = middleware.PanicRecovery(handler, app.Logger)

	return mux
}
