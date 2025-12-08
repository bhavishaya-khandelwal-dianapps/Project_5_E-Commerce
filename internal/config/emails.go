package config

import "os"

type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

// GetEmailConfig reads SMTP/email config from environment variables
func GetEmailConfig() EmailConfig {
	return EmailConfig{
		Host:     os.Getenv("EMAIL_HOST"),
		Port:     os.Getenv("EMAIL_PORT"),
		Username: os.Getenv("EMAIL_USERNAME"),
		Password: os.Getenv("EMAIl_PASSWORD"),
		From:     os.Getenv("EMAIL_FROM"),
	}
}
