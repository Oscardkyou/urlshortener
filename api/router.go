package api

import (
	"net/http"
	"urlshortener/shortener"
)

// NewRouter создает и возвращает новый HTTP роутер для API сокращения URL.
func NewRouter(svc *shortener.ShortenerService) *http.ServeMux {
	mux := http.NewServeMux()

	// Обработчик для сокращения URL.
	mux.HandleFunc("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		SaveURLHandler(svc, w, r)
	})

	// Обработчик для восстановления оригинального URL.
	mux.HandleFunc("/api/resolve", func(w http.ResponseWriter, r *http.Request) {
		ResolveHandler(svc, w, r) // Исправлено имя функции
	})

	return mux
}
