package stl4go

import (
	"testing"
)

func Test_SumAs(t *testing.T) {
	t.Run("sum uint8 to int", func(t *testing.T) {
		a := Range[uint8](1, 101)
		expectEq(t, SumAs[int](a), 5050)
	})

	t.Run("sum int to uint8", func(t *testing.T) {
		a := Range[int](1, 101)
		expectEq(t, SumAs[uint8](a), uint8(5050%256))
	})

	t.Run("sum int64 to float64", func(t *testing.T) {
		a := Range[int64](1, 101)
		expectEq(t, SumAs[float64](a), 5050.)
	})

	t.Run("sum float64 to int64", func(t *testing.T) {
		a := Range[float64](1.1, 101.1)
		expectEq(t, SumAs[int](a), 101.2*50) // 5060
	})
}

func Test_Sum(t *testing.T) {
	a := Range(1, 101)
	expectEq(t, Sum(a), 5050)
}

func Test_Average(t *testing.T) {
	a := Range(1, 101)
	expectEq(t, Average(a), 50)
}

func Test_AverageAs(t *testing.T) {
	a := []int{1, 0}
	expectEq(t, AverageAs[float64](a), 0.5)
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

func Test_Average_Signed(t *testing.T) {
	a := []int{-2, 1, -1, 2, 1, -1, 0}
	expectEq(t, Average(a), 0)
}

func Test_Count(t *testing.T) {
	a := []int{1, 2, 3, 4, 3}
	expectEq(t, Count(a, 3), 2)
	expectEq(t, CountIf(pos, isNegative), 0)
	expectEq(t, CountIf(neg, isNegative), 5)
	expectEq(t, CountIf(mix, isNegative), 2)
}
