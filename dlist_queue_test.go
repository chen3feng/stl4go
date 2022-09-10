package stl4go

import (
	"testing"
)

func Test_Queue_Interface(t *testing.T) {
	q := NewDListQueue[int]()
	_ = Deque[int](q)
}

func Test_Queue_New(t *testing.T) {
	q := NewDListQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_Clear(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	q.Clear()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func Test_Queue_String(t *testing.T) {
	q := NewDListQueue[int]()
	expectEq(t, q.String(), "Queue[int]")
}

func Test_Queue_Front_Back(t *testing.T) {
	q := NewDListQueue[int]()
	expectPanic(t, func() { q.Front() })
	expectPanic(t, func() { q.Back() })
	q.PushBack(1)
	q.PushBack(2)
	expectEq(t, q.Front(), 1)
	expectEq(t, q.Back(), 2)
}

func Test_Queue_PushFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_PushBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func Test_Queue_TryPopFront(t *testing.T) {
	q := NewDListQueue[int]()
	_, ok := q.TryPopFront()
	expectFalse(t, ok)
}
func Test_Queue_TryPopBack(t *testing.T) {
	q := NewDListQueue[int]()
	_, ok := q.TryPopBack()
	expectFalse(t, ok)
}

func Test_Queue_PushFront_PopFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	q.PushFront(2)
	expectEq(t, q.PopFront(), 2)
	expectEq(t, q.PopFront(), 1)
}

func Test_Queue_PushFront_PopBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	expectEq(t, q.PopBack(), 1)
}

func Test_Queue_PushBack_PopFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopFront(), 1)
}

func Test_Queue_PushBack_PopBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopBack(), 1)
}
