package chapters

import (
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	"github.com/go-chi/render"
)

func First(w http.ResponseWriter, r *http.Request) {
	db := database.Manager()
	var chapters entities.Chapter

	db.Preload("User").Preload("WordBase.Word").First(&chapters)
	render.JSON(w, r, chapters)
}
