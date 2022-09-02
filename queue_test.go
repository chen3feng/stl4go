package stl4go

import (
	"testing"
)

func Test_Queue_Interface(t *testing.T) {
	_ = Container(NewQueue[int]())
}

func Test_Queue_New(t *testing.T) {
	q := NewQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_Clear(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.Clear()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_String(t *testing.T) {
	q := NewQueue[int]()
	expectEq(t, q.String(), "Queue[int]")
}

func Test_Queue_PushFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_PushBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_TryPopFront(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.TryPopFront()
	expectFalse(t, ok)
}
func Test_Queue_TryPopBack(t *testing.T) {
	q := NewQueue[int]()
	_, ok := q.TryPopBack()
	expectFalse(t, ok)
}

func Test_Queue_PushFront_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	q.PushFront(2)
	expectEq(t, q.PopFront(), 2)
	expectEq(t, q.PopFront(), 1)
}

func Test_Queue_PushFront_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)
	expectEq(t, q.PopBack(), 1)
}

func Test_Queue_PushBack_PopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopFront(), 1)
}

func Test_Queue_PushBack_PopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopBack(), 1)
}
