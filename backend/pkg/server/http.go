package server

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"

	"github.com/akifkadioglu/vocapedia/pkg/controllers/auth"
	"github.com/akifkadioglu/vocapedia/pkg/controllers/chapters"
	"github.com/akifkadioglu/vocapedia/pkg/controllers/stream"
	"github.com/akifkadioglu/vocapedia/pkg/controllers/user"
	"github.com/akifkadioglu/vocapedia/pkg/embed"
	customMiddleware "github.com/akifkadioglu/vocapedia/pkg/middleware"
	"github.com/akifkadioglu/vocapedia/pkg/token"
)

type Server struct {
	Router *chi.Mux
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountHandlers(host string, port int, allowMethods []string, allowOrigins []string, allowHeaders []string) {
	s.Router.Use(customMiddleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedMethods: allowMethods,
		AllowedHeaders: allowHeaders,
		MaxAge:         300,
	}))
	s.Router.Use(customMiddleware.Language)
	s.Router.Use(customMiddleware.SecurityHeaders)
	//s.Router.Use(customMiddleware.RateLimiter(6, time.Second))
	s.Router.Use(jwtauth.Verifier(token.TokenAuth()))

	s.Router.Route("/v1", func(api chi.Router) {
		api.Group(func(api chi.Router) {
			api.Use(customMiddleware.HandleToken)
			api.Use(jwtauth.Authenticator(token.TokenAuth()))

			api.Route("/chapters", func(api chi.Router) {
				api.Get("/favorite", chapters.Favorites)
				api.Post("/favorite", chapters.Favorite)
				api.Delete("/favorite", chapters.DeleteFavorite)
				api.Post("/compose-by-excel", chapters.ComposeByExcel)
				api.Post("/compose", chapters.Compose)
				api.Put("/compose", chapters.Update)
				api.Delete("/archive/{id}", chapters.Archive)
			})
			api.Route("/user", func(api chi.Router) {
				api.Get("/token", user.Tokens)
				api.Put("/", user.EditUser)
				api.Delete("/token/{id}", user.DeleteToken)
				api.Get("/check", user.Check)
				api.Put("/vocatoken", user.UpdateVocaToken)
				api.Get("/vocatoken", user.GetVocaToken)
				api.Get("/streak", user.DailyStreak)
			})
			api.Route("/auth", func(api chi.Router) {
				api.Delete("/logout", auth.Logout)
			})
			api.Route("/stream", func(api chi.Router) {
				api.Get("/{room}", stream.StartStream)
			})
		})
		api.Route("/public", func(api chi.Router) {
			api.Route("/auth", func(api chi.Router) {
				api.Post("/send-otp", auth.SendOTP)
				api.Post("/verify-otp", auth.VerifyOTP)
			})
			api.Route("/chapters", func(api chi.Router) {
				api.Get("/user", chapters.UserChapters)
				api.Get("/{id}", chapters.GetByID)
				api.Get("/search", chapters.Search)
				api.Get("/trends", chapters.GetTrendingChapters)
				api.Get("/game-format/{id}", chapters.GameFormat)
				api.Get("/hangman/{id}", chapters.GameHangman)
			})
			api.Route("/user", func(api chi.Router) {
				api.Get("/", user.GetByUsername)
				api.Get("/search", user.SearchUsers)
			})
			//api.Post("/speech-to-text", speechtotext.SpeechToText)
		})
		api.Route("/usage", func(api chi.Router) {
			//api.Use(customMiddleware.HandleVocatoken)
			api.Get("/extension/notifier/{id}", chapters.Extension)
		})
	})

	s.Router.HandleFunc("/auth", auth.Token)

	staticsFS, _ := fs.Sub(embed.StaticsFS(), "statics")
	fileServerStatics := http.FileServer(http.FS(staticsFS))

	s.Router.Handle("/og-image.png", http.StripPrefix("/", fileServerStatics))
	s.Router.Handle("/favicon.ico", http.StripPrefix("/", fileServerStatics))

	/* distFS, _ := fs.Sub(embed.DistFS(), "dist")
	fileServer := http.FileServer(http.FS(distFS))
	s.Router.Handle("/assets/*", fileServer)
	s.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		index, err := embed.DistFS().ReadFile("dist/index.html")
		if err != nil {
			http.Error(w, "index.html bulunamadı", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		render.HTML(w, r, string(index))
	}) */

	log.Default().Printf("HTTP Server is running on http://%s:%d", host, port)

}
