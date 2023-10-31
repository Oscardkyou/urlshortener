package api

import (
	"encoding/json"
	"net/http"
	"urlshortener/shortener"
)

type URLHandler struct {
	Service *shortener.ShortenerService
}

type Response struct {
	ShortURL string `json:"short_url"`
}

type ResolveResponse struct {
	LongURL string `json:"long_url"`
}

func (h *URLHandler) SaveURL(w http.ResponseWriter, r *http.Request) {
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

	shortURL, err := h.Service.Shorten(longURL)
	if err != nil {
		http.Error(w, "Ошибка при создании короткого URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := Response{ShortURL: shortURL}
	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) Resolve(w http.ResponseWriter, r *http.Request) {
	shortKey := r.FormValue("key")
	if shortKey == "" {
		http.Error(w, "Ключ не предоставлен", http.StatusBadRequest)
		return
	}

	longURL, err := h.Service.Expand(shortKey)
	if err != nil {
		http.Error(w, "Ошибка при расшифровке короткого URL", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := ResolveResponse{LongURL: longURL}
	json.NewEncoder(w).Encode(response)
}
