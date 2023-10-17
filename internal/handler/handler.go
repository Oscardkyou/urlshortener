package shortenpkg

import (
	"encoding/json"
	"net/http"
)

// SaveURLHandler обрабатывает HTTP-запросы для сохранения длинного URL и возвращает короткий URL.
func SaveURLHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса - POST. Если нет, возвращаем ошибку.
	if r.Method != http.MethodPost {
		http.Error(w, "HELLO WORLD", http.StatusMethodNotAllowed)
		return
	}

	// Создаем словарь для декодирования JSON из тела запроса
	var data map[string]string

	// Декодируем JSON из тела запроса в словарь data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Задаем базовый URL для вашего сократителя ссылок
	baseURL := "http://yandex.ru/" // Замените на ваш базовый URL

	// Создаем словарь с сокращенной ссылкой для ответа
	responseData := map[string]string{"shortened": baseURL}

	// Устанавливаем заголовок ответа на "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Кодируем словарь в JSON и отправляем его в ответе
	json.NewEncoder(w).Encode(responseData)
}
