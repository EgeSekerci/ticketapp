package tasks

import (
	"net/http"

	"github.com/alexedwards/argon2id"

	"ticketapp/db"
	"ticketapp/shared"
)

var hashedPassword string

func createHash(passwordToHash string) string {
	hash, err := argon2id.CreateHash(passwordToHash, argon2id.DefaultParams)
	shared.Check(err, "Error while hashing password")

	return hash
}

func Signup(w http.ResponseWriter, r *http.Request) {
	user := User{}

	email := r.FormValue("email")
	password := r.FormValue("password")
	role := r.FormValue("role")
	name := r.FormValue("name")

	if email == "" || password == "" || role == "" || name == "" {
		http.Error(w, "Please fill all the required fields", http.StatusBadRequest)
		return
	}

	user.Email = email
	user.Password = createHash(password)
	user.Role = role
	user.Name = name

	insert := `INSERT INTO "users" ("email", "password", "role", "name") VALUES ($1, $2, $3, $4)`

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(insert, user.Email, user.Password, user.Role, user.Name)
	shared.Check(err, "Error inserting user")
}
