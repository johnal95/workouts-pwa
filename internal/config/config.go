package config

import (
	"log"
	"os"
)

type Environment string

const (
	EnvDevelopment = Environment("development")
	EnvProduction  = Environment("production")
)

func getRequiredEnv(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("missing required environment variable: '%s'\n", name)
	}
	return value
}

func GetAppEnv() Environment {
	value, exists := os.LookupEnv("APP_ENV")
	if exists && value == "development" {
		return EnvDevelopment
	}
	return EnvProduction
}

func GetDatabaseURL() string {
	return getRequiredEnv("DATABASE_URL")
}

func GetSessionJWTSecret() string {
	return getRequiredEnv("JWT_SESSION_SECRET")
}
