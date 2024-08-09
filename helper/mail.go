package helper

import (
	"github.com/NY-Daystar/corpos-christie/settings"
	"gopkg.in/gomail.v2"
)

// Format and configure mail to send
func NewMail(from string, to string, subject string, body string) *gomail.Message {
	// TODO faire une methode format pour le body du mail + ajouter une signature
	var mail = gomail.NewMessage()
	mail.SetHeader("From", from) // Tester l'envoi sans ce header
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)

	return mail
}

// Configure and return SMTP client to send mail
func NewSMTP(config *settings.Smtp) *gomail.Dialer {
	return gomail.NewDialer(config.Host, config.Port, config.User, config.Password)
}
