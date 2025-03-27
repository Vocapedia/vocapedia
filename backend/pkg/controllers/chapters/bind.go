package chapters

import (
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

type _gameFormat struct {
	Answers []string `json:"answers"`
	Matchs  []_match `json:"matchs"`
}

type _match struct {
	Word   string `json:"word"`
	Answer string `json:"answer"`
}

type _search struct {
	ID          int64          `gorm:"primary_key" json:"id"`
	CreatorID   int64          `json:"creator_id,omitempty" gorm:"foreignKey:CreatorID;references:ID;"`
	Creator     *entities.User `json:"creator,omitempty" gorm:"foreignKey:CreatorID;references:ID;"`
	Title       string         `json:"title" gorm:"not null;type:text;index"`
	Description string         `json:"description" gorm:"not null;type:text;index"`
	Lang        string         `json:"lang" gorm:"not null"`
	TargetLang  string         `json:"target_lang" gorm:"not null"`
	FavCount    int            `json:"fav_count"`
	WordCount   int            `json:"word_count"`
}

type ChapterDTO struct {
	entities.Chapter
	FavCount    int64 `json:"fav_count"`
	WordCount   int64 `json:"word_count"`
	IsFavorited bool  `json:"is_favorited"`
}
