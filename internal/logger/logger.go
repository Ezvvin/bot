package logger

import (
	"bot/internal/domain"
	"fmt"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

// InitLogger - Функция для иницилизации `logrus` и конфигурации паттерна.
func InitLogger() {
	// Set logger formatter
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			functionName := path.Base(f.Function)
			return fmt.Sprintf("%s()", functionName), fmt.Sprintf(" %s:%d", filename, f.Line)
		},
	})
}

// ParseLogLevel - Функция для парсинга уровня логгера.
func ParseLogLevel(cfg domain.Config) log.Level {
	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatal("Level parse error: ", err)
	}

	return level
}
