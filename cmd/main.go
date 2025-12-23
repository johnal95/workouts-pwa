package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/app"
	"github.com/johnal95/workouts-pwa/cmd/config"
	"github.com/johnal95/workouts-pwa/cmd/routes"

	_ "github.com/joho/godotenv/autoload"
)

type ProgramArgs struct {
	Port        int
	DatabaseURL string
}

func getProgramArgs() *ProgramArgs {
	var port int
	flag.IntVar(&port, "port", 8080, "go server port")
	flag.Parse()

	return &ProgramArgs{
		Port:        port,
		DatabaseURL: config.GetDatabaseURL(),
	}
}

func main() {
	args := getProgramArgs()

	app, err := app.NewApplication(&app.ApplicationOptions{
		DatabaseURL: args.DatabaseURL,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", args.Port),
		Handler: routes.SetupRoutesHandler(app),
	}

	slog.Info(fmt.Sprintf("listening on %s", server.Addr))

	err = server.ListenAndServe()

	log.Fatal(err)
}
