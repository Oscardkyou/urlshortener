package api

import (
	"encoding/json"
	"net/http"
	"urlshortener/shortener"
)

type Response struct {
	ShortURL string `json:"short_url"`
}

type ResolveResponse struct {
	LongURL string `json:"long_url"`
}

func SaveURLHandler(s *shortener.ShortenerService, w http.ResponseWriter, r *http.Request) {
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
		// Здесь можно добавить дополнительную обработку ошибок
		http.Error(w, "Ошибка при создании короткого URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Response{ShortURL: shortURL}
	json.NewEncoder(w).Encode(response)
}

func ResolveHandler(s *shortener.ShortenerService, w http.ResponseWriter, r *http.Request) {
	shortKey := r.FormValue("key")
	if shortKey == "" {
		http.Error(w, "Ключ не предоставлен", http.StatusBadRequest)
		return
	}

	longURL, err := s.Expand(shortKey)
	if err != nil {
		// Здесь можно добавить дополнительную обработку ошибок
		http.Error(w, "Ошибка при расшифровке короткого URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := ResolveResponse{LongURL: longURL}
	json.NewEncoder(w).Encode(response)
}

func NewRouter(svc *shortener.ShortenerService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
		SaveURLHandler(svc, w, r)
	})
	mux.HandleFunc("/api/resolve", func(w http.ResponseWriter, r *http.Request) {
		ResolveHandler(svc, w, r)
	})

	return mux
}
