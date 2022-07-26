package contalgo

func Transform[R any, T any](a []T, op func(T) R) []R {
	r := make([]R, 0, len(a))
	for _, v := range a {
		r = append(r, op(v))
	}
	return r
}

// Unique remove adjacent repeated elements from the input slice
// return the processed slice, andthe origin slice is also changed.
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

// Unique remove adjacent repeated elements from the input slice
// return the processed slice, andthe origin slice is kept unchanged.
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
