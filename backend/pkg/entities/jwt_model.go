package entities

type JwtModel struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}
