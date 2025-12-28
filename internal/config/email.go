package config

type EmailConfig struct {
	From     string
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}