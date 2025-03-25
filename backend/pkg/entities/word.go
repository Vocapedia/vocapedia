package entities

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Word struct {
	Base
	Lang        string         `json:"lang" gorm:"not null"`
	Word        string         `json:"word" gorm:"not null;type:text"`
	Description string         `json:"description" gorm:"not null;type:text"`
	Examples    pq.StringArray `json:"examples" gorm:"not null;type:text"`
	WordBaseID  uuid.UUID      `json:"word_base_id" gorm:"not null"`
	ChapterID   uuid.UUID      `json:"chapter_id" gorm:"not null"`
}
