package config

import (
	"fmt"
	"os"
)

// Config holds the configuration for the application
type Config struct {
	ConnectionString string
}

// LoadConfig loads the configuration from environment variables or defaults
func LoadConfig() (*Config, error) {
	connectionString := os.Getenv("AZURE_SERVICEBUS_CONNECTION_STRING")
	if connectionString == "" {
		return nil, fmt.Errorf("AZURE_SERVICEBUS_CONNECTION_STRING environment variable must be set")
	}

	return &Config{
		ConnectionString: connectionString,
	}, nil
}