package routes

import (
	"fmt"
	"net/http"

	"ticketapp/db"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		db := db.Connect()
		if db.Ping() != nil {
			fmt.Println("Error connecting database")
		} else {
			fmt.Println("Connection to the database is successful")
		}
	})
}
