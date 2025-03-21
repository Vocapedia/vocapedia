package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	base
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
