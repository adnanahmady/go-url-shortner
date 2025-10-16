package applog

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	writer := bytes.Buffer{}
	lgr := NewApplicationLoggerWithOptions(&writer, func(opts *slog.HandlerOptions) {
		opts.Level = slog.LevelDebug
	})
	scs := []struct {
		name string
		act  func(lgr Logger)
		exp  string
	}{
		{
			name: "given logger when debug data logged then should contain DEBUG",
			act:  func(lgr Logger) { lgr.Debug("here") },
			exp:  "DEBUG",
		},
		{
			name: "given logger when debug data logged with format then should contain formatted string",
			act:  func(lgr Logger) { lgr.Debugf("here %v", "foo") },
			exp:  "here foo",
		},
		{
			name: "given logger when info data logged then should contain INFO",
			act:  func(lgr Logger) { lgr.Info("here") },
			exp:  "INFO",
		},
		{
			name: "given logger when info data logged with format then should contain formatted string",
			act:  func(lgr Logger) { lgr.Infof("here %v", "foo") },
			exp:  "here foo",
		},
		{
			name: "given logger when error data logged then should contain ERROR",
			act:  func(lgr Logger) { lgr.Error("here") },
			exp:  "ERROR",
		},
		{
			name: "given logger when error data logged with format then should contain formatted string",
			act:  func(lgr Logger) { lgr.Errorf("here %v", "foo") },
			exp:  "here foo",
		},
		{
			name: "given logger when warn data logged then should contain WARN",
			act:  func(lgr Logger) { lgr.Warn("here") },
			exp:  "WARN",
		},
		{
			name: "given logger when warn data logged with format then should contain formatted string",
			act:  func(lgr Logger) { lgr.Warnf("here %v", "foo") },
			exp:  "here foo",
		},
		{
			name: "given logger when section data logged then should contain section and operation",
			act:  func(lgr Logger) { lgr.Section("my-section", "my-operation").Info("here") },
			exp:  "section=my-section operation=my-operation",
		},
		{
			name: "given logger when section data logged with format then should contain formatted string",
			act:  func(lgr Logger) { lgr.Section("my-section", "my-operation").Infof("here %v", "foo") },
			exp:  "here foo",
		},
	}
	for _, tt := range scs {
		t.Run(tt.name, func(t *testing.T) {
			tt.act(lgr)

			assertContains(t, writer.String(), tt.exp)
		})
	}
}

func assertContains(t testing.TB, given, want string) {
	t.Helper()

	if !strings.Contains(given, want) {
		t.Fatalf("failed to assert that (%v) contains (%v)", given, want)
	}
}
