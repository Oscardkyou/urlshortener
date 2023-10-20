package shortenpkg

import (
	"encoding/json"
	"net/http"
	"urlshortener/shortener"
)

var shortenerService *shortener.ShortenerService // предположим, что сервис был инициализирован ранее

type Response struct {
	ShortURL string `json:"short_url"`
}

func SaveURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HELLO WORLD", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	longURL, ok := data["url"]
	if !ok {
		http.Error(w, "URL not provided", http.StatusBadRequest)
		return
	}

	shortURL, err := shortenerService.Shorten(longURL)
	if err != nil {
		http.Error(w, "Error shortening URL", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок ответа на "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Отправляем короткий URL в ответе
	response := Response{ShortURL: shortURL}
	json.NewEncoder(w).Encode(response)
}
