package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world!")
	})

	port := ":8080"
	fmt.Printf("Listening on %s \n", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Println("Error starting the server")
	}
}
