package main

import (
	"log"
	"net/http"
	"os"
	"urlshortener/api" // уберем псевдоним "router"
	"urlshortener/shortener"
	"urlshortener/storage"
)

func main() {
	// Настройка порта через переменную окружения. По умолчанию 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store := storage.NewMemoryStorage()
	shortenerService := shortener.NewShortenerService(store)
	mux := api.NewRouter(shortenerService) // Используем "api" здесь

	log.Printf("Starting server on :%s...", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
