package entities

type JwtModel struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Device   Attrs  `json:"device"`
}
