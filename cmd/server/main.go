package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/internal/app"
	"github.com/johnal95/workouts-pwa/internal/config"
	"github.com/johnal95/workouts-pwa/internal/logging"
	"github.com/johnal95/workouts-pwa/internal/routes"

	_ "github.com/joho/godotenv/autoload"
)

type ProgramArgs struct {
	Port        int
	DatabaseURL string
}

func main() {
	logger := logging.NewLogger(config.GetAppEnv())
	slog.SetDefault(logger)

	app, err := app.NewApplication(&app.ApplicationOptions{
		DatabaseURL: config.GetDatabaseURL(),
		Logger:      logger,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetPort()),
		Handler: routes.SetupRoutesHandler(app),
	}

	logger.Info(fmt.Sprintf("listening on %s", server.Addr))

	err = server.ListenAndServe()

	log.Fatal(err)
}
