package main

import (
	"log"

	"github.com/johnal95/workouts-pwa/internal/config"
	"github.com/johnal95/workouts-pwa/internal/sqlx"
	"github.com/johnal95/workouts-pwa/migrations"

	_ "github.com/joho/godotenv/autoload"
)

type ProgramArgs struct {
	Port        int
	DatabaseURL string
}

func main() {
	pgDB, err := sqlx.Open(config.GetDatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer pgDB.Close()

	err = sqlx.Migrate(pgDB, migrations.FS, ".")
	if err != nil {
		log.Fatal(err)
	}
}
