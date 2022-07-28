package contalgo

import "testing"

func Test_Equal(t *testing.T) {
	expectTrue(t, Equal([]int{}, []int{}))
	expectTrue(t, Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	expectFalse(t, Equal([]int{1, 2}, []int{1, 2, 3}))
	expectFalse(t, Equal([]int{1, 2, 2}, []int{1, 2, 3}))
}

func Test_Compare(t *testing.T) {
	expectEq(t, Compare([]int{}, []int{}), 0)
	expectEq(t, Compare([]int{1, 2, 3}, []int{1, 2, 3}), 0)
	expectEq(t, Compare([]int{1, 2, 2}, []int{1, 2, 3}), -1)
	expectEq(t, Compare([]int{1, 2, 4}, []int{1, 2, 3}), 1)
	expectEq(t, Compare([]int{1, 2}, []int{1, 2, 3}), -1)
	expectEq(t, Compare([]int{1, 2, 3}, []int{1, 2}), 1)
}
