package assert

import (
	"runtime/debug"
	"testing"
)

func NoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("failed to assert no error: (%v)\n(%s)", err, string(debug.Stack()))
	}
}

func Equal[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Fatalf("failed to assert (%v) is equal to (%v)\n(%v)", got, want, string(debug.Stack()))
	}
}

func Truef(t testing.TB, got bool, msg string) {
	t.Helper()

	if !got {
		t.Fatalf("%s\n(%v)", msg, string(debug.Stack()))
	}
}
