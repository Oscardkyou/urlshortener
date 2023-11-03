package app

import (
	"urlshortener/internal/config"
	"urlshortener/internal/logger"
	"urlshortener/internal/server"
)

// Интерфейсы для зависимостей
type ConfigLoader interface {
	Load() *config.Config
}

type Logger interface {
	Info(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

type Server interface {
	Start()
}

type App struct {
	server Server
}

// Инициализация App с зависимостями в виде интерфейсов
func NewApp(cfgLoader ConfigLoader, log logger.LoggerInterface) *App {
	cfg := cfgLoader.Load()
	srv := server.NewServer(cfg, log) // Используйте NewServer, а не Initialize
	return &App{server: srv}
}

func (a *App) Run() {
	a.server.Start()
}

// Теперь в main.go вы создаете реальные экземпляры и передаете их в App
func main() {
	cfgLoader := config.NewConfigLoader() // Фабрика для ConfigLoader
	log := logger.NewLogger()             // Фабрика для Logger
	app := NewApp(cfgLoader, log)         // DI
	app.Run()
}
