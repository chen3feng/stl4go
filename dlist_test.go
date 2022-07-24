package contalgo

import (
	"testing"
)

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

func Test_DList_PushBack(t *testing.T) {
	l := NewDList[int]()
	l.PushBack(1)
	expectFalse(t, l.IsEmpty())
	expectEq(t, 1, l.Len())
}

func Test_DList_PopBack(t *testing.T) {
	l := NewDList[int]()
	v, ok := l.PopBack()
	expectFalse(t, ok)

	l.PushBack(1)
	v, ok = l.PopFront()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_DList_ForEach(t *testing.T) {
	l := NewDListOf(1, 2, 3)
	c := 0
	total := 0
	l.ForEach(func(n int) {
		c++
		total += n
	})
	expectEq(t, c, 3)
	expectEq(t, total, 6)
}
