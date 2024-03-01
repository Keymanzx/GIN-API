package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort  string
	MongoURI string
	// Add other configuration variables as needed
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Read environment variables
	AppPort := os.Getenv("APP_PORT")
	MongoURI := os.Getenv("MONGO_URI")

	// Create a Config object with the read values
	config := &Config{
		AppPort:  AppPort,
		MongoURI: MongoURI,
		// Add other configuration variables as needed
	}

	return config, nil
}
