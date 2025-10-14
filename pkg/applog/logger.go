package applog

import (
	"fmt"
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
	lgr  *slog.Logger
	args []Arg
}

func NewApplicationLogger() *ApplicationLogger {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &ApplicationLogger{lgr: logger}
}

type Arg struct {
	Key   string
	Value any
}

func (al *ApplicationLogger) With(args ...Arg) Logger {
	newLgr := NewApplicationLogger()
	for _, arg := range args {
		newLgr.lgr = newLgr.lgr.With(arg.Key, arg.Value)
	}
	newLgr.args = args
	return newLgr
}

func (al *ApplicationLogger) Section(section string, operation string) Logger {
	newLgr := NewApplicationLogger()
	for _, arg := range al.args {
		newLgr.lgr = newLgr.lgr.With(arg.Key, arg.Value)
	}
	newLgr.args = append(
		newLgr.args,
		Arg{Key: "section", Value: section},
		Arg{Key: "operation", Value: operation},
	)
	return newLgr
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
