package handler

import (
	"encoding/json"
	"net/http"
	"urlshortener/internal/shortener"
)

type Response struct {
	ShortURL string `json:"short_url"`
}

type ResolveResponse struct {
	LongURL string `json:"long_url"`
}

// SaveURLHandler handles the creation of a short URL.
func SaveURLHandler(s *shortener.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var data map[string]string

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		longURL, ok := data["url"]
		if !ok {
			http.Error(w, "URL не предоставлен", http.StatusBadRequest)
			return
		}

		shortURL, err := s.Shorten(longURL)
		if err != nil {
			http.Error(w, "Ошибка при создании короткого URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := Response{ShortURL: shortURL}
		json.NewEncoder(w).Encode(response)
	}
}

// ResolveHandler handles the retrieval of the original URL from a short URL.
func ResolveHandler(s *shortener.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortKey := r.URL.Query().Get("key")
		if shortKey == "" {
			http.Error(w, "Ключ не предоставлен", http.StatusBadRequest)
			return
		}

		longURL, err := s.Expand(shortKey)
		if err != nil {
			http.Error(w, "Ошибка при расшифровке короткого URL", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := ResolveResponse{LongURL: longURL}
		json.NewEncoder(w).Encode(response)
	}
}
