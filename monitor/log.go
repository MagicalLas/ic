package monitor

import "log"

type Logger interface {
	standardLogger

	// Debug level log for debug in local or
	// development env.
	Debug(format string, args ...any)

	// Info level log for debug in production.
	Info(format string, args ...any)
	// Warn level log for warning state.
	// Something wrong, but it's need to attention.
	Warn(format string, args ...any)
	// Error level log for error state. It's critical,
	// need immediately fix this problem.
	Error(format string, args ...any)
}

// standardLogger must have same interface as log.Logger.
type standardLogger interface {

	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...any)

	// Fatalf must have the same semantics as log.Fatalf.
	Fatalf(format string, v ...any)

	// Panicf must have the same semantics as log.Panicf.
	Panicf(format string, v ...any)
}

// KLogger will log to kafka.
type KLogger struct {
	*log.Logger
}

func (l *KLogger) Debug(format string, args ...any) {
	l.Logger.Printf("DEBUG: "+format, args)
}

func (l *KLogger) Info(format string, args ...any) {
	l.Logger.Printf("INFO: "+format+" %v", args)
}

func (l *KLogger) Warn(format string, args ...any) {
	l.Logger.Printf("WARN: "+format, args)
}

func (l *KLogger) Error(format string, args ...any) {
	l.Logger.Printf("ERROR: "+format, args)
}

func New() Logger {
	return &KLogger{
		Logger: log.Default(),
	}
}
