package applog

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Logger interface {
	Info(message string, args ...any)
	Error(message string, args ...any)
	Debug(message string, args ...any)
	Warn(message string, args ...any)
	Infof(message string, args ...any)
	Errorf(message string, args ...any)
	Debugf(message string, args ...any)
	Warnf(message string, args ...any)
	With(args ...Arg) Logger
	Section(section string, operation string) Logger
}

type ApplicationLogger struct {
	lgr    *slog.Logger
	args   []Arg
	writer io.Writer
}

func NewWriter() io.Writer {
	return os.Stdout
}

func NewApplicationLogger(w io.Writer) *ApplicationLogger {
	return NewApplicationLoggerWithOptions(w)
}

func NewApplicationLoggerWithOptions(
	w io.Writer,
	options ...func(*slog.HandlerOptions),
) *ApplicationLogger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	for _, opt := range options {
		opt(opts)
	}
	logger := slog.New(slog.NewTextHandler(w, opts))
	return &ApplicationLogger{lgr: logger, writer: w}
}

type Arg struct {
	Key   string
	Value any
}

func (al *ApplicationLogger) With(args ...Arg) Logger {
	newLgr := NewApplicationLogger(al.writer)
	for _, arg := range args {
		newLgr.lgr = newLgr.lgr.With(arg.Key, arg.Value)
	}
	newLgr.args = args
	return newLgr
}

func (al *ApplicationLogger) Section(section string, operation string) Logger {
	args := append(
		al.args,
		Arg{Key: "section", Value: section},
		Arg{Key: "operation", Value: operation},
	)

	return al.With(args...)
}

func (al *ApplicationLogger) Debug(message string, args ...any) {
	al.lgr.Debug(message, args...)
}

func (al *ApplicationLogger) Error(message string, args ...any) {
	al.lgr.Error(message, args...)
}

func (al *ApplicationLogger) Warn(message string, args ...any) {
	al.lgr.Warn(message, args...)
}

func (al *ApplicationLogger) Info(message string, args ...any) {
	al.lgr.Info(message, args...)
}

func (al *ApplicationLogger) Debugf(message string, args ...any) {
	al.lgr.Debug(fmt.Sprintf(message, args...))
}

func (al *ApplicationLogger) Errorf(message string, args ...any) {
	al.lgr.Error(fmt.Sprintf(message, args...))
}

func (al *ApplicationLogger) Warnf(message string, args ...any) {
	al.lgr.Warn(fmt.Sprintf(message, args...))
}

func (al *ApplicationLogger) Infof(message string, args ...any) {
	al.lgr.Info(fmt.Sprintf(message, args...))
}
