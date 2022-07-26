package contalgo

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
