package shortenpkg

import (
	"database/sql" // Пакет для работы с SQL базами данных
	"net/http"     // Пакет для работы с HTTP сервером и клиентом

	_ "github.com/lib/pq" // Драйвер для PostgreSQL. Используется ниже при вызове sql.Open()
)

// Глобальная переменная для подключения к базе данных
var db *sql.DB

// Set записывает пару короткий URL и оригинальный URL в базу данных
func Set(shortKey, url string) error {
	_, err := db.Exec("INSERT INTO url_shortener (shortKey, url) VALUES ($1, $2)", shortKey, url)
	return err
}

// Get извлекает оригинальный URL из базы данных, используя короткий URL как ключ
func Get(shortKey string) (string, error) {
	var url string
	err := db.QueryRow("SELECT url FROM url_shortener WHERE shortKey = $1", shortKey).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

// ShortenHandler - обработчик HTTP-запроса для создания короткого URL из оригинального
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса POST
	if r.Method != http.MethodPost {
		http.Error(w, "ПРИВЕТ МИР !!!", http.StatusMethodNotAllowed)
		return
	}
	// Извлекаем оригинальный URL из формы запроса
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	// Генерируем короткий URL (функция GenerateShortURL не предоставлена в вашем коде, но, похоже, должна быть реализована)
	shortKey := GenerateShortURL()
	err := Set(shortKey, longURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}
	// Возвращаем сгенерированный короткий URL в ответ
	w.Write([]byte(shortKey))
}

// ResolveHandler - обработчик HTTP-запроса для получения оригинального URL по короткому
func ResolveHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем короткий URL из формы запроса
	shortKey := r.FormValue("key")
	if shortKey == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}
	// Извлекаем оригинальный URL из базы данных по короткому ключу
	longURL, err := Get(shortKey)
	if err != nil {
		http.Error(w, "Failed to get URL", http.StatusInternalServerError)
		return
	}
	// Возвращаем оригинальный URL в ответ
	w.Write([]byte(longURL))
}
