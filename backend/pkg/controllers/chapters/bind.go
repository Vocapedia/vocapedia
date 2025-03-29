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


type ChapterDTO struct {
	entities.Chapter
	FavCount    int64 `json:"fav_count"`
	WordCount   int64 `json:"word_count"`
	IsFavorited bool  `json:"is_favorited"`
}
