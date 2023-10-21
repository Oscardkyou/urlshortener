// api/shortener.go
package apishorten // было shortenpkg

import (
	"net/http"
	"urlshortener/shortener"
	"urlshortener/storage"
)

var store = storage.NewMemoryStorage()
var shortenerService = shortener.NewShortenerService(store)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ПРИВЕТ МИР !!!", http.StatusMethodNotAllowed)
		return
	}
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortURL, err := shortenerService.Shorten(longURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(shortURL))
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := r.FormValue("key")
	if shortKey == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}

	longURL, err := shortenerService.Expand(shortKey)
	if err != nil {
		http.Error(w, "Failed to get URL", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(longURL))
}
