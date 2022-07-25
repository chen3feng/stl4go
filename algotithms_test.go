package contalgo

import (
	"testing"
)

//////////////////////////////////////////////////////////////////////////////
// Generative functions

func TestGenerate(t *testing.T) {
	a := make([]int, 100)
	i := -1
	Generate(a, func() int { i++; return i })
	for i, v := range a {
		expectEq(t, v, i)
	}
}

func TestTransform(t *testing.T) {
	a := Range(0, 100)
	b := Transform(a, func(v int) int { return v * v })
	for i, v := range a {
		expectEq(t, v*v, b[i])
	}
}

//////////////////////////////////////////////////////////////////////////////
// Min, Max

func TestMin(t *testing.T) {
	expectEq(t, Min(1, 2), 1)
	expectEq(t, Min(2, 1), 1)
	expectEq(t, Min(1, 1), 1)
	expectEq(t, Min("hello", "world"), "hello")
}

func TestMax(t *testing.T) {
	expectEq(t, Max(1, 2), 2)
	expectEq(t, Max(2, 1), 2)
	expectEq(t, Max(2, 2), 2)
	expectEq(t, Max("hello", "world"), "world")
}

func TestMinN(t *testing.T) {
	expectEq(t, MinN(1, 2, 3), 1)
	expectEq(t, MinN(2, 1, 3), 1)
	expectEq(t, MinN(1, 1, 1), 1)
	expectEq(t, MinN("hello", "world"), "hello")
}

func TestMaxN(t *testing.T) {
	expectEq(t, MaxN(1, 2), 2)
	expectEq(t, MaxN(2, 1), 2)
	expectEq(t, MaxN(2, 2), 2)
	expectEq(t, MaxN("hello", "world"), "world")
}

func TestMinMax(t *testing.T) {
	min, max := MinMax(1, 2)
	expectEq(t, min, 1)
	expectEq(t, max, 2)
	min, max = MinMax(2, 1)
	expectEq(t, min, 1)
	expectEq(t, max, 2)
}

func TestMinMaxN(t *testing.T) {
	min, max := MinMaxN(4, 3, 2, 1)
	expectEq(t, min, 1)
	expectEq(t, max, 4)
}

func TestSumAs(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, SumAs[int](a), 5050)
}

func TestAverageAs(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, AverageAs[int](a), 50)
}

func TestSum(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Sum(a), 5050)
}

func TestAverage(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Average(a), 50)
}

func TestEqual(t *testing.T) {
	expectTrue(t, Equal([]int{}, []int{}))
	expectTrue(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	expectFalse(t, Equal([]int{1, 2}, []int{1, 2, 3}))
	expectFalse(t, Equal([]int{1, 2, 2}, []int{1, 2, 3}))
}

func TestCompare(t *testing.T) {
	expectEq(t, Compare([]int{}, []int{}), 0)
	expectEq(t, Compare([]int{1, 2, 3}, []int{1, 2, 3}), 0)
	expectEq(t, Compare([]int{1, 2}, []int{1, 2, 3}), -1)
	expectEq(t, Compare([]int{1, 2, 3}, []int{1, 2}), 1)
	expectEq(t, Compare([]int{1, 2, 4}, []int{1, 2, 3}), 1)
}

func TestFind(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	i, ok := Find(a, 3)
	expectTrue(t, ok)
	expectEq(t, i, 2)
	i, ok = Find(a, 5)
	expectFalse(t, ok)
}

func TestIndex(t *testing.T) {
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

func TestAllOf(t *testing.T) {
	expectFalse(t, AllOf(pos, isNegative))
	expectTrue(t, AllOf(neg, isNegative))
	expectFalse(t, AllOf(mix, isNegative))
}

func TestAnyOf(t *testing.T) {
	expectFalse(t, AnyOf(pos, isNegative))
	expectTrue(t, AnyOf(neg, isNegative))
	expectTrue(t, AnyOf(mix, isNegative))
}

func TestNoneOf(t *testing.T) {
	expectTrue(t, NoneOf(pos, isNegative))
	expectFalse(t, NoneOf(neg, isNegative))
	expectFalse(t, NoneOf(mix, isNegative))
}

func TestCount(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	expectEq(t, Count(a, 3), 2)
	expectEq(t, CountIf(pos, isNegative), 0)
	expectEq(t, CountIf(neg, isNegative), 5)
	expectEq(t, CountIf(mix, isNegative), 2)
}

func TestIsSorted(t *testing.T) {
	expectTrue(t, IsSorted([]int{1, 2, 3, 4}))
	expectTrue(t, IsSorted([]int{1, 2, 2, 3, 4}))
	expectFalse(t, IsSorted([]int{1, 2, 2, 3, 2}))
}

func TestUnique(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	b := []int{1, 2, 3, 2, 4}
	expectTrue(t, Equal(Unique(a), b))
}
