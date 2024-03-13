package api

import (
	"context"
	"errors"
	"fmt"
	"freesnow/config"
	"freesnow/data"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "1.0.0"

type Application struct {
	config *config.Config
	Logger *log.Logger
	models data.Models
}

// Run - init and run the server
func Run(cfg *config.Config, logger *log.Logger, models data.Models) {
	// create a done channel for graceful shutdown of the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Initialize app with config and logger
	app := &Application{
		config: cfg,
		Logger: logger,
		models: models,
	}

	logger.Printf("database connection pool established")

	addr := fmt.Sprintf(":%d", cfg.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      app.router(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// begin listening within go routine, handle error
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err)
		}
	}()
	logger.Printf("Starting %s server on %s", cfg.Env, addr)

	<-done // Interrupt begun
	logger.Println("server shutting down...")

	// handle cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling
		cancel()
	}()

	// shutdown server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutdown failed: %+v", err)
	}
	logger.Println("server shutdown successful")
}
