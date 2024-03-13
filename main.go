package main

import (
	"freesnow/api"
	"freesnow/config"
	"freesnow/data"
	"freesnow/log"
)

func main() {
	// configure logger
	logger := log.NewLogger()

	// get config for server
	cfg := config.LoadConfig()

	// connect to the database
	models := data.Initialize(cfg, logger)

	// run the app
	api.Run(cfg, logger, models)

	defer models.CloseConnection(logger)
}
