package tasks

import (
	"net/http"
	"strings"

	"github.com/alexedwards/argon2id"

	"ticketapp/db"
	"ticketapp/shared"
)

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

	insert := `INSERT INTO users (email, password, role, name) VALUES ($1, $2, $3, $4)`

	db := db.Connect()
	defer db.Close()

	_, err := db.Exec(insert, user.Email, user.Password, user.Role, user.Name)
	shared.Check(err, "Error inserting user")

	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := User{}

	email := strings.ToLower(r.FormValue("email"))
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
	}

	db := db.Connect()
	defer db.Close()

	emailRow := db.QueryRow("SELECT email, password FROM users WHERE email = $1", email)

	err := emailRow.Scan(
		&user.Email,
		&user.Password,
	)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	match, err := argon2id.ComparePasswordAndHash(password, user.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if match {
		userRow := db.QueryRow("SELECT id, role, name FROM users WHERE email = $1", email)

		err := userRow.Scan(
			&user.Id,
			&user.Role,
			&user.Name,
		)

		tokenString, err := createJWT(&user)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}

		cookie := http.Cookie{
			Name:     "Authorization",
			Value:    tokenString,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Path:     "/",
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   -1,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}
