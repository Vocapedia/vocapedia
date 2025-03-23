package auth

type _login struct {
	Email string `json:"email"`
}

type _emailData struct {
	Header      string
	Description string
	Warning     string
	Action      string
	Code        string
	Host        string
}
