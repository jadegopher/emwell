package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type ILogger interface {
	InfoKV(ctx context.Context, msg string, kvs ...interface{})
	WarnKV(ctx context.Context, msg string, kvs ...interface{})
	ErrorKV(ctx context.Context, msg string, kvs ...interface{})
}

const (
	incorrectLogText = "Key values incorrect in log"
)

type Logger struct {
	zapLogger *zap.Logger
}

func NewLogger() (*Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Logger{
		zapLogger: zapLogger,
	}, nil
}

func (l *Logger) InfoKV(_ context.Context, msg string, _ ...interface{}) {
	l.zapLogger.Info(msg)
}

func (l *Logger) WarnKV(_ context.Context, msg string, _ ...interface{}) {
	l.zapLogger.Warn(msg)
}

func (l *Logger) ErrorKV(_ context.Context, msg string, args ...interface{}) {
	fields, err := argsToFields(args...)
	if err != nil {
		l.zapLogger.Error(incorrectLogText, zap.Error(err))
		l.zapLogger.Error(msg)
		return
	}

	l.zapLogger.Error(msg, fields...)
}

func argsToFields(args ...interface{}) ([]zap.Field, error) {
	list := append([]interface{}{}, args...)
	if len(list)%2 != 0 {
		return nil, fmt.Errorf("incorrect count of kv: %d", len(list))
	}

	result := make([]zap.Field, 0, len(list))
	key := ""
	for i := range list {
		if i%2 != 0 {
			result = append(result, zap.Any(key, list[i]))
			continue
		}

		val, ok := list[i].(string)
		if !ok {
			return nil, fmt.Errorf("key on position %d should be type string", i)
		}
		key = val
	}

	return result, nil
}
