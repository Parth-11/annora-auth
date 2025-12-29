package mailer

import (
	"bytes"
	"html/template"
)

func renderTemplate(path string, data any) (string, error) {
	tmpl, err := template.ParseFiles("internal/mailer/templates/" + path)
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
