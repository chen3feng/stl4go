package stl4go

// Max return the larger value between `a` and `b`.
//
// Complexity: O(1).
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Min return the smaller value between `a` and `b`.
//
// Complexity: O(1).
func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MaxN return the maximum value in the sequence `a`.
//
// Complexity: O(len(a)).
func MaxN[T Ordered](a ...T) T {
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

// MinN return the minimum value in the sequence `a`.
//
// Complexity: O(len(a)).
func MinN[T Ordered](a ...T) T {
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

// MinMax returns both min and max between a and b.
//
// Complexity: O(1).
func MinMax[T Ordered](a, b T) (min, max T) {
	if a < b {
		return a, b
	}
	return b, a
}

// MinMaxN returns both min and max in slice a.
//
// Complexity: O(len(a))
func MinMaxN[T Ordered](a ...T) (min, max T) {
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

// Find find the value x in the given slice a linearly.
// return (index, true) if found,
// return (_, false) if not found.
// Complexity: O(len(a)).
func Find[T comparable](a []T, x T) (index int, ok bool) {
	for i, v := range a {
		if v == x {
			return i, true
		}
	}
	return -1, false
}

// Index find the value x in the given slice a linearly.
//
// Return index if found, -1 if not found.
//
// Complexity: O(len(a)).
func Index[T comparable](a []T, x T) int {
	for i, v := range a {
		if v == x {
			return i
		}
	}
	return -1
}

// AllOf return true if pred(e) returns true for all emements e in a.
//
// Complexity: O(len(a)).
func AllOf[T any](a []T, pred func(T) bool) bool {
	for _, v := range a {
		if !pred(v) {
			return false
		}
	}
	return true
}

// AnyOf return true if pred(e) returns true for any emements e in a.
//
// Complexity: O(len(a)).
func AnyOf[T any](a []T, pred func(T) bool) bool {
	for _, v := range a {
		if pred(v) {
			return true
		}
	}
	return false
}

// NoneOf return true pred(e) returns true for none emements e in a.
//
// Complexity: O(len(a)).
func NoneOf[T any](a []T, pred func(T) bool) bool {
	return !AnyOf(a, pred)
}
