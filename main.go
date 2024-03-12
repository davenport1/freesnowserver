package main

import (
	"freesnow/api"
	"freesnow/config"
	"freesnow/data"
	log2 "freesnow/log"
)

func main() {
	// configure logger
	log := log2.NewLogger()

	// get config for server
	cfg := config.LoadConfig()

	// connect to the database
	models := data.Initialize(cfg, log)

	// run the app
	api.Run(cfg, log, models)

	defer models.CloseConnection(log)
}
