package contalgo

import "reflect"

// nameOfType return type name as a string.
func nameOfType[T any]() string {
	var t *T = nil
	return reflect.TypeOf(t).String()[1:]
}
