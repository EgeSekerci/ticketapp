package main

import (
	"embed"
	"fmt"
	"net/http"

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
	mux := http.NewServeMux()

	routes.PageRoutes(mux)
	routes.ServeRoutes(mux, content)
	routes.TicketRoutes(mux)

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	shared.Check(err, "Error starting the server")
}
