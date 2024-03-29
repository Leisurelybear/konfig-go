package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

var (
	Logger = NewDefaultLogger()
)

type LogLevel int

const (
	Info  LogLevel = iota // 0
	Warn                  // 1
	Debug                 // 2
	Error                 // 3
)

type ILog interface {
	Info(ctx context.Context, message string, args ...interface{})
	Warn(ctx context.Context, message string, args ...interface{})
	Debug(ctx context.Context, message string, args ...interface{})
	Error(ctx context.Context, message string, args ...interface{})
}

type DefaultLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
}

func NewDefaultLogger() *DefaultLogger {
	infoLogger := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger := log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Llongfile)
	debugLogger := log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Llongfile)
	errorLogger := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile)

	return &DefaultLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		debugLogger: debugLogger,
		errorLogger: errorLogger,
	}
}

func (l *DefaultLogger) Info(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, Info, message, args...)
}

func (l *DefaultLogger) Warn(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, Warn, message, args...)
}

func (l *DefaultLogger) Debug(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, Debug, message, args...)
}

func (l *DefaultLogger) Error(ctx context.Context, message string, args ...interface{}) {
	l.log(ctx, Error, message, args...)
}

func (l *DefaultLogger) log(ctx context.Context, level LogLevel, message string, args ...interface{}) {
	logMessage := fmt.Sprintf(message, args...)
	switch level {
	case Info:
		l.infoLogger.Println(logMessage)
	case Warn:
		l.warnLogger.Println(logMessage)
	case Debug:
		l.debugLogger.Println(logMessage)
	case Error:
		l.errorLogger.Println(logMessage)
	}
}
