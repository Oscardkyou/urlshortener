package main

import (
	"urlshortener/api"
	"urlshortener/app"
	"urlshortener/internal/storage"
)

func main() {
	application := app.NewApplication()
	err := application.Run()
	if err != nil {
		// обработка ошибки (можно использовать ваш логгер)
	}
}

func initLogger() {
	// Инициализация логгера
}

func configureServerFlags() {
	// Конфигурация флагов
}

func loadConfigurationIfNeeded() {
	// Загрузка конфигурации, если нужно
}

func selectStorageType() storage.StorageInterface {
	// Выбор типа хранилища
	return nil // пока возвращает nil; необходимо добавить реальную реализацию
}

func startServer(mux *api.RouterType) { // замените api.RouterType на реальный тип вашего роутера
	// Запуск сервера
}
