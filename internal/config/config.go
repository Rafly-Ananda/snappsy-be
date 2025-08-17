package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MinIOEndpoint        string
	MinIOAccessKey       string
	MinIOSecretKey       string
	MinIOBucket          string
	MinioPresignedExpiry int

	AppPort string
}

func Load() *Config {
	cfg := &Config{
		MinIOEndpoint:        getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey:       getEnv("MINIO_ACCESS_KEY", "minioadmin"),
		MinIOSecretKey:       getEnv("MINIO_SECRET_KEY", "minioadmin"),
		MinIOBucket:          getEnv("MINIO_BUCKET", "images"),
		MinioPresignedExpiry: getEnvInt("MINIO_EXPIRY_IN_MINUTES", 30),
		AppPort:              getEnv("APP_PORT", "8080"),
	}
	return cfg
}

func getEnv(key, fallback string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if val := os.Getenv(key); val != "" {
		return val
	}

	log.Printf("using fallback for %s", key)
	return fallback
}

func getEnvInt(key string, fallback int) int {
	// load .env once (you might want to move this out of here to avoid reloading every call)
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
		log.Printf("invalid int for %s: %s, using fallback", key, val)
	}
	return fallback
}
