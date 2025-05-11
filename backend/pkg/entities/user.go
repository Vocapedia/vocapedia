package entities

import (
	"fmt"
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/search"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

type User struct {
	Base
	PhotoURL     string    `json:"photo_url" gorm:"type:text;index"`
	Name         string    `json:"name"`
	Vocatoken    string    `json:"-" gorm:"not null;unique"`
	VocatokenVal int       `json:"-" gorm:"default:0"`
	Biography    string    `json:"biography"`
	Username     string    `json:"username" gorm:"type:citext;not null;unique"`
	Email        string    `json:"-" gorm:"not null;unique"`
	Approved     bool      `json:"approved" gorm:""`
	Chapter      []Chapter `json:"chapters,omitempty" gorm:"foreignkey:creator_id"`
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	client := search.Meili()
	index := client.Index("users")

	document := map[string]any{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username,
		"photo_url": user.PhotoURL,
	}

	_, err = index.AddDocuments([]map[string]any{document})
	if err != nil {
		log.Fatal("Error adding document to Meilisearch: ", err)
	}
	return nil
}

func (user *User) BeforeDelete(tx *gorm.DB) (err error) {
	client := search.Meili()
	index := client.Index("users")
	searchRes, err := index.Search("", &meilisearch.SearchRequest{
		Filter: fmt.Sprintf("id = %d", user.ID),
		Limit:  1,
	})
	if err != nil {
		log.Printf("❌ Search failed: %v", err)
		return err
	}

	var ids []string
	for _, hit := range searchRes.Hits {
		if id, ok := hit.(map[string]any)["id"]; ok {
			ids = append(ids, fmt.Sprintf("%v", id))
		}
	}

	if len(ids) > 0 {
		_, err = index.DeleteDocuments(ids)
		if err != nil {
			log.Printf("❌ Failed to delete entries from MeiliSearch: %v", err)
			return err
		}
		log.Printf("✅ Deleted %d entries of user %d from MeiliSearch", len(ids), user.ID)
	}
	return
}

func (user *User) AfterUpdate(tx *gorm.DB) (err error) {
	client := search.Meili()
	index := client.Index("users")

	document := map[string]any{
		"id":        user.ID,
		"name":      user.Name,
		"username":  user.Username,
		"photo_url": user.PhotoURL,
	}

	_, err = index.UpdateDocuments([]map[string]any{document})
	if err != nil {
		log.Printf("❌ Error updating chapter in Meilisearch: %v", err)
		return err
	}

	return nil
}
