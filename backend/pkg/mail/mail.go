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
	err := smtp.SendMail(config.ReadValue().SMTP.Host+":"+config.ReadValue().SMTP.Port, auth, config.ReadValue().SMTP.From, []string{to}, subject)
	return err
}
