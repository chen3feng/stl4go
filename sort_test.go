package contalgo

import (
	"testing"
)

func TestIsSorted(t *testing.T) {
	expectTrue(t, IsSorted([]int{}))
	expectTrue(t, IsSorted([]int{1, 2, 3, 4}))
	expectTrue(t, IsSorted([]int{1, 2, 2, 3, 4}))
	expectFalse(t, IsSorted([]int{1, 2, 2, 3, 2}))
}
