package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Username string    `json:"username" gorm:"not null;unique"`
	Password string    `json:"-"`
	Email    string    `json:"email" gorm:"not null;unique"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
