package handler

import (
	"GoUrlShortener/internal/store"
	"encoding/json"
	"net/http"
)

func SaveURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	shortKey := data["shortKey"]
	if shortKey == "" {
		shortKey = GenerateShortURL() // Ensure this function is implemented elsewhere
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"shortened": "{{.BaseURL}}" + shortKey})
}



	http.Redirect(w, r, url, http.StatusFound)
}
