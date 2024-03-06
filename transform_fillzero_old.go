//go:build !go1.21
// +build !go1.21

package stl4go

// FillZero fills each element in slice a with zero value.
//
// Complexity: O(len(a)).
func FillZero[T any](a []T) {
	var zero T
	Fill(a, zero)
}
