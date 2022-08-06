package stl4go

// SumAs summarize all elements in a.
// returns the result as type R, this is useful when T is too small to hold the result.
// Complexity: O(len(a)).
func SumAs[R, T Numeric](a []T) R {
	var total R
	for _, v := range a {
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

// AverageAs returns the average value of a as type R.
func AverageAs[R, T Numeric](a []T) R {
	return SumAs[R](a) / R(len(a))
}

// Average returns the average value of a.
func Average[T Numeric](a []T) T {
	var x T
	var i interface{} = x
	switch i.(type) {
	case int, int8, uint8, int16, uint16, int32, uint32:
		return T(AverageAs[int64](a))
	case uint64:
		return T(AverageAs[uint64](a))
	case float32, float64:
		return T(AverageAs[float64](a))
	}
	// int64, uint64, uintptr ...
	return AverageAs[T](a)
}

// Count returns the number of elements in the slice equals to x.
//
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

// CountIf returns the number of elements in the slice which pred returns true.
//
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
