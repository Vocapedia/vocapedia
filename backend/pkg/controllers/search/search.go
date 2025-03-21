package search

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
)

func Search(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	chapters := []entities.Chapter{}
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 50 {
		limit = l
	}
	offset := (page - 1) * limit

	err := db.Raw(`
		WITH ranked_chapters AS (
			SELECT *,
				(0.7 * similarity(title, ?) + 0.3 * similarity(description, ?)) AS sim_score,
				LEAST(levenshtein(title, ?), levenshtein(description, ?)) AS lev_score
			FROM chapters
			WHERE title % ? OR description % ?
		)
		SELECT * FROM ranked_chapters
		ORDER BY sim_score DESC, lev_score ASC
		LIMIT ? OFFSET ?;
	`, query, query, query, query, query, query, limit, offset).Find(&chapters).Error

	if err != nil {
		render.JSON(w, r, map[string]string{
			"error": err.Error(),
		})
		return
	}
	render.JSON(w, r, map[string]interface{}{
		"page":    page,
		"limit":   limit,
		"results": chapters,
		"hasNext": len(chapters) == limit,
	})
}
