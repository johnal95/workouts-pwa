package config

import (
	"flag"
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
	env := os.Getenv("APP_ENV")
	switch env {
	case "development":
		return EnvDevelopment
	default:
		return EnvProduction
	}
}

func GetPort() int {
	var port int
	flag.IntVar(&port, "port", 8080, "go server port")
	flag.Parse()
	return port
}

func GetDatabaseURL() string {
	return getRequiredEnv("DATABASE_URL")
}

func GetSessionJWTSecret() string {
	return getRequiredEnv("JWT_SESSION_SECRET")
}
