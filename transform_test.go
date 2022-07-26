package contalgo

import (
	"strconv"
	"testing"
)

var (
	emptyInts = []int{}
)

func TestCopy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := Copy(a)
	expectTrue(t, Equal(a, b))
}

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

func TestRemove(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	b := Remove(a, 2)
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal([]int{1, 3, 4}, a[:len(b)]))
}

func TestRemoveCopy(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := []int{1, 2, 2, 3, 2, 4}
	b := RemoveCopy(a, 2)
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal(a1, a))
}

func TestRemoveIf(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	b := RemoveIf(a, func(v int) bool { return v == 2 })
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal([]int{1, 3, 4}, a[:len(b)]))
}

func TestRemoveIfCopy(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := []int{1, 2, 2, 3, 2, 4}
	b := RemoveIfCopy(a, func(v int) bool { return v == 2 })
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal(a1, a))
}
