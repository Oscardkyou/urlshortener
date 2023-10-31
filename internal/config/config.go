
package config

import (
	"flag"
	"go.uber.org/zap"
)

func Load() *Config {
	var port string
	var storageType string

	flag.StringVar(&port, "port", "8080", "port to run the server on")
	flag.StringVar(&storageType, "storage", "memory", "type of storage ('memory' or 'db')")
	flag.Parse()

	if port == "" {
		// TODO: Add logic to load port from external configuration if needed
		// cfg, err := externalConfigLoadFunction()
		// if err != nil {
		// 	zap.L().Fatal("Failed to load configuration", zap.Error(err))
		// }
		// port = cfg.Port
	}

	return &Config{
		Port:        port,
		StorageType: storageType,
	}
}

type Config struct {
	Port        string
	StorageType string
}
