package server

import (
	"freesnow/server/middleware"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return middleware.Logging(next, logger)
}

func PanicRecoveryMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return middleware.PanicRecovery(next, logger)
}
