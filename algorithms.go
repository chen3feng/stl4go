package contalgo

import (
	"golang.org/x/exp/constraints"
)

//////////////////////////////////////////////////////////////////////////////
// Generative functions

// Range make a []T filled with values in the `[first, last)` sequence
func Range[T Numeric](first, last T) []T {
	a := make([]T, 0, int(last-first))
	for v := first; v < last; v++ {
		a = append(a, v)
	}
	return a
}

// Generate fill each element of `a`` with `gen()``
func Generate[T Numeric](a []T, gen func() T) {
	for i := range a {
		a[i] = gen()
	}
}

func Transform[R any, T any](a []T, op func(T) R) []R {
	r := make([]R, 0, len(a))
	for _, v := range a {
		r = append(r, op(v))
	}
	return r
}

//////////////////////////////////////////////////////////////////////////////
// Min, Max

// Max return the larger value between `a` and `b`
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Max return the smaller value between `a` and `b`
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxN return the maximum value in the sequence `a`
func MaxN[T constraints.Ordered](a ...T) T {
	if len(a) == 0 {
		panic("can't call MaxN() with empty arguments list")
	}
	v := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > v {
			v = a[i]
		}
	}
	return v
}

// MaxN return the minimum value in the sequence `a`
func MinN[T constraints.Ordered](a ...T) T {
	if len(a) == 0 {
		panic("can't call MaxN() with empty arguments list")
	}
	v := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < v {
			v = a[i]
		}
	}
	return v
}

func MinMax[T constraints.Ordered](a, b T) (min, max T) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinMaxN[T constraints.Ordered](a ...T) (min, max T) {
	if len(a) == 0 {
		panic("can't call MaxN() with empty arguments list")
	}
	min = a[0]
	max = a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}
	}
	return
}

func SumAs[R, T Numeric](a []T) R {
	var total R
	for v := range a {
		total += R(v)
	}
	return total
}

func Sum[T Numeric](a []T) T {
	return SumAs[T](a)
}

func AverageAs[R, T Numeric](a []T) R {
	return SumAs[R](a) / R(len(a))
}

func Average[T Numeric](a []T) T {
	return AverageAs[T](a)
}

//////////////////////////////////////////////////////////////////////////////
// Compare

// Equal returns whether two slices are equal.
// Return true if they are the same length and all elements equal.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Compare compares the elements of a and b.
func Compare[E constraints.Ordered](a, b []E) int {
	bl := len(b)
	for i, v1 := range a {
		if i >= bl {
			return 1
		}
		v2 := b[i]
		switch {
		case v1 < v2:
			return -1
		case v1 > v2:
			return 1
		}
	}
	if len(a) < bl {
		return -1
	}
	return 0
}

//////////////////////////////////////////////////////////////////////////////
// Lookup

// Find find the value x in the given slice a linearly
// return (index, true) if found
// return (_, false) if not found
func Find[T comparable](a []T, x T) (index int, ok bool) {
	for i, v := range a {
		if v == x {
			return i, true
		}
	}
	return -1, false
}

// Index find the value x in the given slice a linearly
// return index if found
// return -1 if not found
func Index[T comparable](a []T, x T) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

// AllOf return true if pred(e) returns true for all emements e in a
func AllOf[T any](a []T, pred func(T) bool) bool {
	for _, v := range a {
		if !pred(v) {
			return false
		}
	}
	return true
}

// AllOf return true if pred(e) returns true for any emements e in a
func AnyOf[T any](a []T, pred func(T) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}

// AllOf return true pred(e) returns true for none emements e in a
func NoneOf[T any](a []T, pred func(T) bool) bool {
	return !AnyOf(a, pred)
}

// Count returns the number of elements in the slice equals to x
func Count[T comparable](a []T, x T) int {
	c := 0
	for _, v := range a {
		if v == x {
			c++
		}
	}
	return c
}

// Count returns the number of elements in the slice which pred returns true
func CountIf[T comparable](a []T, pred func(T) bool) int {
	c := 0
	for _, v := range a {
		if pred(v) {
			c++
		}
	}
	return c
}

//////////////////////////////////////////////////////////////////////////////
// Sort

func Sort[T constraints.Ordered](a []T) {
}

func IsSorted[T constraints.Ordered](a []T) bool {
	if len(a) == 0 {
		return true
	}
	prev := a[0]
	for _, v := range a[1:] {
		if v < prev {
			return false
		}
		prev = v
	}
	return true
}

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
