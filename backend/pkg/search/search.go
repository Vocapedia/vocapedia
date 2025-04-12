package search

import (
	"fmt"
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
		log.Fatalf("Meilisearch is not healthy")
	}
	return ms

}
func InitMeili() {
	host := config.ReadValue().Meili.Host
	apiKey := config.ReadValue().Meili.APIKey
	indexName := config.ReadValue().Meili.Index

	fmt.Println("Connecting to Meilisearch at:", host)

	ms := meilisearch.New(host, meilisearch.WithAPIKey(apiKey))

	if !ms.IsHealthy() {
		log.Fatalf("❌ Meilisearch is not healthy")
	}

	log.Println("✅ Meilisearch is healthy")

	index := ms.Index(indexName)

	_, err := index.FetchInfo()
	if err != nil {
		_, err := ms.CreateIndex(&meilisearch.IndexConfig{
			Uid: indexName,
		})
		if err != nil {
			log.Fatalf("❌ Failed to create index: %s", err)
		}
		log.Println("📦 Index created:", indexName)
	} else {
		log.Println("📦 Index already exists:", indexName)
	}

	_, err = index.UpdateSearchableAttributes(&[]string{"title", "description"})
	if err != nil {
		log.Fatalf("❌ Failed to set searchable attributes: %s", err)
	}
	log.Println("🔍 Searchable attributes set: title, description")
}
