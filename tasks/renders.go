package tasks

import (
	"embed"
	"html/template"
	"net/http"
	"path/filepath"

	"ticketapp/shared"
)

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
	)
	return err
}

func RenderHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(w, "layout", nil)
	shared.Check(err, "Error executing template")
}
