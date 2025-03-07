package server

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/akifkadioglu/vocapedia/pkg/controllers/search"
	"github.com/akifkadioglu/vocapedia/pkg/embed"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func HttpServer(host string, port int, allowMethods []string, allowOrigins []string, allowHeaders []string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: allowMethods,
		AllowedHeaders: allowHeaders,
		MaxAge:         300,
	}))

	r.Route("/api/v1", func(api chi.Router) {
		api.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
		api.Get("/search", search.Search)
	})

	distFS, _ := fs.Sub(embed.StaticsFS(), "dist")
	fileServer := http.FileServer(http.FS(distFS))
	r.Handle("/assets/*", fileServer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		index, err := embed.StaticsFS().ReadFile("dist/index.html")
		if err != nil {
			http.Error(w, "index.html bulunamadı", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		render.HTML(w, r, string(index))
	})

	log.Default().Printf("HTTP Server is running on http://%s:%d", host, port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)

}
