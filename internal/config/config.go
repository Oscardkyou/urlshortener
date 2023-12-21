package config

import "flag"

// Config - структура конфигурации для приложения.
type Config struct {
	Port        string
	StorageType string
}

// ConfigLoader - интерфейс для загрузчика конфигурации.
type ConfigLoader interface {
	Load() *Config
}

// NewConfigLoader - функция для создания нового загрузчика конфигурации.
func NewConfigLoader() ConfigLoader {
	return &defaultConfigLoader{}
}

type defaultConfigLoader struct{}

// Load - метод для загрузки конфигурации.
func (dl *defaultConfigLoader) Load() *Config {
	var port string
	var storageType string

	flag.StringVar(&port, "port", "8080", "port to run the server on")
	flag.StringVar(&storageType, "storage", "memory", "type of storage ('memory' or 'db')")
	flag.Parse()

	return &Config{
		Port:        port,
		StorageType: storageType,
	}
}
