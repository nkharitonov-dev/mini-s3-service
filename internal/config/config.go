package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port      string `env:"PORT"`
	DebugMode bool   `env:"DEBUG_MODE"`
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	debugStr := os.Getenv("DEBUG_MODE")
	debug := false
	if debugStr != "" {
		var err error
		debug, err = strconv.ParseBool(debugStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DEBUG_MODE: %w", err)
		}
	}

	return &Config{
		Port:      port,
		DebugMode: debug,
	}, nil
}
