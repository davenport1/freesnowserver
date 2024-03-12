package api

import (
	"net/http"
)

func (app *application) router() *http.ServeMux {
	// Instantiate our ServeMux and add health check
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", nil)
	mux.HandleFunc("/v1/books", nil)
	mux.HandleFunc("/v1/books/", nil)
	return mux
}
