package logger

import "errors"

// ILogger is the package logging interface
type ILogger interface {
	Info(template string, args ...interface{})
	Debug(template string, args ...interface{})
	Error(template string, args ...interface{})

	SetLogLevelInfo()
	SetLogLevelDebug()
}

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Error LogLevel = "error"
)

func NewLogLevel(logLevelStr string) (LogLevel, error) {
	logLevel := LogLevel(logLevelStr)
	switch logLevel {
	case Debug, Info, Error:
		return logLevel, nil
	}

	return logLevel, errors.New("Invalid log level: " + logLevelStr)
}

type LoggerConfig struct {
	LogLevel    string
	LogFilePath string
}
