package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger is the package logging implementation
type ZapLogger struct {
	loggerImpl *zap.Logger
}

// NewZapLogger creates a logger at log level Error
func NewZapLogger(logConf LoggerConfig) ILogger {
	config := zap.NewDevelopmentConfig()

	logLevel, err := NewLogLevel(logConf.LogLevel)
	if err != nil {
		config.Level.SetLevel(zapcore.ErrorLevel)
	}

	config.Level.SetLevel(toZapLogLevel(logLevel))

	callerOpt := zap.AddCallerSkip(1)
	loggerImpl, _ := config.Build(callerOpt)

	return ZapLogger{
		loggerImpl: loggerImpl,
	}
}

func (l ZapLogger) Info(template string, args ...interface{}) {
	if len(args) == 0 {
		l.loggerImpl.Sugar().Info(template)
		return
	}

	l.loggerImpl.Sugar().Infof(template, args...)
}

func (l ZapLogger) Debug(template string, args ...interface{}) {
	if len(args) == 0 {
		l.loggerImpl.Sugar().Debug(template)
		return
	}

	l.loggerImpl.Sugar().Debugf(template, args...)
}

func (l ZapLogger) Error(template string, args ...interface{}) {
	if len(args) == 0 {
		l.loggerImpl.Sugar().Error(template)
		return
	}

	l.loggerImpl.Sugar().Errorf(template, args...)
}

func toZapLogLevel(logLevel LogLevel) zapcore.Level {
	switch logLevel {
	case Error:
		return zapcore.ErrorLevel
	case Info:
		return zapcore.InfoLevel
	case Debug:
		return zapcore.DebugLevel
	default:
		return zapcore.ErrorLevel
	}
}

// SetLogLevelInfo sets the current log level to info
func (l ZapLogger) SetLogLevelInfo() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.InfoLevel)

	l.loggerImpl, _ = config.Build()
}

// SetLogLevelDebug sets the current log level to debug
func (l ZapLogger) SetLogLevelDebug() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.DebugLevel)

	l.loggerImpl, _ = config.Build()
}

// SetLogLevelError sets the current log level to error
func (l ZapLogger) SetLogLevelError() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.ErrorLevel)

	l.loggerImpl, _ = config.Build()
}
