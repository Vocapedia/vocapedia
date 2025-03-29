package entities

type Token struct {
	Base
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
	Device Attrs `json:"device"`
}

type Device struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Os       string `json:"os"`
	Browser  string `json:"browser"`
	Platform string `json:"platform"`
	Language string `json:"language"`
}
