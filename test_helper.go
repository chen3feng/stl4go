package stl4go

import (
	"fmt"
	"runtime"
	"testing"
)

func report(t *testing.T, file string, line int, msg string) {
	t.Errorf("%v:%v: Wrong: %v\n", file, line, msg)
}

func reportMismatch[T comparable](t *testing.T, a T, op string, b T, file string, line int) {
	report(t, file, line, fmt.Sprintf("%v %v %v", a, op, b))
}

func expectEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "==", b, file, line)
	}
}

func expectNe[T comparable](t *testing.T, a, b T) {
	if !(a != b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "!=", b, file, line)
	}
}

func expectLt[T Ordered](t *testing.T, a, b T) {
	if !(a < b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "<", b, file, line)
	}
}

func expectGt[T Ordered](t *testing.T, a, b T) {
	if !(a > b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, ">", b, file, line)
	}
}

func expectLe[T Ordered](t *testing.T, a, b T) {
	if !(a <= b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "<=", b, file, line)
	}
}

func expectGe[T Ordered](t *testing.T, a, b T) {
	if !(a >= b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, ">=", b, file, line)
	}
}

func expectTrue(t *testing.T, actual bool) {
	if !actual {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, actual, "==", true, file, line)
	}
}

func expectFalse(t *testing.T, actual bool) {
	if actual {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, actual, "==", false, file, line)
	}
}

func expactPanic(t *testing.T, f func()) {
	_, file, line, _ := runtime.Caller(1)
	defer func() {
		if r := recover(); r == nil {
			report(t, file, line, "din't panic")
		}
	}()

	f()
}
