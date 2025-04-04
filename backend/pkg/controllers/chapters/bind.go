package chapters

import (
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

type _compose struct {
	ChapterID   string `json:"chapter_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	TargetLang  string `json:"target_lang"`
	WordBase    []struct {
		Type string `json:"type"`
		Word []struct {
			Lang        string `json:"lang"`
			Word        string `json:"word"`
			Description string `json:"description"`
			Examples    string `json:"examples"`
		} `json:"words,omitempty"`
	} `json:"word_bases,omitempty"`
	Tutorial string `json:"tutorial"`
}
type ChapterDTO struct {
	entities.Chapter
	FavCount    int64 `json:"fav_count"`
	WordCount   int64 `json:"word_count"`
	IsFavorited bool  `json:"is_favorited"`
}
