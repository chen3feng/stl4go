package stl4go

import (
	"testing"
)

func Test_Queue_Interface(t *testing.T) {
	_ = Container[int](NewQueue[int]())
}

func Test_Queue_New(t *testing.T) {
	q := NewQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, 0, q.Len())
}

func Test_Queue_Clear(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.Clear()
	expectTrue(t, q.IsEmpty())
	expectEq(t, 0, q.Len())
}

func Test_Queue_String(t *testing.T) {
	q := NewQueue[int]()
	expectEq(t, "Queue[int]", q.String())
}

func Test_Queue_PushFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, 1, q.Len())
}

func Test_Queue_PushBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, 1, q.Len())
}

func Test_Queue_PopFront(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.PopFront()
	expectFalse(t, ok)
}
func Test_Queue_PopBack(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.PopBack()
	expectFalse(t, ok)
}

func Test_Queue_PushFront_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	v, ok := q.PopFront()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_Queue_PushFront_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	v, ok := q.PopBack()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_Queue_PushBack_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	v, ok := q.PopFront()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_Queue_PushBack_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	v, ok := q.PopBack()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}
