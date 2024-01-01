package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type LoggerInterface interface {
	GetZapLogger() *zap.Logger
	Named(s string) *Logger
	With(fields ...zap.Field) *Logger
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Panic(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
	extract(ctx context.Context) []zap.Field
}

var _ LoggerInterface = (*Logger)(nil)

type Logger struct {
	logger *zap.Logger
}

func New() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{
		logger: logger,
	}, nil
}

func (l *Logger) GetZapLogger() *zap.Logger {
	return l.logger
}

func (l *Logger) Named(s string) *Logger {
	l2 := l.logger.Named(s)
	return &Logger{l2}
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	l2 := l.logger.With(fields...)
	return &Logger{l2}
}

func (l *Logger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Debug(msg, fields...)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Info(msg, fields...)
}

func (l *Logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Warn(msg, fields...)
}

func (l *Logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Error(msg, fields...)
}

func (l *Logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Panic(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.logger.With(l.extract(ctx)...).Fatal(msg, fields...)
}

func (l *Logger) extract(ctx context.Context) []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.String("time", time.Now().Format(time.RFC3339)))
	return fields
}
