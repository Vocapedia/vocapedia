package search

import (
	"log"

	"github.com/akifkadioglu/vocapedia/pkg/config"
	"github.com/meilisearch/meilisearch-go"
)

func Meili() meilisearch.ServiceManager {
	ms := meilisearch.New(
		config.ReadValue().Meili.Host,
		meilisearch.WithAPIKey(config.ReadValue().Meili.APIKey),
	)

	if ms.IsHealthy() {
		log.Println("Meilisearch is healthy")

	} else {
		log.Println("Meilisearch is not healthy")
	}
	return ms

}
func InitMeili() {
	host := config.ReadValue().Meili.Host
	apiKey := config.ReadValue().Meili.APIKey
	indexName := config.ReadValue().Meili.Index

	ms := meilisearch.New(host, meilisearch.WithAPIKey(apiKey))

	if !ms.IsHealthy() {
		log.Println("Meilisearch is not healthy")
	}

	index := ms.Index(indexName)

	_, err := index.FetchInfo()
	if err != nil {
		_, err := ms.CreateIndex(&meilisearch.IndexConfig{
			Uid: indexName,
		})
		if err != nil {
			log.Println("Meilisearch is not healthy")
		}
	}

	_, err = index.UpdateSearchableAttributes(&[]string{"title", "description"})
	if err != nil {
		log.Println("❌ Failed to set searchable attributes:", err)
	}
	log.Println("Meili is ready")
}
