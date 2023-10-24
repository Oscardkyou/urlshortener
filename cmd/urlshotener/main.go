package main

import (
	"log"
	"net/http"
	apishorten "urlshortener/api"
	"urlshortener/shortener"
	"urlshortener/storage"
)

func main() {
	store := storage.NewMemoryStorage()
	shortenerService := shortener.NewShortenerService(store)

	// передаем shortenerService в обработчики
	http.HandleFunc("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		apishorten.SaveURLHandler(shortenerService, w, r)
	})
	http.HandleFunc("/api/resolve", func(w http.ResponseWriter, r *http.Request) {
		apishorten.ResolveHandler(shortenerService, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
