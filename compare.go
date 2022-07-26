package contalgo

import "golang.org/x/exp/constraints"

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
