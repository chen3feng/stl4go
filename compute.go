package contalgo

// Sum summarize all elements in a.
// returns the result as type R, this is useful when T is too small to hold the result.
// Complexity: O(len(a)).
func SumAs[R, T Numeric](a []T) R {
	var total R
	for v := range a {
		total += R(v)
	}
	return total
}

// Sum summarize all elements in a.
// returns the result as type R, you should use SumAs if T can't hold the result.
// Complexity: O(len(a)).
func Sum[T Numeric](a []T) T {
	return SumAs[T](a)
}

func AverageAs[R, T Numeric](a []T) R {
	return SumAs[R](a) / R(len(a))
}

func Average[T Numeric](a []T) T {
	return AverageAs[T](a)
}

// Count returns the number of elements in the slice equals to x.
// Complexity: O(len(a)).
func Count[T comparable](a []T, x T) int {
	c := 0
	for _, v := range a {
		if v == x {
			c++
		}
	}
	return c
}

// Count returns the number of elements in the slice which pred returns true.
// Complexity: O(len(a)).
func CountIf[T comparable](a []T, pred func(T) bool) int {
	c := 0
	for _, v := range a {
		if pred(v) {
			c++
		}
	}
	return c
}
