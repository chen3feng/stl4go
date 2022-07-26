package contalgo

import (
	"strconv"
	"testing"
)

var (
	emptyInts = []int{}
)

func TestTransform(t *testing.T) {
	a := Range(0, 100)
	Transform(a, func(v int) int { return v * v })
	for i, v := range a {
		expectEq(t, v, i*i)
	}
}

func TestTransformTo(t *testing.T) {
	a := Range(0, 100)
	b := make([]string, len(a))
	TransformTo(a, func(v int) string { return strconv.Itoa(v) }, b)
	for i, v := range a {
		expectEq(t, strconv.Itoa(v), b[i])
	}
	expactPanic(t, func() {
		c := make([]string, len(a)-1)
		TransformTo(a, func(v int) string { return strconv.Itoa(v) }, c)
	})
}

func TestTransformCopy(t *testing.T) {
	a := Range(0, 100)
	b := TransformCopy(a, func(v int) string { return strconv.Itoa(v) })
	for i, v := range b {
		expectEq(t, v, strconv.Itoa(i))
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
