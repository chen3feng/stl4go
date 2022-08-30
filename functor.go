package stl4go

// Less wraps the '<' operator for ordered types.
func Less[T Ordered](a, b T) bool {
	return a < b
}

// Greater wraps the '>' operator for ordered types.
func Greater[T Ordered](a, b T) bool {
	return a > b
}

// OrderedCompare provide default CompareFn for ordered types.
func OrderedCompare[T Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
