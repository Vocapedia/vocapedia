package mail

import (
	"log"
	"net/http"
)

var srv EmailIntf

func InitMail(host, from, password string, port int) {
	srv = NewEmailService(port, host, from, password)

	log.Println("Mail is ready")
}

func Send(r *http.Request, to, subject string, template string) (bool, error) {
	isEmailSent, err := srv.SendEmail(to, subject, template, r.Header.Get("Accept-Language"))

	return isEmailSent, err
}
