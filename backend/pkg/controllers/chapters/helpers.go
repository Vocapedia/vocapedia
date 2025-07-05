package chapters

import (
	"encoding/json"
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"gorm.io/gorm"
)

func indexOf(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func applyLanguageFilter(query *gorm.DB, userID string) *gorm.DB {
	db := database.Manager()
	var user entities.User
	db.Where("id = ?", userID).First(&user)

	var knownLangs, targetLangs []string

	_ = json.Unmarshal(user.KnownLanguages, &knownLangs)
	_ = json.Unmarshal(user.TargetLanguages, &targetLangs)

	log.Println(knownLangs)
	log.Println(targetLangs)
	
	query = query.Where("lang IN ?", knownLangs)
	query = query.Where("target_lang IN ?", targetLangs)

	return query
}
