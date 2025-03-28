package auth

type _login struct {
	Email string `json:"email"`
}
type _otp struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type _emailData struct {
	Header      string
	Description string
	Warning     string
	Action      string
	Code        string
	Host        string
}
