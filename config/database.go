package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config struct for application configuration
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

// LoadConfig loads the configuration from environment variables or .env file
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found")
	}

	// Create a new Config struct
	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
	}

	// Check if all required environment variables are set
	if config.DBHost == "" || config.DBPort == "" || config.DBUser == "" || config.DBPassword == "" || config.DBName == "" || config.AppPort == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}
