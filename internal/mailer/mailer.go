package mailer

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/config"
)

type Mailer struct {
	Addr string
	Auth smtp.Auth
	From string
	TokenTTL time.Duration
}

func NewMailer(cfg config.MailerConfig) *Mailer {
	auth := smtp.PlainAuth(
		"",
		cfg.Username,
		cfg.Password,
		cfg.SMTPHost,
	)

	return &Mailer{
		Addr: fmt.Sprintf("%s:%d",cfg.SMTPHost,cfg.SMTPPort),
		Auth: auth,
		From: cfg.From,
		TokenTTL: cfg.TokenTTL,
	}
}

func (m *Mailer) SendVerificationEmail(to, token string) error {
	body, err := renderTemplate(
		"verify_email.html",
		map[string]string{
			"url": "https://localhost:8080/auth/verify-email?token=" + token,
		},
	)

	if err != nil {
		return err
	}

	body = "Subject: Verify your email\n" + body

	return smtp.SendMail(
		m.Addr,
		m.Auth,
		m.From,
		[]string{to},
		[]byte(body),
	)
}
