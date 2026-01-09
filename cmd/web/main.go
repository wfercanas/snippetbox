package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	staticDir := flag.String("staticDir", "./ui/static/", "Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(*staticDir))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
