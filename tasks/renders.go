package tasks

import (
	"embed"
	"html/template"
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

