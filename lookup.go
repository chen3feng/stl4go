package contalgo

import "golang.org/x/exp/constraints"

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
