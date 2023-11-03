package server

import (
	"api/"
	"net/http"
	"urlshortener/internal/config"
	"urlshortener/internal/shortener"
	"urlshortener/internal/storage"

	"go.uber.org/zap"
)

func Initialize(cfg *config.Config, logger *zap.Logger) *Server {
	var store storage.StorageInterface
	switch cfg.StorageType {
	case "memory":
		store = storage.NewMemoryStorage()
	case "db":
		// TODO: Add DB storage implementation if available
		// store = storage.NewDBStorage()
		logger.Fatal("DB storage is not implemented yet")
	default:
		logger.Fatal("Unsupported storage type", zap.String("storageType", cfg.StorageType))
	}

	shortenerService := shortener.NewShortenerService(store)
	mux := api.NewRouter(shortenerService)

	return &Server{
		Port:   cfg.Port,
		Router: mux,
		Logger: logger,
	}
}

type Server struct {
	Port   string
	Router *http.ServeMux
	Logger *zap.Logger
}

func (s *Server) Start() {
	s.Logger.Info("Starting server", zap.String("port", s.Port))
	err := http.ListenAndServe(":"+s.Port, s.Router)
	if err != nil {
		s.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
