
package logger

import (
	"go.uber.org/zap"
)

func Initialize() *zap.Logger {
	// Initialize the zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}
