package logger

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

// Error exibe detalhes do erro
func Error(args ...interface{}) {
	log.Error(args...)
}

// ErrorContext exibe detalhes do erro com o contexto
func ErrorContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Error(args...)
}

// Info exibe detalhes do log info
func Info(args ...interface{}) {
	log.Info(args...)
}

// InfoContext exibe detalhes do log info com o contexto
func InfoContext(ctx context.Context, args ...interface{}) {
	log.WithContext(ctx).Info(args...)
}

// Fatal exibe detalhes do erro
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}
