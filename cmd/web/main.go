package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger    *slog.Logger
	staticDir *string
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	staticDir := flag.String("staticDir", "./ui/static/", "Path to static assets")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger:    logger,
		staticDir: staticDir,
	}

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
