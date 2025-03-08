package search

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

func Search(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	chapters := []entities.Chapter{}
	query := r.URL.Query().Get("q")

	err := db.Raw(`
	SELECT * FROM chapters
	WHERE title % ? OR description % ?
	ORDER BY GREATEST(similarity(title, ?), similarity(description, ?)) DESC,
			 LEAST(levenshtein(title, ?), levenshtein(description, ?))
	LIMIT 10;
`, query, query, query, query, query, query).Scan(&chapters).Error

	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	render.JSON(w, r, chapters)
}
