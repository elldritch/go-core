package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	dbg    bool
	logger Logger
)

func Init(debug bool) {
	dbg = debug
	logger = Logger{&log.Logger}
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}
