package routes

import (
	"net/http"

	"ticketapp/tasks"
)

func PageRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.HandlerFunc(tasks.RenderHome))
}

func TicketRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/v1/tickets", http.HandlerFunc(tasks.GetTickets))
	mux.Handle("POST /api/v1/addTicket", http.HandlerFunc(tasks.AddTicket))
}
