package main

import (
	"net/http"
	"urlshortener/api"
	"urlshortener/config" // добавьте этот импорт
	"urlshortener/shortener"
	"urlshortener/storage"

	"go.uber.org/zap"
)

func main() {
	// Инициализация логгера zap
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	store := storage.NewMemoryStorage()
	shortenerService := shortener.NewShortenerService(store)
	mux := api.NewRouter(shortenerService)

	// Используем logger от zap вместо стандартного log
	logger.Info("Starting server", zap.String("port", cfg.Port))

	err = http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
