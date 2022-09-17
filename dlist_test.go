package stl4go

import (
	"testing"
)

func TestDlist_Interface(t *testing.T) {
	_ = Container(&DList[int]{})
}

func TestDList_New(t *testing.T) {
	l := DList[int]{}
	expectTrue(t, l.IsEmpty())
	expectEq(t, l.Len(), 0)
}

func TestDListOf(t *testing.T) {
	l := DListOf(1, 2, 3)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 3)
}

func TestDList_String(t *testing.T) {
	l := DList[int]{}
	expectEq(t, l.String(), "DList[int]")
}

func TestDList_Iterate(t *testing.T) {
	l := DListOf(1, 2, 3)
	i := 1
	for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), i)
		i++
	}
	expectEq(t, i, 4)
}

func TestDList_Iterate_Empty(t *testing.T) {
	l := DList[int]{}
	i := 0
	for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		i++
	}
	expectEq(t, i, 0)
}

func TestDList_FrontBack(t *testing.T) {
	l := DListOf(1, 2, 3)
	expectEq(t, l.Front(), 1)
	expectEq(t, l.Back(), 3)
}

func TestDList_PushFront(t *testing.T) {
	l := DList[int]{}
	l.PushFront(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 1)
}

func TestDList_PushBack(t *testing.T) {
	l := DList[int]{}
	l.PushBack(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 1)
}

func TestDList_PopFront(t *testing.T) {
	l := DListOf(1, 2)
	expectEq(t, l.PopFront(), 1)
	n, ok := l.TryPopFront()
	expectEq(t, n, 2)
	expectTrue(t, ok)
	n, ok = l.TryPopFront()
	expectFalse(t, ok)
	expectPanic(t, func() { l.PopFront() })
}

func TestDList_PopBack(t *testing.T) {
	l := DListOf(1, 2)
	expectEq(t, l.PopBack(), 2)
	n, ok := l.TryPopBack()
	expectTrue(t, ok)
	expectEq(t, n, 1)
	n, ok = l.TryPopBack()
	expectFalse(t, ok)
	expectPanic(t, func() { l.PopBack() })
}

func TestDList_PushBack_PopFront(t *testing.T) {
	l := DList[int]{}
	l.PushBack(1)
	v := l.PopFront()
	expectEq(t, v, 1)
}

func TestDList_PushBack_PopBack(t *testing.T) {
	l := DList[int]{}
	l.PushBack(1)
	v := l.PopBack()
	expectEq(t, v, 1)
}

func TestDList_PushFront_PopBack(t *testing.T) {
	l := DList[int]{}
	l.PushFront(1)
	v := l.PopBack()
	expectEq(t, v, 1)
}

func TestDList_PushFront_PopFront(t *testing.T) {
	l := DList[int]{}
	l.PushFront(1)
	v := l.PopFront()
	expectEq(t, v, 1)
}

func TestDList_ForEach(t *testing.T) {
	a := []int{1, 2, 3}
	l := DListOf(a...)
	var b []int
	l.ForEach(func(n int) {
		b = append(b, n)
	})
	expectEq(t, len(b), 3)
	expectTrue(t, Equal(a, b))
}

func TestDList_ForEachIf(t *testing.T) {
	l := DListOf(1, 2, 3)
	c := 0
	l.ForEachIf(func(n int) bool {
		c = n
		return n != 2
	})
	expectEq(t, c, 2)
}

func TestDList_ForEachMutable(t *testing.T) {
	a := []int{1, 2, 3}
	l := DListOf(a...)
	l.ForEachMutable(func(n *int) {
		*n = -*n
	})
	var b []int
	l.ForEach(func(n int) {
		b = append(b, n)
	})
	expectEq(t, len(b), 3)
	for i := range b {
		expectEq(t, a[i], -b[i])
	}
}

func TestDList_ForEachMutableIf(t *testing.T) {
	l := DListOf(1, 2, 3)
	c := 0
	l.ForEachMutableIf(func(n *int) bool {
		c = *n
		return *n != 2
	})
	expectEq(t, c, 2)
}

func TestDList_ForEach_EmptyOK(t *testing.T) {
	l := DList[int]{}
	l.ForEach(func(n int) {})
	l.ForEachIf(func(n int) bool { return true })
	l.ForEachMutable(func(n *int) {})
	l.ForEachMutableIf(func(n *int) bool { return true })
}

func Benchmark_DList_Iterate(b *testing.B) {
	l := DListOf(Range(1, 10000)...)
	b.Run("Iterator", func(b *testing.B) {
		sum := 0
		for i := 0; i < b.N; i++ {
			for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
				sum += it.Value()
			}
		}
	})
	b.Run("ForEach", func(b *testing.B) {
		sum := 0
		for i := 0; i < b.N; i++ {
			l.ForEach(func(val int) { sum += val })
		}
	})
}
