package config

import (
	"log"
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
	appEnv := os.Getenv("APP_ENV")

	switch appEnv {
	case "local":
		_ = godotenv.Load(".env.local")
	case "production":
		// DO NOT load .env in prod
		log.Println("Running in production, using system env / secrets")
	default:
		_ = godotenv.Load(".env")
	}

	return &Config{
		DBHost: os.Getenv("POSTGRES_HOST"),
		DBPort: os.Getenv("POSTGRES_PORT"),
		DBUser: os.Getenv("POSTGRES_USERNAME"),
		DBPass: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),
	}, nil
}
