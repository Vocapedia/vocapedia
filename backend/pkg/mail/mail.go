package mail

import (
	"net/smtp"

	"github.com/akifkadioglu/vocapedia/pkg/config"
)

var auth smtp.Auth

func InitMail() {
	auth = smtp.PlainAuth("", config.ReadValue().SMTP.From, config.ReadValue().SMTP.Password, config.ReadValue().SMTP.Host)
}

func Send(to string, subject []byte, template string) error {
	err := smtp.SendMail("smtp.example.com:587", auth, "seninmailin@example.com", []string{to}, subject)
	return err
}
