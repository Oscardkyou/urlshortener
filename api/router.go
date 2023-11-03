package router

import (
	"urlshortener/internal/handler" // Предполагаем, что обработчики находятся здесь

	"github.com/go-chi/chi"
)

func NewRouter(shortenerService interface{}) *chi.Mux {
	r := chi.NewRouter()

	// Регистрация обработчиков из пакета handler
	r.Get("/home", handler.HomePageHandler)
	r.Get("/redirect/{shortURL}", handler.RedirectHandler)

	// API маршруты
	r.Route("/api", func(r chi.Router) {
		r.Post("/shorten", handler.SaveURLHandler(shortenerService))
		r.Get("/resolve", handler.ResolveHandler(shortenerService))
		// Если у вас есть обработчики для следующих путей, раскомментируйте их
		// r.Post("/shorten/batch", handler.ShortenURLBatchHandler)
		// r.Delete("/user/urls", handler.DeleteURLBatchHandler)
	})

	// Служебные маршруты
	r.Get("/health", handler.HealthCheckHandler)
	r.Get("/ping", handler.PingHandler)

	// Включение отладочных маршрутов, если они у вас есть
	// if debug {
	// 	r.Route("/debug", func(r chi.Router) {
	// 		r.Get("/", handler.DebugProfileHandler)
	// 	})
	// }

	return r
}
