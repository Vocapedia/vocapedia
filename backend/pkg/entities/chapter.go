package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chapter struct {
	Base
	CreatorID   uuid.UUID  `json:"creator_id" gorm:"not null"`
	Creator     *User      `json:"creator,omitempty" gorm:"foreignKey:CreatorID;references:ID;"`
	Title       string     `json:"title" gorm:"not null;type:text;index"`
	Description string     `json:"description" gorm:"not null;type:text;index"`
	Lang        string     `json:"lang" gorm:"not null"`
	TargetLang  string     `json:"target_lang" gorm:"not null"`
	WordBase    []WordBase `json:"word_bases,omitempty"`
	Tutorial    string     `json:"tutorial" gorm:"type:text"`
}

type UserFavorite struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	ChapterID uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
}

func (u *UserFavorite) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return nil
}
