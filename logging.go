package logging

import (
	"errors"
	"log"
)

var (
	Default *Logger
	loggers map[string]*Logger
)

// errors
var (
	ErrLoggerNotFounded = errors.New("Logger name is not in loggers")
)

func init() {
	loggers = map[string]*Logger{}
	Default = &Logger{
		Name:   "default",
		Level:  DEBUG,
		Logger: GetDefaultLogger(),
	}
}

func GetLogger(name string) *Logger {
	logger, ok := loggers[name]
	if !ok {
		if name != "default" {
			Default.Warn("cannot find logger \"" + name + "\" - return to default logger")
		}
		return Default
	}

	return logger
}

func SetLogger(name string, level Level, logger *log.Logger) {
	l, ok := loggers[name]
	if ok {
		l.Level = level
		l.Logger = logger
	} else {
		l = &Logger{
			Name:   name,
			Level:  level,
			Logger: logger,
		}
	}

	loggers[l.Name] = l
}

func SetLevel(name string, level Level) error {
	logger, ok := loggers[name]
	if !ok {
		Default.Error("cannot find logger " + name)
		return ErrLoggerNotFounded
	}

	logger.Level = level
	return nil
}
