package api

import (
	"freesnow/api/middleware"
	"net/http"
)

func (app *Application) router() *http.ServeMux {
	// Instantiate our ServeMux and add health check
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", nil)

	// Configure routes and handlers
	mux.HandleFunc("/v1/books", nil)
	mux.HandleFunc("/v1/books/", nil)

	// Setup the middleware pipeline
	handler := middleware.Logging(app, mux)
	handler = middleware.PanicRecovery(app, handler)

	return mux
}
