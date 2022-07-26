package contalgo

import (
	"testing"
)

var (
	emptyInts = []int{}
)

func TestTransform(t *testing.T) {
	a := Range(0, 100)
	b := Transform(a, func(v int) int { return v * v })
	for i, v := range a {
		expectEq(t, v*v, b[i])
	}
}

func TestUnique(t *testing.T) {
	expectTrue(t, Equal(Unique(emptyInts), emptyInts))
	a := []int{1, 2, 2, 3, 2, 4}
	b := []int{1, 2, 3, 2, 4}
	expectTrue(t, Equal(Unique(a), b))
	expectTrue(t, Equal(Unique(a[:len(b)]), b))
}

func TestUniqueCopy(t *testing.T) {
	expectTrue(t, Equal(UniqueCopy(emptyInts), emptyInts))
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := append([]int{}, a...)
	b := []int{1, 2, 3, 2, 4}
	expectTrue(t, Equal(UniqueCopy(a), b))
	expectTrue(t, Equal(a, a1))
}
