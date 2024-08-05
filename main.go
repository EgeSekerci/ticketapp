package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error starting the server")
	}
}
