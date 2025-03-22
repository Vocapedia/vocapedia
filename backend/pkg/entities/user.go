package entities

type User struct {
	Base
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"-" gorm:"not null;unique"`
}
