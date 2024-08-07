package main

import (
	"fmt"
	"net/http"

	"ticketapp/routes"
	"ticketapp/shared"
)

func main() {
	mux := http.NewServeMux()

	routes.HandleRoutes(mux)

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	shared.Check(err, "Error starting the server")
}
