package tasks

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ticketapp/db"
	"ticketapp/shared"
)

type Ticket struct {
	Id        int
	Title     string
	Desc      string
	CreatedAt time.Time
	SolvedAt  time.Time
	IsSolved  bool
	CreatedBy float64
}
type TemplateData struct {
	Tickets   []Ticket
	CreatedAt []string
	SolvedAt  []string
}

func GetTickets(w http.ResponseWriter, r *http.Request) {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tickets WHERE is_solved = false")
	shared.Check(err, "Error receiving tickets")

	defer rows.Close()

	var tickets []Ticket
	var createdAtString, solvedAtString []string

	for rows.Next() {
		var ticket Ticket

		var solvedAt sql.NullTime
		var createdAt sql.NullTime

		err := rows.Scan(
			&ticket.Title,
			&ticket.Desc,
			&createdAt,
			&solvedAt,
			&ticket.IsSolved,
			&ticket.Id,
			&ticket.CreatedBy,
		)
		shared.Check(err, "Error scanning ticket data")

		if createdAt.Valid {
			createdAtString = append(createdAtString, createdAt.Time.Format("02.01.2006 15:04"))
		} else {
			createdAtString = append(createdAtString, "")
		}
		if solvedAt.Valid {
			solvedAtString = append(solvedAtString, solvedAt.Time.Format("02.01.2006 15:04"))
		} else {
			solvedAtString = append(solvedAtString, "")
		}

		tickets = append(tickets, ticket)
	}
	shared.Check(rows.Err(), "Error on rows.Next()")

	templateData := TemplateData{
		Tickets:   tickets,
		CreatedAt: createdAtString,
		SolvedAt:  solvedAtString,
	}
	err = tmpl.ExecuteTemplate(w, "getTickets", templateData)
	shared.Check(err, "Error executing template")
}

func AddTicket(w http.ResponseWriter, r *http.Request) {
	ticket := Ticket{}

	claims := r.Context().Value(claimsContextKey).(jwt.MapClaims)
	if claims != nil {
		if userId, ok := claims["userId"]; ok {
			ticket.CreatedBy = userId.(float64)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	ticket.Title = title

	desc := r.FormValue("desc")
	if desc == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}
	ticket.Desc = desc

	ticket.CreatedAt = time.Now()

	insert := `INSERT INTO "tickets" ("title", "desc", "created_at", "created_by") VALUES ($1, $2, $3, $4)`

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(insert, ticket.Title, ticket.Desc, ticket.CreatedAt, ticket.CreatedBy)
	shared.Check(err, "Error inserting ticket")

	RenderHome(w, r)
}

func SolveTicket(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db := db.Connect()
	defer db.Close()

	time := time.Now()

	update := `UPDATE tickets SET is_solved = true, solved_at = $1 WHERE id = $2`

	_, err := db.Exec(update, time, id)

	shared.Check(err, "Error updating ticket")

	err = tmpl.ExecuteTemplate(w, "getTickets", nil)

	shared.Check(err, "Error executing template")
}
