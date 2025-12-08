package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendSimpleEmail(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_HOST"),
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("EMAIL_FROM"),
		// []string{"kartik.bisht@dianapps.com"},
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}
