package contalgo

import "golang.org/x/exp/constraints"

// IsSorted returns whether the slice a is sorted in ascending order.
// Complexity: O(len(a)).
func IsSorted[T constraints.Ordered](a []T) bool {
	if len(a) == 0 {
		return true
	}
	prev := a[0]
	for _, v := range a[1:] {
		if v < prev {
			return false
		}
		prev = v
	}
	return true
}
