package tasks

import (
	"database/sql"

	"ticketapp/db"
	"ticketapp/shared"
)

func GetTicket(query string, userId interface{}) TemplateData {
	db := db.Connect()
	defer db.Close()
	rows, err := db.Query(query, userId)
	shared.Check(err, "Error receiving tickets")
	defer rows.Close()

	var tickets []Ticket
	var createdAtString []string

	for rows.Next() {
		var ticket Ticket

		var createdAt sql.NullTime

		err := rows.Scan(
			&ticket.Title,
			&ticket.Description,
			&createdAt,
		)
		shared.Check(err, "Error scanning ticket data")

		if createdAt.Valid {
			createdAtString = append(
				createdAtString,
				createdAt.Time.Format("02.01.2006 15:04"),
			)
		} else {
			createdAtString = append(createdAtString, "")
		}

		tickets = append(tickets, ticket)
	}
	shared.Check(rows.Err(), "Error on rows.Next()")

	TempData.Tickets = tickets
	TempData.CreatedAt = createdAtString

	return TempData
}

func GetAllTickets(query string) TemplateData {
	db := db.Connect()
	defer db.Close()
	rows, err := db.Query(query)
	shared.Check(err, "Error receiving tickets")
	defer rows.Close()

	var tickets []Ticket
	var createdAtString []string
	var userName []string

	for rows.Next() {
		var ticket Ticket

		var createdAt sql.NullTime

		err := rows.Scan(
			&ticket.Title,
			&ticket.Description,
			&createdAt,
			&ticket.CreatedBy,
		)
		shared.Check(err, "Error scanning ticket data")

		if createdAt.Valid {
			createdAtString = append(
				createdAtString,
				createdAt.Time.Format("02.01.2006 15:04"),
			)
		} else {
			createdAtString = append(createdAtString, "")
		}
		userName = append(userName, getUserName(ticket.CreatedBy))

		tickets = append(tickets, ticket)
	}
	shared.Check(rows.Err(), "Error on rows.Next()")

	TempData.Tickets = tickets
	TempData.CreatedAt = createdAtString
	TempData.UserName = userName

	return TempData
}

func getUserName(userId float64) string {
	user := User{}
	db := db.Connect()
	defer db.Close()

	row := db.QueryRow("SELECT name FROM users WHERE id = $1", userId)

	err := row.Scan(
		&user.Name,
	)
	if err != nil {
		shared.Check(err, "Unauthorized")
	}
	return user.Name
}
