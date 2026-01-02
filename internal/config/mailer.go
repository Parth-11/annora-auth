package config

import "time"

type MailerConfig struct {
	From     string
	SMTPHost string
	SMTPPort int
	Username string
	Password string
	TokenTTL time.Duration
	ResendLimit int
	ResendLimitTTL time.Duration
}