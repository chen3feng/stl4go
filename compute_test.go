package contalgo

import "testing"

func TestSumAs(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, SumAs[int](a), 5050)
}

func TestSum(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Sum(a), 5050)
}

func TestAverage(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Average(a), 50)
}

func TestAverage_Float(t *testing.T) {
	a := Range(0.0, 101.0)
	expectEq(t, Average(a), 50.0)
}

func TestAverage_SmallType(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, Average(a), 50)
}

func TestAverage_UintPtr(t *testing.T) {
	a := Range[uintptr](0, 101)
	expectEq(t, Average(a), 50)
}

func TestCount(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	expectEq(t, Count(a, 3), 2)
	expectEq(t, CountIf(pos, isNegative), 0)
	expectEq(t, CountIf(neg, isNegative), 5)
	expectEq(t, CountIf(mix, isNegative), 2)
}
