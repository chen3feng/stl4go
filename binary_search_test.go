package stl4go

import "testing"

func TestLowerBound(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	expectEq(t, LowerBound(a, 1), 0)
	expectEq(t, LowerBound(a, 5), 3)
	expectEq(t, LowerBound(a, 7), len(a))
}

func TestLowerBoundFunc(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	expectEq(t, LowerBoundFunc(a, 1, Less[int]), 0)
	expectEq(t, LowerBoundFunc(a, 5, Less[int]), 3)
	expectEq(t, LowerBoundFunc(a, 7, Less[int]), len(a))
}

func TestUpperBound(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	expectEq(t, UpperBound(a, 1), 1)
	expectEq(t, UpperBound(a, 5), 5)
	expectEq(t, UpperBound(a, 7), len(a))
}

func TestUpperBoundFunc(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}
	expectEq(t, UpperBoundFunc(a, 1, Less[int]), 1)
	expectEq(t, UpperBoundFunc(a, 5, Less[int]), 5)
	expectEq(t, UpperBoundFunc(a, 7, Less[int]), len(a))
}

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}

	i, ok := BinarySearch(a, 4)
	expectTrue(t, ok)
	expectEq(t, i, 2)

	i, ok = BinarySearch(a, 5)
	expectTrue(t, ok)
	expectEq(t, i, 3)

	i, ok = BinarySearch(a, 3)
	expectFalse(t, ok)
}

func TestBinarySearchFunc(t *testing.T) {
	a := []int{1, 2, 4, 5, 5, 6}

	i, ok := BinarySearchFunc(a, 4, Less[int])
	expectTrue(t, ok)
	expectEq(t, i, 2)

	i, ok = BinarySearchFunc(a, 5, Less[int])
	expectTrue(t, ok)
	expectEq(t, i, 3)

	i, ok = BinarySearchFunc(a, 3, Less[int])
	expectFalse(t, ok)
}
