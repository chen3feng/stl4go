package contalgo

// Equal returns whether two slices are equal.
// Return true if they are the same length and all elements are equal.
//
// Complexity: O(min(len(a), len(b))).
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

// Compare compares each elements in a and b.
//
// return 0 if they are equals,
// return 1 if a > b,
// return -1 if a < b.
//
// Complexity: O(min(len(a), len(b))).
func Compare[E Ordered](a, b []E) int {
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
