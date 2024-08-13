package tasks

import (
	"fmt"
	"net/http"
	"time"

	"ticketapp/db"
	"ticketapp/shared"
)

type Ticket struct {
	Id        int
	Title     string
	Desc      string
	CreatedAt time.Time
}

func AddTicket(w http.ResponseWriter, r *http.Request) {
	ticket := Ticket{}

	ticket.Title = r.FormValue("title")
	ticket.Desc = r.FormValue("desc")

	ticket.CreatedAt = time.Now()

	insert := `INSERT INTO "tickets" ("title", "desc", "created_at") VALUES ($1, $2, $3)`

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(insert, ticket.Title, ticket.Desc, ticket.CreatedAt)
	shared.Check(err, "Error inserting ticket")

	fmt.Println("Ticket added successfully\n", ticket)
}
