package main

import (
	"freesnow/config"
	"freesnow/data"
	"freesnow/log"
	"freesnow/server"
)

func main() {
	// configure logger
	logger := log.NewLogger()

	// get config for server
	cfg := config.LoadConfig()

	// connect to the database
	models := data.Initialize(cfg, logger)

	// run the app
	server.Run(cfg, logger, models)

	defer models.CloseConnection(logger)
}
