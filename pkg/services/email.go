package services

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	smtpServer := os.Getenv("EMAIL_SMTP_SERVER")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")

	from := "noreply@example.com" // Use a generic 'from' address
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Create an SMTP client without authentication
	err := smtp.SendMail(smtpServer+":"+smtpPort,
		smtp.PlainAuth("", "", "", smtpServer),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Print("Email sent successfully")
	return nil
}
