package tasks

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/golang-jwt/jwt/v5"

	"ticketapp/shared"
)

func ServeStaticFiles(w http.ResponseWriter, r *http.Request, content embed.FS) {
	fs := http.FileServer(http.FS(content))
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error serving static files: %v", r)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		fs.ServeHTTP(w, r)
	}).ServeHTTP(w, r)
}

func getTemplateFiles(patterns ...string) []string {
	var files []string
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		shared.Check(err, "Failed to get templates")

		files = append(files, matches...)
	}
	return files
}

var tmpl *template.Template

func ParseAllFiles(content embed.FS) error {
	var err error
	tmpl, err = template.ParseFS(
		content,
		"templates/*.html",
		"templates/*/*.html",
	)
	return err
}

func RenderHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	claims := r.Context().Value(claimsContextKey).(jwt.MapClaims)
	if claims == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if page, ok := claims["userRole"].(string); ok {
		switch page {
		case "admin":
			page = "adminhome"
		default:
			page = "userhome"
		}
		err := tmpl.ExecuteTemplate(w, page, nil)
		shared.Check(err, "Error executing template")
	}
}

func RenderSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(w, "signupLayout", nil)
	shared.Check(err, "Error executing template")
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(w, "loginLayout", nil)
	shared.Check(err, "Error executing template")
}
