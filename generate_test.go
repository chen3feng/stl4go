package stl4go

import "testing"

func Test_Range(t *testing.T) {
	a := Range(0, 100)
	expectEq(t, len(a), 100)
	expectEq(t, a[0], 0)
	expectEq(t, a[99], 99)
}

func Test_Generate(t *testing.T) {
	a := make([]int, 100)
	i := -1
	Generate(a, func() int { i++; return i })
	for i, v := range a {
		expectEq(t, v, i)
	}
}
