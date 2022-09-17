package stl4go

import (
	"testing"
)

func TestQueue_Interface(t *testing.T) {
	q := NewDListQueue[int]()
	_ = Deque[int](q)
}

func TestQueue_New(t *testing.T) {
	q := NewDListQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func TestQueue_Clear(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	q.Clear()
	expectTrue(t, q.IsEmpty())
	expectEq(t, q.Len(), 0)
}

func TestQueue_String(t *testing.T) {
	q := NewDListQueue[int]()
	expectEq(t, q.String(), "Queue[int]")
}

func TestQueue_Front_Back(t *testing.T) {
	q := NewDListQueue[int]()
	expectPanic(t, func() { q.Front() })
	expectPanic(t, func() { q.Back() })
	q.PushBack(1)
	q.PushBack(2)
	expectEq(t, q.Front(), 1)
	expectEq(t, q.Back(), 2)
}

func TestQueue_PushFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func TestQueue_PushBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, q.Len(), 1)
}

func TestQueue_TryPopFront(t *testing.T) {
	q := NewDListQueue[int]()
	_, ok := q.TryPopFront()
	expectFalse(t, ok)
}
func TestQueue_TryPopBack(t *testing.T) {
	q := NewDListQueue[int]()
	_, ok := q.TryPopBack()
	expectFalse(t, ok)
}

func TestQueue_PushFront_PopFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	q.PushFront(2)
	expectEq(t, q.PopFront(), 2)
	expectEq(t, q.PopFront(), 1)
}

func TestQueue_PushFront_PopBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushFront(1)
	expectEq(t, q.PopBack(), 1)
}

func TestQueue_PushBack_PopFront(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopFront(), 1)
}

func TestQueue_PushBack_PopBack(t *testing.T) {
	q := NewDListQueue[int]()
	q.PushBack(1)
	expectEq(t, q.PopBack(), 1)
}
