package logger

import "go.uber.org/zap"

// LoggerInterface - интерфейс для логгера.
type LoggerInterface interface {
	Info(msg string, keysAndValues ...interface{})
	Fatal(msg string, keysAndValues ...interface{})
}

// NewLogger - функция для создания нового логгера.
func NewLogger() LoggerInterface {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err) // или обработать ошибку более изящно
	}
	return logger
}
