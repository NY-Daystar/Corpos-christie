package helper

// TODO mettre un popup d'erreur si le mail s'envoie pas

import (
	"github.com/NY-Daystar/corpos-christie/settings"
	"gopkg.in/gomail.v2"
)

// Format and configure mail to send
func NewMail(from string, to string, subject string, body string) *gomail.Message {
	var mail = gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", body)

	return mail
}

// Configure and return SMTP client to send mail
func NewSMTP(config *settings.Smtp) *gomail.Dialer {
	return gomail.NewDialer(config.Host, config.Port, config.User, config.Password)
}
