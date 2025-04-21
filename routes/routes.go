package routes

import (
	"embed"
	"net/http"

	"ticketapp/tasks"
)

func PageRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", tasks.WithJWTAuth(http.HandlerFunc(tasks.RenderHome)))
	mux.Handle("GET /signup", http.HandlerFunc(tasks.RenderSignup))
	mux.Handle("GET /login", http.HandlerFunc(tasks.RenderLogin))
}

func ServeRoutes(mux *http.ServeMux, content embed.FS) {
	mux.Handle(
		"GET /static/",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) { tasks.ServeStaticFiles(w, r, content) },
		),
	)
}

func TicketRoutes(mux *http.ServeMux) {
	mux.Handle("GET /api/v1/tickets", tasks.WithJWTAuth(http.HandlerFunc(tasks.GetTickets)))
	mux.Handle("POST /api/v1/addTicket", tasks.WithJWTAuth(http.HandlerFunc(tasks.AddTicket)))
	mux.Handle(
		"PATCH /api/v1/solveTicket/{id}",
		tasks.WithJWTAuth(http.HandlerFunc(tasks.SolveTicket)),
	)
}

func AuthRoutes(mux *http.ServeMux) {
	mux.Handle("POST /api/v1/signup", http.HandlerFunc(tasks.Signup))
	mux.Handle("POST /api/v1/login", http.HandlerFunc(tasks.Login))
	mux.Handle("POST /api/v1/logout", http.HandlerFunc(tasks.Logout))
}
