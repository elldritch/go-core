package log

import (
	"context"
	"runtime"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func From(ctx context.Context) *Logger {
	reqLogger := zerolog.Ctx(ctx).With().Str("RequestID", middleware.GetReqID(ctx)).Logger()
	return &Logger{&reqLogger}
}

// TODO: alternatively, logging frame + 1 caller upwards? use pkg/file:line
func (l *Logger) Panic() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Panic().Str("file", file).Int("line", line)
	}
	return l.Logger.Panic()
}

func (l *Logger) Fatal() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Fatal().Str("file", file).Int("line", line)
	}
	return l.Logger.Fatal()
}

func (l *Logger) Error() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Error().Str("file", file).Int("line", line)
	}
	return l.Logger.Error()
}

func (l *Logger) Warn() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Warn().Str("file", file).Int("line", line)
	}
	return l.Logger.Warn()
}

func (l *Logger) Info() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Info().Str("file", file).Int("line", line)
	}
	return l.Logger.Info()
}

func (l *Logger) Debug() *zerolog.Event {
	if dbg {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			panic("unable to get caller")
		}
		return l.Logger.Debug().Str("file", file).Int("line", line)
	}
	return l.Logger.Debug()
}
