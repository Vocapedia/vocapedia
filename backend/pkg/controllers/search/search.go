package search

import (
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/go-chi/render"
)

func Search(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	chapters := []entities.Chapter{}
	query := r.URL.Query().Get("q")

	err := db.Raw(`
        SELECT * FROM chapters
        WHERE title % ? OR content % ?
        ORDER BY similarity(title, ?) DESC
        LIMIT 10
    `, query, query, query).Scan(&chapters).Error
	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	render.JSON(w, r, chapters)

}
