// Package config is responsible for initializing and providing access
// to global application configuration such as logging.
package config

import (
	"log/slog"
	"os"
)

// logger is a package-level variable that holds the reference to the singleton logger instance.
var logger *slog.Logger

// This is where the logger is set up and will be executed before the main function.
func init() {
	logLevel := new(slog.LevelVar)
	logLevel.Set(slog.LevelDebug)
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
}

// GetLogger returns a pointer to the singleton logger instance.
func GetLogger() *slog.Logger {
	return logger
}
