package main

import (
	"urlshortener/internal/config"
	"urlshortener/internal/logger"
	"urlshortener/internal/server"
)

// App инкапсулирует сервер и его зависимости
type App struct {
	server server.ServerInterface // Используйте интерфейс вместо конкретного типа
}

// Initialize настраивает приложение с необходимыми компонентами
func (a *App) Initialize(cfgLoader config.ConfigLoader, log logger.LoggerInterface) {
	cfg := cfgLoader.Load() // Загрузка конфигурации с использованием переданного загрузчика

	// Инициализация сервера с использованием загруженной конфигурации и логгера
	a.server = server.NewServer(cfg, log)
}

// Run запускает приложение
func (a *App) Run() {
	a.server.Start()
}

func main() {
	// Инициализация зависимостей
	cfgLoader := config.NewConfigLoader() // Предполагается, что это возвращает реализацию ConfigLoader
	log := logger.NewLogger()             // Предполагается, что это возвращает реализацию LoggerInterface

	app := &App{}
	app.Initialize(cfgLoader, log) // Внедрение зависимостей в приложение
	app.Run()                      // Запуск приложения
}
