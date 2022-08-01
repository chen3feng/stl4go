package stl4go

import "testing"

func Test_SumAs(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, SumAs[int](a), 5050)
}

func Test_Sum(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Sum(a), 5050)
}

func Test_Average(t *testing.T) {
	a := Range(0, 101)
	expectEq(t, Average(a), 50)
}

func Test_Average_U64(t *testing.T) {
	a := Range[uint64](0, 101)
	expectEq(t, Average(a), 50)
}

func Test_Average_Float(t *testing.T) {
	a := Range(0.0, 101.0)
	expectEq(t, Average(a), 50.0)
}

func Test_Average_SmallType(t *testing.T) {
	a := Range[uint8](0, 101)
	expectEq(t, Average(a), 50)
}

func Test_Average_UintPtr(t *testing.T) {
	a := Range[uintptr](0, 101)
	expectEq(t, Average(a), 50)
}

func Test_Count(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	expectEq(t, Count(a, 3), 2)
	expectEq(t, CountIf(pos, isNegative), 0)
	expectEq(t, CountIf(neg, isNegative), 5)
	expectEq(t, CountIf(mix, isNegative), 2)
}
