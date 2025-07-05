package entities

type JwtModel struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	IsTeacher bool   `json:"is_teacher"`
	Device    Attrs  `json:"device"`
}
