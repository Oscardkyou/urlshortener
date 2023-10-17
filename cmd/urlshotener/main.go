package main

import (
    "GoUrlShortener_final\shortener\shortener.go"
	// Импортируем необходимые пакеты
	"database/sql" // Пакет для работы с SQL базами данных
	"log"          // Пакет для логирования
	"net/http"     // Пакет для работы с HTTP сервером и клиентом
	"os"           // Пакет для работы с операционной системой

	_ "github.com/lib/pq" // Драйвер для PostgreSQL. Используется ниже при вызове sql.Open()
)

// Глобальная переменная для подключения к базе данных
var db *sql.DB

func main() {
	// Используем переменную окружения для безопасного хранения строк соединения с базой данных
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set") // Если переменная окружения не установлена, выводим ошибку и завершаем выполнение
	}

	// Пытаемся создать соединение с базой данных
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) // Если возникла ошибка при соединении, выводим ошибку и завершаем выполнение
	}
	// Закрываем соединение с базой данных при завершении работы программы
	defer db.Close()

	// Проверяем соединение с базой данных
	if err = db.Ping(); err != nil {
		log.Fatal(err) // Если соединение не установлено, выводим ошибку и завершаем выполнение
	}

	// Устанавливаем обработчики URL для нашего API
	http.HandleFunc("/api/shorten", shortener.ShortenHandler) // Обработчик для сокращения URL
	http.HandleFunc("/api/resolve", shortener.ResolveHandler) // Обработчик для разрешения сокращенного URL

	// Запускаем HTTP сервер на порту 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
