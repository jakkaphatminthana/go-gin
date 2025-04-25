package utils

import (
	"log"
	"os"
)

type Logger interface {
	Info(args ...any)
	Error(args ...any)
	Warn(args ...any)
	Debug(args ...any)

	Infof(format string, args ...any)
	Errorf(format string, args ...any)
}

type appLogger struct {
	logger *log.Logger
	level  string
}

func newLogger() *appLogger {
	return &appLogger{
		logger: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile),
		level:  "INFO",
	}
}

func (l *appLogger) Info(args ...any) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Println(args...)
}

func (l *appLogger) Error(args ...any) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Println(args...)
}

func (l *appLogger) Warn(args ...any) {
	l.logger.SetPrefix("[WARN] ")
	l.logger.Println(args...)
}

func (l *appLogger) Debug(args ...any) {
	l.logger.SetPrefix("[DEBUG] ")
	l.logger.Println(args...)
}

func (l *appLogger) Infof(format string, args ...any) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Printf(format, args...)
}

func (l *appLogger) Errorf(format string, args ...any) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Printf(format, args...)
}

// Global instance
var defaultLogger = newLogger()

// Exported global log functions
func Info(args ...any)                  { defaultLogger.Info(args...) }
func Error(args ...any)                 { defaultLogger.Error(args...) }
func Warn(args ...any)                  { defaultLogger.Warn(args...) }
func Debug(args ...any)                 { defaultLogger.Debug(args...) }
func Infof(format string, args ...any)  { defaultLogger.Infof(format, args...) }
func Errorf(format string, args ...any) { defaultLogger.Errorf(format, args...) }

func GetLogger() Logger {
	return defaultLogger
}

//> example usage
// func main() {
// 	logger := utils.GetLogger()
// 	logger.Info("This is an info message")
