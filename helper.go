package stl4go

import (
	"reflect"
	"unsafe"
)

// nameOfType return type name as a string.
func nameOfType[T any]() string {
	var t *T
	return reflect.TypeOf(t).String()[1:]
}

// pointerCast case a pointer from *F to *T regardlessly.
func pointerCast[R *T, T any, F any](f *F) R {
	//#nosec G103
	return (R)(unsafe.Pointer(f))
}
