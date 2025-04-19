package entities

import (
	"fmt"
	"log"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/search"
	"gorm.io/gorm"
)

type Chapter struct {
	Base
	CreatorID   int64      `json:"creator_id" gorm:"not null"`
	Creator     *User      `json:"creator,omitempty" gorm:"foreignKey:CreatorID;references:ID;"`
	Title       string     `json:"title" gorm:"not null;type:text;index"`
	Description string     `json:"description" gorm:"not null;type:text;index"`
	Lang        string     `json:"lang" gorm:"not null"`
	TargetLang  string     `json:"target_lang" gorm:"not null"`
	WordBase    []WordBase `json:"word_bases,omitempty"`
	Tutorial    string     `json:"tutorial" gorm:"type:text"`
}

type UserFavorite struct {
	UserID    int64     `gorm:"primaryKey"`
	ChapterID int64     `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
}

func (u *UserFavorite) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return nil
}

func addChapterToMeilisearch(chapter Chapter) {
	client := search.Meili()
	index := client.Index("chapters")

	document := map[string]any{
		"id":          chapter.ID,
		"title":       chapter.Title,
		"description": chapter.Description,
	}

	_, err := index.AddDocuments([]map[string]any{document})
	if err != nil {
		log.Fatal("Error adding document to Meilisearch: ", err)
	}

	fmt.Println("Chapter added to Meilisearch:", chapter.Title)
}

func (chapter *Chapter) AfterCreate(tx *gorm.DB) (err error) {
	addChapterToMeilisearch(*chapter)
	return nil
}

func (chapter *Chapter) BeforeDelete(tx *gorm.DB) (err error) {
	client := search.Meili()
	index := client.Index("chapters")
	_, err = index.DeleteDocument(fmt.Sprintf("%v", chapter.ID))
	if err != nil {
		log.Printf("❌ Failed to delete chapter from Meilisearch: %v", err)
	}
	return
}

func (chapter *Chapter) AfterUpdate(tx *gorm.DB) (err error) {
	client := search.Meili()
	index := client.Index("chapters")

	document := map[string]any{
		"id":          chapter.ID,
		"title":       chapter.Title,
		"description": chapter.Description,
	}

	_, err = index.UpdateDocuments([]map[string]any{document})
	if err != nil {
		log.Printf("❌ Error updating chapter in Meilisearch: %v", err)
		return err
	}

	fmt.Println("Chapter updated in Meilisearch:", chapter.Title)
	return nil
}

