package server

import (
	"net/http"
	"urlshortener/internal/config"
	"urlshortener/internal/handler"
	"urlshortener/internal/logger"
	"urlshortener/internal/shortener"
	"urlshortener/internal/storage"

	"go.uber.org/zap"
)

// ServerInterface - интерфейс для сервера.
type ServerInterface interface {
	Start()
}

// Server - структура сервера.
type Server struct {
	Port   string
	Router *http.ServeMux
	Logger logger.LoggerInterface
}

// NewServer - функция для создания нового сервера.
func NewServer(cfg *config.Config, log logger.LoggerInterface) ServerInterface {
	var store storage.StorageInterface
	switch cfg.StorageType {
	case "memory":
		store = storage.NewMemoryStorage()
	case "db":
		// TODO: реализация хранения в БД
		log.Fatal("DB storage is not implemented yet")
	default:
		log.Fatal("Unsupported storage type", zap.String("storageType", cfg.StorageType))
	}

	shortenerService := shortener.NewShortenerService(store)
	mux := handler.NewRouter(shortenerService)

	return &Server{
		Port:   cfg.Port,
		Router: mux,
		Logger: log,
	}
}

// Start - метод для запуска сервера.
func (s *Server) Start() {
	s.Logger.Info("Starting server", zap.String("port", s.Port))
	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		s.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
