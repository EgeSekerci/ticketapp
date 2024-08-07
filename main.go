package main

import (
	"fmt"
	"log"
	"net/http"

	"ticketapp/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.HandleRoutes(mux)

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error starting the server")
	}
}
