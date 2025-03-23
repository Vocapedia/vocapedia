package mail

import (
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/config"
)

var srv EmailIntf

func InitMail() {
	srv = NewEmailService(587, config.ReadValue().SMTP.Host, config.ReadValue().SMTP.From, config.ReadValue().SMTP.Password)

	log.Println("Mail is ready")
}
func Send(to, subject string, template string) (bool, error) {
	isEmailSent, err := srv.SendEmail(to, subject, template)

	return isEmailSent, err
}
