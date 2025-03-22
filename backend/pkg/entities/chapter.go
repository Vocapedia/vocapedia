package entities

import "github.com/google/uuid"

type Chapter struct {
	Base
	UserID      uuid.UUID  `json:"user_id" gorm:"not null"`
	User        *User      `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title       string     `json:"title" gorm:"not null;type:text;index"`
	Description string     `json:"description" gorm:"not null;type:text;index"`
	Lang        string     `json:"lang" gorm:"not null"`
	TargetLang  string     `json:"target_lang" gorm:"not null"`
	WordBase    []WordBase `json:"word_bases,omitempty"`
}
