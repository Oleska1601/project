package logger

import (
	"log/slog"
	"os"
	"project/config"
	"strings"
)

type Logger struct {
	logger *slog.Logger
}

func New(cfg *config.Config) *Logger {
	logLevel := cfg.Log.Level
	var l slog.Level
	switch strings.ToLower(logLevel) {
	case "debug":
		l = slog.LevelDebug
	case "info":
		l = slog.LevelInfo
	case "warn":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}
	slog.SetLogLoggerLevel(l)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &Logger{logger: logger}
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args)
}

func (l *Logger) Error(err error, args ...any) {
	l.logger.Error(err.Error(), args)

}
