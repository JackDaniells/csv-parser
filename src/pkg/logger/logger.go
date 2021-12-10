package logger

import "fmt"

const (
	LOGGER_LEVEL_DEBUG = "debug"
	LOGGER_LEVEL_INFO  = "info"
	LOGGER_LEVEL_WARN  = "warn"
	LOGGER_LEVEL_ERROR = "error"
)

type logger struct {
	level string
}

func Debug() *logger {
	return &logger{level: LOGGER_LEVEL_DEBUG}
}

func Info() *logger {
	return &logger{level: LOGGER_LEVEL_INFO}
}

func Warn() *logger {
	return &logger{level: LOGGER_LEVEL_WARN}
}

func Error() *logger {
	return &logger{level: LOGGER_LEVEL_ERROR}
}

func (l *logger) Log(message string) {
	fmt.Println(fmt.Sprintf("[level:%s] message: %s", l.level, message))
}
