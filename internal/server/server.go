package server

import (
	"net/http"
	"urlshortener/internal/config"
	"urlshortener/internal/handler"
	"urlshortener/internal/logger"
	"urlshortener/internal/shortener"
	"urlshortener/internal/storage"

	"github.com/go-chi/chi"
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

func NewServer(cfg *config.Config, log logger.LoggerInterface) ServerInterface {
	// ... [оставьте предыдущий код]

	shortenerService := shortener.NewShortenerService(storage.NewMemoryStorage())
	mux := handler.NewRouter(shortenerService)

	return &Server{
		Port:   cfg.Port,
		Router: chi.Mux,
		Logger: log,
	}
}

// Start - метод для запуска сервера.
func (s *Server) Start() {
	s.Logger.Info("Starting server on port " + s.Port)
	if err := http.ListenAndServe(":"+s.Port, s.Router); err != nil {
		s.Logger.Fatal("Server failed to start", zap.Error(err))
	}
}
