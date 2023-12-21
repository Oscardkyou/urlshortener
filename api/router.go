package router

import (
	"urlshortener/internal/handler"
	"urlshortener/internal/shortener"

	"github.com/go-chi/chi"
)

func NewRouter(shortenerService *shortener.ShortenerService) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/home", handler.HomePageHandler)
	r.Get("/{shortURL}", handler.RedirectHandler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", handler.SaveURLHandler(shortenerService))
		r.Get("/resolve", handler.ResolveHandler(shortenerService))
	})
	r.Get("/health", handler.HealthCheckHandler)
	r.Get("/ping", handler.PingHandler)
	return r
}
