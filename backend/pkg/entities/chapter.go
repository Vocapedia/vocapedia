package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title        string    `json:"name" gorm:"not null;type:text;index"`
	Description string    `json:"description" gorm:"not null;type:text;index"`
}

func (c *Chapter) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
