package routes

import (
	"net/http"

	"ticketapp/tasks"
)

func TicketRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.HandlerFunc(tasks.GetTickets))
	mux.Handle("POST /api/v1/addTicket", http.HandlerFunc(tasks.AddTicket))
}
