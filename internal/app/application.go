package app

import (
	"urlshortener/internal/storage"
	"urlshortener/shortener"

	"go.uber.org/zap"
)

type Application struct {
	logger *zap.Logger
}

func NewApplication() *Application {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return &Application{
		logger: logger,
	}
}

func (a *Application) Run() error {
	// Здесь просто пример. Вы должны определить свой способ выбора типа хранилища.
	storageType := "memory"

	var store storage.StorageInterface
	switch storageType {
	case "memory":
		store = storage.NewMemoryStorage()
	case "db":
		// Добавьте вашу реализацию DBStorage здесь, если у вас она есть.
		a.logger.Error("DB storage is not implemented yet")
		return nil
	default:
		a.logger.Fatal("Unsupported storage type")
		return nil
	}