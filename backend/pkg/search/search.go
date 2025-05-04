package search

import (
	"log"

	"github.com/meilisearch/meilisearch-go"
)

var ms meilisearch.ServiceManager

func Meili() meilisearch.ServiceManager {
	if ms.IsHealthy() {
		log.Println("Meilisearch is healthy")

	} else {
		log.Println("Meilisearch is not healthy")
	}
	return ms

}
func InitMeili(host string, apiKey string, indexName string) {

	ms = meilisearch.New(host, meilisearch.WithAPIKey(apiKey))

	if !ms.IsHealthy() {
		log.Println("Meilisearch is not healthy")
	}

	index := ms.Index(indexName)

	_, err := index.FetchInfo()
	if err != nil {
		_, err := ms.CreateIndex(&meilisearch.IndexConfig{
			Uid: indexName,
			PrimaryKey: apiKey,
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
