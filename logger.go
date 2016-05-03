package logging

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

var (
	DEBUG = Level{
		Value: 10,
		Name:  "DEBUG",
	}

	INFO = Level{
		Value: 20,
		Name:  "INFO",
	}

	WARN = Level{
		Value: 30,
		Name:  "WARN",
	}

	ERROR = Level{
		Value: 40,
		Name:  "ERROR",
	}

	PANIC = Level{
		Value: 50,
		Name:  "PANIC",
	}
)

// log level
type Level struct {
	// bigger number has bigger priority
	Value int

	// level name
	Name string
}

type Logger struct {
	// logger name
	Name string `json:"name"`

	Level Level `json:"-"`

	Logger *log.Logger
}

func GetDefaultLogger() *log.Logger {
	logFlag := log.LstdFlags | log.Lmicroseconds
	return log.New(os.Stdout, "[default] ", logFlag)
}

func (l *Logger) isVerbose(lvl Level) bool {
	if lvl.Value < l.Level.Value {
		return false
	}

	return true
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l.isVerbose(DEBUG) {
		l.Logger.Printf("["+DEBUG.Name+"] "+format, v...)
	}
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l.isVerbose(INFO) {
		l.Logger.Printf("["+INFO.Name+"] "+format, v...)
	}
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l.isVerbose(WARN) {
		l.Logger.Printf("["+WARN.Name+"] "+format, v...)
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l.isVerbose(ERROR) {
		l.Logger.Printf("["+ERROR.Name+"] "+format, v...)
	}
}

func (l *Logger) Trace(recovered interface{}, format string, v ...interface{}) {
	msg := fmt.Sprintf("["+ERROR.Name+"] "+format, v...)
	l.Logger.Printf("%s\npanic: %v\n%s", msg, recovered, debug.Stack())
}

func (l *Logger) Panic(format string, v ...interface{}) {
	if l.isVerbose(PANIC) {
		l.Logger.Printf("["+PANIC.Name+"] "+format, v...)
	}
	panic(fmt.Sprintf(format, v...))
}
