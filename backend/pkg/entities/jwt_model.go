package entities

type JwtModel struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Device   Attrs  `json:"device"`
}
