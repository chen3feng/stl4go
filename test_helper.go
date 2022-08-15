package stl4go

import (
	"fmt"
	"testing"
)

func report(t *testing.T, msg string) {
	t.Helper()
	t.Errorf("Wrong: %v\n", msg)
}

func reportMismatch[T comparable](t *testing.T, a T, op string, b T) {
	t.Helper()
	report(t, fmt.Sprintf("%v %v %v", a, op, b))
}

func expectEq[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a != b {
		reportMismatch(t, a, "==", b)
	}
}

func expectNe[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if !(a != b) {
		reportMismatch(t, a, "!=", b)
	}
}

func expectLt[T Ordered](t *testing.T, a, b T) {
	t.Helper()
	if !(a < b) {
		reportMismatch(t, a, "<", b)
	}
}

func expectGt[T Ordered](t *testing.T, a, b T) {
	t.Helper()
	if !(a > b) {
		reportMismatch(t, a, ">", b)
	}
}

func expectLe[T Ordered](t *testing.T, a, b T) {
	t.Helper()
	if !(a <= b) {
		reportMismatch(t, a, "<=", b)
	}
}

func expectGe[T Ordered](t *testing.T, a, b T) {
	t.Helper()
	if !(a >= b) {
		reportMismatch(t, a, ">=", b)
	}
}

func expectTrue(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		reportMismatch(t, actual, "==", true)
	}
}

func expectFalse(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		reportMismatch(t, actual, "==", false)
	}
}

func expactPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() {
		t.Helper()
		if r := recover(); r == nil {
			report(t, "din't panic")
		}
	}()

	f()
}
