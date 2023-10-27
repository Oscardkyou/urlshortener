package main

import (
	"flag"
	"net/http"
	"urlshortener/api"
	"urlshortener/config"
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

	// Используйте флаги для конфигурации
	var port string
	var storageType string

	flag.StringVar(&port, "port", "8080", "port to run the server on")
	flag.StringVar(&storageType, "storage", "memory", "type of storage ('memory' or 'db')")
	flag.Parse()

	if port == "" { // Если порт не указан, загрузите его из конфигурации
		cfg, err := config.Load()
		if err != nil {
			logger.Fatal("Failed to load configuration", zap.Error(err))
		}
		port = cfg.Port
	}

	var store storage.StorageInterface
	switch storageType {
	case "memory":
		store = storage.NewMemoryStorage()
	case "db":
		// Например, если у вас есть реализация DBStorage
		// store = storage.NewDBStorage()
		logger.Fatal("DB storage is not implemented yet")
	default:
		logger.Fatal("Unsupported storage type", zap.String("storageType", storageType))
	}

	shortenerService := shortener.NewShortenerService(store)
	mux := api.NewRouter(shortenerService)

	logger.Info("Starting server", zap.String("port", port))

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
