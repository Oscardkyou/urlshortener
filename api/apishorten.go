package apishorten

import (
	"net/http"
	"urlshortener/shortener"
)

func SaveURLHandler(s *shortener.ShortenerService, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "ПРИВЕТ МИР !!!", http.StatusMethodNotAllowed)
		return
	}
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL работает", http.StatusBadRequest)
		return
	}

	shortURL, err := s.Shorten(longURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(shortURL))
}

func ResolveHandler(s *shortener.ShortenerService, w http.ResponseWriter, r *http.Request) {
	shortKey := r.FormValue("key")
	if shortKey == "" {
		http.Error(w, "Ключ получен", http.StatusBadRequest)
		return
	}

	longURL, err := s.Expand(shortKey)
	if err != nil {
		http.Error(w, "Ошибка УРЛА(((())))", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(longURL))
}
