package api

import (
	"context"
	"database/sql"
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

type application struct {
	config *config.Config
	logger *log.Logger
	models data.Models
}

func Run(cfg *config.Config, logger *log.Logger) {
	// create a done channel for graceful shutdown of the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// connect to database
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		logger.Fatal(err)
	}

	models := data.NewModels(db)

	// Initialize app with config and logger
	app := &application{
		config: cfg,
		logger: logger,
		models: models,
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Fatal(err)
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
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
