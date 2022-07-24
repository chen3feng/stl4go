package contalgo

import (
	"testing"
)

func Test_Queue_New(t *testing.T) {
	q := NewQueue[int]()
	expectTrue(t, q.IsEmpty())
	expectEq(t, 0, q.Len())
}

func Test_Queue_Clean(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.Clean()
	expectTrue(t, q.IsEmpty())
	expectEq(t, 0, q.Len())
}

func Test_Queue_String(t *testing.T) {
	q := NewQueue[int]()
	expectEq(t, "Queue[int]", q.String())
}

func Test_Queue_PushBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	expectFalse(t, q.IsEmpty())
	expectEq(t, 1, q.Len())
}

func Test_Queue_PopBack(t *testing.T) {
	q := NewQueue[int]()
	v, ok := q.PopBack()
	expectFalse(t, ok)

	q.PushBack(1)
	v, ok = q.PopFront()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}
