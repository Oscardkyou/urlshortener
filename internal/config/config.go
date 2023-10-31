package config

import (
	"flag"
	"net/http"
	"urlshortener/internal/api"
	"urlshortener/internal/shortener"
	"urlshortener/internal/storage"

	"go.uber.org/zap"
	"honnef.co/go/tools/config"
)

var (
	logger      *zap.Logger
	port        *string
	storageType *string
)

func initLogger() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

func configureServerFlags() {
	port = flag.String("port", "8080", "port to run the server on")
	storageType = flag.String("storage", "memory", "type of storage ('memory' or 'db')")
	flag.Parse()
}

func loadConfigurationIfNeeded() {
	if *port == "" {
		cfg, err := config.Load()
		if err != nil {
			logger.Fatal("Failed to load configuration", zap.Error(err))
		}
		*port = cfg.Port
	}
}

func selectStorageType() storage.StorageInterface {
	switch *storageType {
	case "memory":
		return storage.NewMemoryStorage()
	case "db":
		// Если у вас будет реализация DBStorage, замените следующую строку
		// return storage.NewDBStorage()
		logger.Fatal("DB storage is not implemented yet")
	default:
		logger.Fatal("Unsupported storage type", zap.String("storageType", *storageType))
		return nil
	}
}

func startServer(mux *chi.Mux) { // Предполагая, что api.NewRouter возвращает *chi.Mux
	logger.Info("Starting server", zap.String("port", *port))
	err := http.ListenAndServe(":"+*port, mux)
	if err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}

func main() {
	initLogger()
	configureServerFlags()
	loadConfigurationIfNeeded()

	store := selectStorageType()
	shortenerService := shortener.NewShortenerService(store)
	mux := api.NewRouter(shortenerService)

	startServer(mux)
}
