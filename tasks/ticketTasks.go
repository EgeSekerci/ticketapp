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

func GetTickets(w http.ResponseWriter, r *http.Request) {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tickets")
	shared.Check(err, "Error receiving tickets")

	defer rows.Close()

	var tickets []Ticket

	for rows.Next() {
		var ticket Ticket

		err := rows.Scan(
			&ticket.Title,
			&ticket.Desc,
			&ticket.CreatedAt,
		)
		shared.Check(err, "Error scanning ticket data")
		tickets = append(tickets, ticket)
	}
	shared.Check(rows.Err(), "Error on rows.Next()")
	fmt.Println(tickets)
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

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	fmt.Println("Ticket added successfully\n", ticket)
}
