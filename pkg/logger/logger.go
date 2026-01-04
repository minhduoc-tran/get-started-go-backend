package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger interface định nghĩa các method logging
type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
}

// SimpleLogger implementation
type SimpleLogger struct {
	logger *log.Logger
}

// NewLogger - Tạo logger mới
func NewLogger() Logger {
	return &SimpleLogger{
		logger: log.New(os.Stdout, "", 0),
	}
}

func (l *SimpleLogger) Info(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[%s] [INFO] %s\n", timestamp, msg)
}

func (l *SimpleLogger) Error(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[%s] [ERROR] %s\n", timestamp, msg)
}

func (l *SimpleLogger) Debug(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[%s] [DEBUG] %s\n", timestamp, msg)
}

func (l *SimpleLogger) Warn(msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Printf("[%s] [WARN] %s\n", timestamp, msg)
}

// Global logger instance
var globalLogger Logger

// Init - Khởi tạo global logger
func Init() {
	globalLogger = NewLogger()
}

// GetLogger - Lấy global logger
func GetLogger() Logger {
	if globalLogger == nil {
		Init()
	}
	return globalLogger
}

// Convenience functions - Sử dụng global logger
func Info(msg string) {
	GetLogger().Info(msg)
}

func Error(msg string) {
	GetLogger().Error(msg)
}

func Debug(msg string) {
	GetLogger().Debug(msg)
}

func Warn(msg string) {
	GetLogger().Warn(msg)
}

// Formatted functions
func Infof(format string, args ...interface{}) {
	GetLogger().Info(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	GetLogger().Error(fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...interface{}) {
	GetLogger().Debug(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
	GetLogger().Warn(fmt.Sprintf(format, args...))
}
