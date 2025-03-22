package entities

type User struct {
	Base
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"-" gorm:"not null;unique"`
	Favorites []Chapter `json:"favorites,omitempty" gorm:"many2many:user_favorites;"`
}
