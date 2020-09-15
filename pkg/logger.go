package ldsview

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger ILogger

func init() {
	Logger = PkgLogger{
		loggerImpl: zap.NewNop(),
	}
}

// EnableLogging turns on package logging
func EnableLogging() {
	Logger = NewLogger()
}

// ILogger is the package logging interface
type ILogger interface {
	Info(template string, args ...interface{})
	Debug(template string, args ...interface{})
	Error(template string, args ...interface{})

	SetLogLevelInfo()
	SetLogLevelDebug()
}

// PkgLogger is the package logging implementation
type PkgLogger struct {
	loggerImpl *zap.Logger
}

// NewLogger creates a logger at log level Error
func NewLogger() ILogger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.ErrorLevel)

	loggerImpl, _ := config.Build()
	return PkgLogger{
		loggerImpl: loggerImpl,
	}
}

func (l PkgLogger) Info(template string, args ...interface{}) {
	l.loggerImpl.Sugar().Infof(template, args)
}

func (l PkgLogger) Debug(template string, args ...interface{}) {
	l.loggerImpl.Sugar().Debugf(template, args)
}

func (l PkgLogger) Error(template string, args ...interface{}) {
	l.loggerImpl.Sugar().Errorf(template, args)
}

// SetLogLevelInfo sets the current log level to info
func (l PkgLogger) SetLogLevelInfo() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.InfoLevel)

	l.loggerImpl, _ = config.Build()
}

// SetLogLevelDebug sets the current log level to debug
func (l PkgLogger) SetLogLevelDebug() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.DebugLevel)

	l.loggerImpl, _ = config.Build()
}
