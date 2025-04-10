package main

import (
	"context"
	"embed"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ticketapp/routes"
	"ticketapp/shared"
	"ticketapp/tasks"
)

//go:embed templates/** static/*
var content embed.FS

func init() {
	var err error
	err = tasks.ParseAllFiles(content)

	shared.Check(err, "Error parsing templates")
}

func main() {
	log.Println("Starting...")
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	routes.PageRoutes(mux)
	routes.ServeRoutes(mux, content)
	routes.TicketRoutes(mux)
	routes.AuthRoutes(mux)

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
}
