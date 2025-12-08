package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

func SendMailSimple(subject string, body string, to []string) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailFrom := os.Getenv("EMAIL_FROM")

	auth := smtp.PlainAuth(
		"",
		emailUsername,
		emailPassword,
		emailHost,
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		emailHost+":"+emailPort,
		auth,
		emailFrom,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func SendEmail(subject string, templatePath string, to []string, data interface{}) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailPort := os.Getenv("EMAIL_PORT")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailFrom := os.Getenv("EMAIL_FROM")

	// Parse HTML Template
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	// t.Execute(&body, struct {
	// 	FirstName string `json:"firstName"`
	// 	Email     string `json:"email"`
	// }{FirstName: "Bhavishaya Khandelwal", Email: "bhavishaya.khandelwal@dianapps.com"})

	// Pass dynamic struct here
	if err := t.Execute(&body, data); err != nil {
		fmt.Println(err)
		return
	}

	auth := smtp.PlainAuth(
		"",
		emailUsername,
		emailPassword,
		emailHost,
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		emailHost+":"+emailPort,
		auth,
		emailFrom,
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
