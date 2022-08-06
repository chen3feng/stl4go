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
	expectEq(t, 0, l.Len())
}

func Test_DList_NewOf(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	expectFalse(t, l.IsEmpty())
	expectEq(t, 3, l.Len())
}

func Test_DList_String(t *testing.T) {
	l := NewDList[int]()
	expectEq(t, "DList[int]", l.String())
}

func Test_DList_PushFront(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, 1, l.Len())
}

func Test_DList_PushBack(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, 1, l.Len())
}

func Test_DList_PopFront(t *testing.T) {
	l := NewDList[int]()
	_, ok := l.PopFront()
	expectFalse(t, ok)
}

func Test_DList_PopBack(t *testing.T) {
	l := NewDList[int]()
	_, ok := l.PopBack()
	expectFalse(t, ok)
}

func Test_DList_PushBack_PopFront(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	v, ok := l.PopFront()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_DList_PushBack_PopBack(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	v, ok := l.PopBack()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_DList_PushFront_PopBack(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	v, ok := l.PopBack()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_DList_PushFront_PopFront(t *testing.T) {
	l := NewDList[int]()
	l.PushFront(1)
	v, ok := l.PopFront()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_DList_ForEach(t *testing.T) {
	a := []int{1, 2, 3}
	l := NewDListOf(a...)
	c := 0
	var b []int
	l.ForEach(func(n int) {
		c++
		b = append(b, n)
	})
	expectEq(t, c, 3)
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
