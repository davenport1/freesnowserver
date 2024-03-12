// Package config handles configuration of the web server
package config

import (
	"flag"
	"os"
)

// Config struct containing the port, environment being run, and connection string for postgresql
type Config struct {
	Port int
	Env  string
	DSN  string
}

// LoadConfig loads configuration from environment variables/input
func LoadConfig() *Config {
	var cfg Config
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.DSN, "db-dsn", os.Getenv("FREESNOW_DB_DSN"), "POSTGRESQL DSN")
	flag.Parse()
	return &cfg
}
