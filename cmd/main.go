package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/johnal95/workouts-pwa/cmd/app"
	"github.com/johnal95/workouts-pwa/cmd/routes"
)

type ProgramArgs struct {
	Port        int
	DatabaseURL string
}

func getProgramArgs() *ProgramArgs {
	var port int
	flag.IntVar(&port, "port", 8080, "go server port")
	flag.Parse()

	tempDatabaseURL := "postgres://postgres:postgres@127.0.0.1:5432/devdb"

	return &ProgramArgs{
		Port:        port,
		DatabaseURL: tempDatabaseURL,
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
