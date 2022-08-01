package contalgo

import "math/rand"

// Copy make a copy of slice a.
//
// Complexity: O(len(a)).
func Copy[T any](a []T) []T {
	b := append([]T{}, a...)
	return b
}

// TransformTo applies the function op to each element in slice a and fill it to slice b.
//
// The len(b) must not lesser than len(a).
//
// Complexity: O(len(a)).
func TransformTo[R any, T any](a []T, op func(T) R, b []R) {
	if len(b) < len(a) {
		panic("TransformTo: len(b) < len(a)")
	}
	for i, v := range a {
		b[i] = op(v)
	}
}

// Transform applies the function op to each element in slice a and set it back to the same place in a.
//
// Complexity: O(len(a)).
func Transform[T any](a []T, op func(T) T) {
	for i, v := range a {
		a[i] = op(v)
	}
}

// TransformCopy applies the function op to each element in slice a and return all the result as a slice.
//
// Complexity: O(len(a)).
//
func TransformCopy[R any, T any](a []T, op func(T) R) []R {
	r := make([]R, 0, len(a))
	for _, v := range a {
		r = append(r, op(v))
	}
	return r
}

// Unique remove adjacent repeated elements from the input slice.
// return the processed slice, and the content of the input slice is also changed.
//
// Complexity: O(len(a)).
func Unique[T comparable](a []T) []T {
	if len(a) == 0 {
		return a
	}
	i := 1
	prev := a[0]
	for _, v := range a[1:] {
		if v != prev {
			a[i] = v
			i++
			prev = v
		}
	}
	return a[:i]
}

// UniqueCopy remove adjacent repeated elements from the input slice.
// return the result slice, and the input slice is kept unchanged.
//
// Complexity: O(len(a)).
func UniqueCopy[T comparable](a []T) []T {
	var r []T
	if len(a) == 0 {
		return r
	}

	for _, v := range a {
		if len(r) > 0 && r[len(r)-1] == v {
			continue
		}
		r = append(r, v)
	}

	return r
}

// Remove remove the elements which equals to x from the input slice.
// return the processed slice, and the content of the input slice is also changed.
//
// Complexity: O(len(a)).
func Remove[T comparable](a []T, x T) []T {
	j := 0
	for _, v := range a {
		if v != x {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// RemoveCopy remove all elements which equals to x from the input slice.
// return the processed slice, and the content of the input slice is also changed.
//
// Complexity: O(len(a)).
func RemoveCopy[T comparable](a []T, x T) []T {
	r := make([]T, 0, len(a))
	for _, v := range a {
		if v != x {
			r = append(r, v)
		}
	}
	return r
}

// RemoveIf remove each element which make cond(x) returns true from the input slice,
// copy other elements to a new slice and return it. The input slice is kept unchanged.
//
// Complexity: O(len(a)).
func RemoveIf[T any](a []T, cond func(T) bool) []T {
	j := 0
	for _, v := range a {
		if !cond(v) {
			a[j] = v
			j++
		}
	}
	return a[:j]
}

// RemoveIfCopy drops each element which make cond(x) returns true from the input slice,
// copy other elements to a new slice and return it. The input slice is kept unchanged.
//
// Complexity: O(len(a)).
func RemoveIfCopy[T any](a []T, cond func(T) bool) []T {
	r := make([]T, 0, len(a))
	for _, v := range a {
		if !cond(v) {
			r = append(r, v)
		}
	}
	return r
}

// Shuffle pseudo-randomizes the order of elements.
//
// Complexity: O(len(a)).
func Shuffle[T any](a []T) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

// Reverse reverses the order of the elements in the slice a.
//
// Complexity: O(len(a)).
func Reverse[T any](a []T) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

// ReverseCopy returns a reversed copy of slice a.
//
// Complexity: O(len(a)).
func ReverseCopy[T any](a []T) []T {
	b := make([]T, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}
