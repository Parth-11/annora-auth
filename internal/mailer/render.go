package mailer

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func renderTemplate(path string, data any) (string, error) {
	pwd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	tmpl, err := template.ParseFiles(filepath.Join(pwd, "internal", "mailer", "templates", path))
	if err != nil {
		return "", err
	}

	var renderedContent bytes.Buffer
	if err := tmpl.Execute(&renderedContent, data); err != nil {
		return "", err
	}

	content := renderedContent.String()

	return "Content-Type: text/html; charset=\"utf-8\"\r\n" + content, nil
}
