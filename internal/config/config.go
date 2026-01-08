package config

import (
	"log"
	"os"
)

func getRequiredEnv(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("missing required environment variable: '%s'\n", name)
	}
	return value
}

func GetDatabaseURL() string {
	return getRequiredEnv("DATABASE_URL")
}

func GetSessionJWTSecret() string {
	return getRequiredEnv("JWT_SESSION_SECRET")
}
