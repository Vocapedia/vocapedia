package server

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"

	"github.com/akifkadioglu/vocapedia/pkg/controllers/auth"
	"github.com/akifkadioglu/vocapedia/pkg/controllers/search"
	"github.com/akifkadioglu/vocapedia/pkg/embed"
	customMiddleware "github.com/akifkadioglu/vocapedia/pkg/middleware"
	"github.com/akifkadioglu/vocapedia/pkg/token"
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
	r.Use(customMiddleware.Language)

	r.Route("/api/v1", func(api chi.Router) {
		api.Group(func(api chi.Router) {
			api.Use(jwtauth.Verifier(token.TokenAuth()))
			api.Use(jwtauth.Authenticator(token.TokenAuth()))

			api.Get("/search", search.Search)
		})
		api.Get("/get-token", auth.GetToken)

		api.Group(func(api chi.Router) {
			api.Route("/auth", func(api chi.Router) {
				api.Post("/login", auth.Login)
			})
		})
	})

	r.HandleFunc("/auth", auth.Token)

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
