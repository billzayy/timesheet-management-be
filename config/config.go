package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return &Config{}, err
	}

	return &Config{
		DBHost: os.Getenv("POSTGRES_HOST"),
		DBPort: os.Getenv("POSTGRES_PORT"),
		DBUser: os.Getenv("POSTGRES_USERNAME"),
		DBPass: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
	}, nil
}
