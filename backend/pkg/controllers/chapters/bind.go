package chapters

import (
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)


type _compose struct {
	Title       string `json:"title" gorm:"not null;type:text;index"`
	Description string `json:"description" gorm:"not null;type:text;index"`
	Lang        string `json:"lang" gorm:"not null"`
	TargetLang  string `json:"target_lang" gorm:"not null"`
	WordBase    []struct {
		Type string `json:"type" gorm:"not null"`
		Word []struct {
			Lang        string `json:"lang" gorm:"not null"`
			Word        string `json:"word" gorm:"not null;type:text"`
			Description string `json:"description" gorm:"not null;type:text"`
			Examples    string `json:"examples" gorm:"not null;type:text"`
		} `json:"words,omitempty"`
	} `json:"word_bases,omitempty"`
	Tutorial string `json:"tutorial" gorm:"type:text"`
}
type ChapterDTO struct {
	entities.Chapter
	FavCount    int64 `json:"fav_count"`
	WordCount   int64 `json:"word_count"`
	IsFavorited bool  `json:"is_favorited"`
}
