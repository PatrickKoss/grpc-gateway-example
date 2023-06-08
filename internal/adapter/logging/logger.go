package logging

import (
	nativeLogger "log"

	"go.uber.org/zap"
)

// logger to init zap logger and a lock for singleton.
type logger struct {
	sugar  *zap.SugaredLogger
	l      *zap.Logger
	fields []Field
}

type Field = zap.Field

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	Boolp      = zap.Boolp
	ByteString = zap.ByteString
	Int        = zap.Int
	Float64    = zap.Float64
	Float64p   = zap.Float64p
	Float32    = zap.Float32
	Float32p   = zap.Float32p
	Durationp  = zap.Durationp
	Any        = zap.Any
	String     = zap.String
)

// Logger should have a info and error log level.
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
}

// Info logs on info level.
func (l logger) Info(msg string, fields ...Field) {
	if l.sugar != nil {
		l.l.Info(msg, l.mergeFields(fields...)...)
	} else {
		nativeLogger.Println(msg)
	}
}

// Error logs on error level.
func (l logger) Error(msg string, fields ...Field) {
	if l.sugar != nil {
		l.l.Error(msg, l.mergeFields(fields...)...)
	} else {
		nativeLogger.Println(msg)
	}
}

// Fatal logs on fatal level.
func (l logger) Fatal(msg string, fields ...Field) {
	if l.sugar != nil {
		l.l.Fatal(msg, l.mergeFields(fields...)...)
	} else {
		nativeLogger.Fatalf(msg)
	}
}

// Debug logs on debug level if it is enabled.
func (l logger) Debug(msg string, fields ...Field) {
	if l.sugar != nil {
		l.l.Debug(msg, l.mergeFields(fields...)...)
	} else {
		nativeLogger.Println(msg)
	}
}

// New initializes a new logger.
func New(fields ...Field) Logger {
	l, err := zap.NewProduction()
	if err != nil {
		return logger{sugar: nil}
	}

	newLogger := logger{
		sugar:  l.Sugar(),
		l:      l,
		fields: fields,
	}

	return newLogger
}

func (l logger) mergeFields(fields ...Field) []Field {
	return append(l.fields, fields...)
}
