package gorest

import (
	// "errors"
)

const (
	LoggerLevelDebug = 0
	LoggerLevelInfo = 1
	LoggerLevelWarn = 2
	LoggerLevelError = 3
	LoggerLevelFatal = 4
)

type Logger interface {
	SetLevel(level uint)
	GetLevel() (uint)

	Debug(component string, message string, err error)
	Info(component string, message string, err error)
	Warn(component string, message string, err error)
	Error(component string, message string, err error)
	Fatal(component string, message string, err error)
}

type HasLogger interface {
	GetLogger() (Logger)
	SetLogger(log Logger)
}

type HasName interface {
	GetName() (string)
	SetName(name string)
}