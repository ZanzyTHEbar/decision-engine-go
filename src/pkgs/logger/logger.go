package goflexprologger

import (
	"log/slog"
	"os"

	"github.com/ZanzyTHEbar/decision-engine-go/pkgs/config"

	"github.com/labstack/gommon/log"
)

var loggerLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

// Set up the logger based on the configuration
// Must be called before server is started
func InitLogger(config config.Config) {

	var handler slog.Handler

	logLevel, ok := loggerLevels[string(config.Logger.Level)]
	log.Debugf("Log Level Set to %v, with: %v", config.Logger.Level, logLevel)

	if !ok {
		log.Debugf("Invalid log level: %s, Using default value INFO", config.Logger.Level)
		logLevel = slog.LevelInfo
	}

	switch config.Logger.Style {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	case "text":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	case "dev":
		handler = NewDevHandler(os.Stdout, logLevel)
	default:
		log.Debugf("Invalid log style: %s, Using default value JSON", config.Logger.Style)
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}