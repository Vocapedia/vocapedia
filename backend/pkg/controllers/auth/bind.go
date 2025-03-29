package auth

import "github.com/akifkadioglu/vocapedia/pkg/entities"

type _login struct {
	Email string `json:"email"`
}
type _otp struct {
	Email  string          `json:"email"`
	OTP    string          `json:"otp"`
	Device entities.Device `json:"device"`
}

type _emailData struct {
	Header      string
	Description string
	Warning     string
	Action      string
	Code        string
	Host        string
}
