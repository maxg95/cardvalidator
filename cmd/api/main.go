package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

// Define a config struct to hold all the configuration settings for application.
type config struct {
	port int
}

// Define an application struct to hold the dependencies.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	// Declare an instance of the config struct.
	var cfg config

	// Read the value of the command-line flags into the config struct.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	// Initialize a new structured logger which writes log entries.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Declare an instance of the application struct.
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a HTTP server.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server.
	logger.Info("starting server", "addr", srv.Addr)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
