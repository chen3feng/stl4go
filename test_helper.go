package contalgo

import (
	"runtime"
	"testing"

	"golang.org/x/exp/constraints"
)

func reportMismatch[T comparable](t *testing.T, a T, op string, b T, file string, line int) {
	if a != b {
		t.Errorf("%v:%v: Wrong: %v %v %v\n", file, line, a, op, b)
	}
}

func expectEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "==", b, file, line)
	}
}

func expectLt[T constraints.Ordered](t *testing.T, a, b T) {
	if !(a < b) {
		_, file, line, _ := runtime.Caller(1)
		reportMismatch(t, a, "<", b, file, line)
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
