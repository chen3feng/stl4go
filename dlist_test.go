package stl4go

import (
	"testing"
)

func Test_Dlist_Interface(t *testing.T) {
	_ = Container(NewDList[int]())
}

func Test_DList_New(t *testing.T) {
	l := NewDList[int]()
	expectTrue(t, l.IsEmpty())
	expectEq(t, l.Len(), 0)
}

func Test_DList_NewOf(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 3)
}

func Test_DList_String(t *testing.T) {
	l := NewDList[int]()
	expectEq(t, l.String(), "DList[int]")
}

func Test_DList_Iterate(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	i := 1
	for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), i)
		i++
	}
	expectEq(t, i, 4)
}

func Test_DList_PushFront(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 1)
}

func Test_DList_PushBack(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, l.Len(), 1)
}

func Test_DList_PopFront(t *testing.T) {
	l := NewDListOf(1, 2)
	expectEq(t, l.PopFront(), 1)
	n, ok := l.TryPopFront()
	expectEq(t, n, 2)
	expectTrue(t, ok)
	n, ok = l.TryPopFront()
	expectFalse(t, ok)
	expectPanic(t, func() { l.PopFront() })
}

func Test_DList_PopBack(t *testing.T) {
	l := NewDListOf(1, 2)
	expectEq(t, l.PopBack(), 2)
	n, ok := l.TryPopBack()
	expectTrue(t, ok)
	expectEq(t, n, 1)
	n, ok = l.TryPopBack()
	expectFalse(t, ok)
	expectPanic(t, func() { l.PopBack() })
}

func Test_DList_PushBack_PopFront(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	v := l.PopFront()
	expectEq(t, v, 1)
}

func Test_DList_PushBack_PopBack(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	v := l.PopBack()
	expectEq(t, v, 1)
}

func Test_DList_PushFront_PopBack(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	v := l.PopBack()
	expectEq(t, v, 1)
}

func Test_DList_PushFront_PopFront(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	v := l.PopFront()
	expectEq(t, v, 1)
}

func Test_DList_ForEach(t *testing.T) {
	a := []int{1, 2, 3}
	l := NewDListOf(a...)
	var b []int
	l.ForEach(func(n int) {
		b = append(b, n)
	})
	expectEq(t, len(b), 3)
	expectTrue(t, Equal(a, b))
}

func Test_DList_ForEachIf(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	c := 0
	l.ForEachIf(func(n int) bool {
		c = n
		return n != 2
	})
	expectEq(t, c, 2)
}

func Test_DList_ForEachMutable(t *testing.T) {
	a := []int{1, 2, 3}
	l := NewDListOf(a...)
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

func Test_DList_ForEachMutableIf(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	c := 0
	l.ForEachMutableIf(func(n *int) bool {
		c = *n
		return *n != 2
	})
	expectEq(t, c, 2)
}

func Benchmark_DList_Iterate(b *testing.B) {
	l := NewDListOf(Range(1, 10000)...)
	b.Run("Iterate", func(b *testing.B) {
		sum := 0
		for i := 0; i < b.N; i++ {
			for it := l.Iterate(); it.IsNotEnd(); it.MoveToNext() {
				sum += it.Value()
			}
		}
	})
	b.Run("Iterate", func(b *testing.B) {
		sum := 0
		for i := 0; i < b.N; i++ {
			l.ForEach(func(val int) { sum += val })
		}
	})
}
