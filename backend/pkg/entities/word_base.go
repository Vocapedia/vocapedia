package entities

import (
	"github.com/google/uuid"
)

type WordBase struct {
	Base
	ChapterID uuid.UUID `json:"chapter_id" gorm:"not null"`
	Chapter   *Chapter  `json:"chapter,omitempty" gorm:"foreignKey:ChapterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type      string    `json:"type" gorm:"not null"`
	Word      []Word    `json:"words,omitempty"`
}
