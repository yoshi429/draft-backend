package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

type logger struct {
	zap   *zap.Logger
	sugar *zap.SugaredLogger
}

func (l logger) Debug(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l logger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l logger) Warn(msg string, fields ...zap.Field) {
	l.zap.Warn(msg, fields...)
}

func (l logger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func New() Logger {
	zap, err := zap.NewDevelopment(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("ERROR AT Init Logger %s\n", err.Error())
		panic(err)
	}
	return logger{
		zap:   zap,
		sugar: zap.Sugar(),
	}
}
