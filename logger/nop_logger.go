package logger

// NopLogger is a logger impl that does absolutely nothing
type NopLogger struct{}

// NewNopLogger creates a logger at log level Error
func NewNopLogger() ILogger {
	return NopLogger{}
}

func (l NopLogger) Info(template string, args ...interface{}) {
}

func (l NopLogger) Debug(template string, args ...interface{}) {
}

func (l NopLogger) Error(template string, args ...interface{}) {
}

// SetLogLevelInfo sets the current log level to info
func (l NopLogger) SetLogLevelInfo() {
}

// SetLogLevelDebug sets the current log level to debug
func (l NopLogger) SetLogLevelDebug() {
}
