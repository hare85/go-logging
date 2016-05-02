package logging

import (
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

type TestWriter struct {
	count int
}

var testWriter = &TestWriter{count: 0}

func (l *TestWriter) Write(p []byte) (n int, err error) {
	l.count += 1

	return l.count, nil
}

func GetTestLogger() *log.Logger {
	return log.New(io.MultiWriter(os.Stdout, testWriter), "[testLogger] ", log.LstdFlags|log.Lmicroseconds)
}

func cleanUpTest() {
	loggers = map[string]*Logger{}
	Default = &Logger{
		Name:   "default",
		Level:  DEBUG,
		Logger: GetDefaultLogger(),
	}

	testWriter.count = 0
}

func TestGetLogger(t *testing.T) {
	defer cleanUpTest()

	logger := GetLogger("logger")
	assert.NotEqual(t, logger.Name, "logger")
	assert.Equal(t, logger.Name, "default")

	SetLogger("logger", DEBUG, GetTestLogger())

	_logger := GetLogger("logger")
	assert.NotEqual(t, _logger.Name, "default")
	assert.Equal(t, _logger.Name, "logger")

	assert.NotEqual(t, logger, _logger)
}

func TestSetLevel(t *testing.T) {
	defer cleanUpTest()

	ok := SetLevel("logger", DEBUG)
	assert.Equal(t, ok, ErrLoggerNotFounded)

	SetLogger("logger", DEBUG, GetTestLogger())
	ok = SetLevel("logger", DEBUG)
	assert.NotEqual(t, ok, ErrLoggerNotFounded)
	assert.Equal(t, ok, nil)

	logger := GetLogger("logger")
	logger.Debug("1")
	logger.Debug("2")
	logger.Debug("3")

	assert.Equal(t, 3, testWriter.count)

	testWriter.count = 0

	SetLevel("logger", INFO)
	logger.Debug("no count")
	logger.Debug("no count")
	logger.Debug("no count")
	logger.Warn("count")
	logger.Info("count")
	logger.Error("count")

	assert.Equal(t, 3, testWriter.count)

}
