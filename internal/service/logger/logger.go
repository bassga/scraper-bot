package logger

import (
	"log"

	"github.com/bassga/scraper-bot/internal/domain/logger"
)

type LoggerImpl struct{}

// Error implements logger.Logger.
func (l *LoggerImpl) Info(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

// Info implements logger.Logger.
func (l *LoggerImpl) Error(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}

func NewLogger() logger.Logger {
	return &LoggerImpl{}
}
