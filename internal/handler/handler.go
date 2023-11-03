package handler

import (
	"encoding/json"
	"net/http"
	"urlshortener/internal/shortener"

	"github.com/go-chi/chi"
)

func NewRouter(s *shortener.ShortenerService) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/shorten", SaveURLHandler(s))
	router.Get("/{shortURL}", RedirectHandler(s))
	router.Get("/resolve", ResolveHandler(s))
	router.Get("/health", HealthCheckHandler)
	router.Get("/ping", PingHandler)
	router.Get("/", HomePageHandler)

	return router
}

func NewConfigLoader() {

}

type Response struct {
	ShortURL string `json:"short_url"`
}

type ResolveResponse struct {
	LongURL string `json:"long_url"`
}

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

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) //статус 200 OK
	w.Write([]byte("OK"))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong")) //тест
}

// HomePageHandler может отображать домашнюю страницу или какой-то информационный контент
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь может быть логика для отображения домашней страницы
	w.Write([]byte("Welcome to the home page!"))
}

// RedirectHandler обрабатывает перенаправление с короткого URL на исходный URL
func RedirectHandler(s *shortener.ShortenerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := chi.URLParam(r, "shortURL")
		longURL, err := s.Expand(shortURL)
		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
