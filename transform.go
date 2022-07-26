package contalgo

// Transform applies the function op to each element in slice a and fill it to slice b.
// The len(b) must not lesser than len(a).
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
// Complexity: O(len(a)).
func Transform[T any](a []T, op func(T) T) {
	for i, v := range a {
		a[i] = op(v)
	}
}

// Transform applies the function op to each element in slice a and return all the result as a slice.
// Complexity: O(len(a)).
func TransformCopy[R any, T any](a []T, op func(T) R) []R {
	r := make([]R, 0, len(a))
	for _, v := range a {
		r = append(r, op(v))
	}
	return r
}

// Unique remove adjacent repeated elements from the input slice.
// return the processed slice, andthe origin slice is also changed.
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

// Unique remove adjacent repeated elements from the input slice.
// return the processed slice, andthe origin slice is kept unchanged.
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
