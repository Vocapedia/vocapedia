package stream

import (
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/token"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func StartStream(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room")
	if room == "" {
		http.Error(w, "room is required", http.StatusBadRequest)
		return
	}

	jitsiPass := token.GenerateDeterministicToken(room, 8)

	render.JSON(w, r, map[string]any{
		"password": jitsiPass,
	})
}
