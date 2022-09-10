package stl4go

import (
	"strconv"
	"testing"
)

var (
	emptyInts = []int{}
)

func Test_Copy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := Copy(a)
	expectTrue(t, Equal(a, b))
}

func Test_Fill(t *testing.T) {
	a := []int{1, 2, 3, 4}
	Fill(a, 1)
	expectTrue(t, Equal(a, []int{1, 1, 1, 1}))
	b := a[:0]
	Fill(b, 2)
	expectTrue(t, Equal(a, []int{1, 1, 1, 1}))
}

func naiveFill[T any](a []T, v T) {
	for i := range a {
		a[i] = v
	}
}

func BenchmarkFill(b *testing.B) {
	a := make([]int, 10000)
	b.Run("Fill", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Fill(a, 1)
		}
	})
	b.Run("naiveFill", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			naiveFill(a, 1)
		}
	})
}

func Test_FillPattern(t *testing.T) {
	a := make([]int, 10)
	p := []int{1, 2, 3}
	FillPattern(a, p)
	expectTrue(t, Equal(a, []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}))
	expectPanic(t, func() { FillPattern(a, []int{}) })
}

func Test_Transform(t *testing.T) {
	a := Range(0, 100)
	Transform(a, func(v int) int { return v * v })
	for i, v := range a {
		expectEq(t, v, i*i)
	}
}

func Test_TransformTo(t *testing.T) {
	a := Range(0, 100)
	b := make([]string, len(a))
	TransformTo(a, func(v int) string { return strconv.Itoa(v) }, b)
	for i, v := range a {
		expectEq(t, b[i], strconv.Itoa(v))
	}
	expectPanic(t, func() {
		c := make([]string, len(a)-1)
		TransformTo(a, func(v int) string { return strconv.Itoa(v) }, c)
	})
}

func Test_TransformCopy(t *testing.T) {
	a := Range(0, 100)
	b := TransformCopy(a, func(v int) string { return strconv.Itoa(v) })
	for i, v := range b {
		expectEq(t, v, strconv.Itoa(i))
	}
}

func Test_Replace(t *testing.T) {
	a := []int{1, 2, 2, 4}
	Replace(a, 2, 3)
	expectTrue(t, Equal(a, []int{1, 3, 3, 4}))
}

func Test_ReplaceIf(t *testing.T) {
	a := []int{1, 2, 2, 4}
	ReplaceIf(a, func(n int) bool { return n == 2 }, 3)
	expectTrue(t, Equal(a, []int{1, 3, 3, 4}))
}

func Test_Unique(t *testing.T) {
	expectTrue(t, Equal(Unique(emptyInts), emptyInts))
	a := []int{1, 2, 2, 3, 2, 4}
	b := []int{1, 2, 3, 2, 4}
	expectTrue(t, Equal(Unique(a), b))
	expectTrue(t, Equal(Unique(a[:len(b)]), b))
}

func Test_UniqueCopy(t *testing.T) {
	expectTrue(t, Equal(UniqueCopy(emptyInts), emptyInts))
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := append([]int{}, a...)
	b := []int{1, 2, 3, 2, 4}
	expectTrue(t, Equal(UniqueCopy(a), b))
	expectTrue(t, Equal(a, a1))
}

func Test_Remove(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	b := Remove(a, 2)
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal([]int{1, 3, 4}, a[:len(b)]))
}

func Test_RemoveCopy(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := []int{1, 2, 2, 3, 2, 4}
	b := RemoveCopy(a, 2)
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal(a1, a))
}

func Test_RemoveIf(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	b := RemoveIf(a, func(v int) bool { return v == 2 })
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal([]int{1, 3, 4}, a[:len(b)]))
}

func Test_RemoveIfCopy(t *testing.T) {
	a := []int{1, 2, 2, 3, 2, 4}
	a1 := []int{1, 2, 2, 3, 2, 4}
	b := RemoveIfCopy(a, func(v int) bool { return v == 2 })
	expectTrue(t, Equal([]int{1, 3, 4}, b))
	expectTrue(t, Equal(a1, a))
}

func Test_Shuffle(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	Shuffle(a)
}

func Test_Reverse(t *testing.T) {
	a := []int{1, 2, 3, 4}
	Reverse(a)
	expectTrue(t, Equal(a, []int{4, 3, 2, 1}))
}

func Test_Reverse_Odd(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	Reverse(a)
	expectTrue(t, Equal(a, []int{5, 4, 3, 2, 1}))
}

func Test_ReverseCopy(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := ReverseCopy(a)
	expectTrue(t, Equal(b, []int{5, 4, 3, 2, 1}))
}
