package main

import (
	"GoUrlShortener_final/internal/pkg/shortener"
	"GoUrlShortener_final/internal/pkg/storage"
	"log"
	"net/http"
)

func main() {
	// Создаем экземпляр MemoryStorage
	store := storage.NewMemoryStorage()
	shortenerService := shortener.NewShortener(store)

	// Обработчики URL используют shortenerService
	http.HandleFunc("/api/shorten", shortenerService.ShortenHandler)
	http.HandleFunc("/api/resolve", shortenerService.ResolveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
