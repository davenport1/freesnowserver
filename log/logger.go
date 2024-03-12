// Package log logger for freesnow server
package log

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime)
}
