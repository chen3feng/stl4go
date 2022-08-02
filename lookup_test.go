package stl4go

import "testing"

func Test_Min(t *testing.T) {
	expectEq(t, Min(1, 2), 1)
	expectEq(t, Min(2, 1), 1)
	expectEq(t, Min(1, 1), 1)
	expectEq(t, Min("hello", "world"), "hello")
}

func Test_Max(t *testing.T) {
	expectEq(t, Max(1, 2), 2)
	expectEq(t, Max(2, 1), 2)
	expectEq(t, Max(2, 2), 2)
	expectEq(t, Max("hello", "world"), "world")
}

func Test_MinN(t *testing.T) {
	expectEq(t, MinN(1, 2, 3), 1)
	expectEq(t, MinN(2, 1, 3), 1)
	expectEq(t, MinN(1, 1, 1), 1)
	expectEq(t, MinN("hello", "world"), "hello")
	expactPanic(t, func() { MinN(emptyInts...) })
}

func Test_MaxN(t *testing.T) {
	expectEq(t, MaxN(1, 2), 2)
	expectEq(t, MaxN(2, 1), 2)
	expectEq(t, MaxN(2, 2), 2)
	expectEq(t, MaxN("hello", "world"), "world")
	expactPanic(t, func() { MaxN(emptyInts...) })
}

func Test_MinMax(t *testing.T) {
	min, max := MinMax(1, 2)
	expectEq(t, min, 1)
	expectEq(t, max, 2)
	min, max = MinMax(2, 1)
	expectEq(t, min, 1)
	expectEq(t, max, 2)
}

func Test_MinMaxN(t *testing.T) {
	min, max := MinMaxN(3, 4, 1, 2)
	expectEq(t, min, 1)
	expectEq(t, max, 4)
	expactPanic(t, func() { MinMaxN(emptyInts...) })
}

func Test_Find(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	i, ok := Find(a, 3)
	expectTrue(t, ok)
	expectEq(t, i, 2)
	i, ok = Find(a, 5)
	expectFalse(t, ok)
}

func Test_FindIf(t *testing.T) {
	isNeg := func(x int) bool { return x < 0 }
	a := []int{1, 2, -3, 4, 3}
	i, ok := FindIf(a, isNeg)
	expectTrue(t, ok)
	expectEq(t, i, 2)
	i, ok = FindIf([]int{1, 2, 3, 4, 3}, isNeg)
	expectFalse(t, ok)
}

func Test_Index(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	expectEq(t, Index(a, 3), 2)
	expectEq(t, Index(a, 5), -1)
}

var (
	pos = []int{1, 2, 3, 4, 5}
	neg = []int{-1, -2, -3, -4, -5}
	mix = []int{1, -2, 3, -4, 5}
)

func isNegative(n int) bool { return n < 0 }

func Test_AllOf(t *testing.T) {
	expectFalse(t, AllOf(pos, isNegative))
	expectTrue(t, AllOf(neg, isNegative))
	expectFalse(t, AllOf(mix, isNegative))
}

func Test_AnyOf(t *testing.T) {
	expectFalse(t, AnyOf(pos, isNegative))
	expectTrue(t, AnyOf(neg, isNegative))
	expectTrue(t, AnyOf(mix, isNegative))
}

func Test_NoneOf(t *testing.T) {
	expectTrue(t, NoneOf(pos, isNegative))
	expectFalse(t, NoneOf(neg, isNegative))
	expectFalse(t, NoneOf(mix, isNegative))
}
