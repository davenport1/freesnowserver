package server

import (
	"freesnow/server/api"
	"net/http"
)

func (app *Application) router() http.Handler {
	// Instantiate our ServeMux and add health check
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", api.HealthCheck)

	// Configure routes and handlers
	mux.HandleFunc("/v1/resorts", func(w http.ResponseWriter, r *http.Request) {
		api.GetAllResorts(w, r, &app.models)
	})
	mux.HandleFunc("v1/resort", func(w http.ResponseWriter, r *http.Request) {
		api.Resort(w, r, &app.models)
	})
	//mux.HandleFunc("/v1/books", nil)
	//mux.HandleFunc("/v1/books/", nil)

	// Setup the middleware pipeline
	handler := LoggingMiddleware(mux, app.Logger)
	handler = PanicRecoveryMiddleware(handler, app.Logger)

	return handler
}
