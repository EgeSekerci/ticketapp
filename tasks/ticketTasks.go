package tasks

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ticketapp/db"
	"ticketapp/shared"
)

func GetTickets(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(claimsContextKey).(jwt.MapClaims)
	if claims == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	userId, ok := claims["userId"]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	role, ok := claims["userRole"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	var tempData TemplateData
	if ok {
		if role == "user" {
			query := "SELECT title, description, created_at FROM tickets WHERE is_solved = false AND created_by = $1"
			tempData = GetTicket(query, userId)
		}
		if role == "admin" {
			query := "SELECT title, description, created_at, created_by FROM tickets WHERE is_solved = false"
			tempData = GetAllTickets(query)

			tempData.IsAdmin = true
		}
	}
	err := tmpl.ExecuteTemplate(w, "getTickets", tempData)
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

	description := r.FormValue("description")
	if description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}
	ticket.Description = description

	ticket.CreatedAt = time.Now()

	insert := `INSERT INTO tickets (title, description, created_at, created_by) VALUES ($1, $2, $3, $4)`

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(insert, ticket.Title, ticket.Description, ticket.CreatedAt, ticket.CreatedBy)
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
