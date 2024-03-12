package main

import (
	"freesnow/api"
	"freesnow/config"
	log2 "freesnow/log"
)

func main() {
	// configure logger
	log := log2.NewLogger()

	// get config for server
	cfg := config.LoadConfig()

	// run the app
	api.Run(cfg, log)
}
