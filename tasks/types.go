package tasks

import (
	"time"
)

type Ticket struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	SolvedAt    time.Time
	IsSolved    bool
	CreatedBy   float64
}

type TemplateData struct {
	UserInfo  User
	Tickets   []Ticket
	CreatedAt []string
	SolvedAt  []string
	UserName  []string
	IsAdmin   bool
}

var TempData TemplateData

type User struct {
	Id       int
	Email    string
	Password string
	Role     string
	Name     string
}

type contextKey string

const claimsContextKey contextKey = "claims"
