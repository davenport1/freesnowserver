package data

import (
	"database/sql"
	"freesnow/config"
	"log"
)

// Initialize - initializes the database connection and returns a Models struct to the calling function
// Parameters - cfg configuration for connection string, logger
// Throws an error if there is an error opening the connection or pinging the database
func Initialize(cfg *config.Config, logger *log.Logger) Models {
	// connect to database
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		logger.Fatal(err)
	}

	// Check connection
	if err := db.Ping(); err != nil {
		logger.Fatal(err)
	}

	return NewModels(db)
}

// CloseConnection - Closes the database connection once the server is stopped and main exits.
func (m Models) CloseConnection(logger *log.Logger) {
	logger.Println("closing database connection...")
	if err := m.DB.Close(); err != nil {
		logger.Fatalf("error closing the database connection: %v", err)
	}
}
