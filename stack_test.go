package stl4go

import (
	"testing"
)

func Test_Stack_Interface(t *testing.T) {
	_ = Container(NewStack[int]())
}
func Test_NewStack(t *testing.T) {
	si := NewStack[int]()
	ss := NewStack[string]()
	expectEq(t, si.Len(), 0)
	expectEq(t, ss.Len(), 0)
}

func Test_NewStackCap(t *testing.T) {
	s := NewStackCap[int](10)
	expectEq(t, s.Len(), 0)
	expectEq(t, s.Cap(), 10)
}

func Test_StackCap(t *testing.T) {
	s := NewStackCap[int](10)
	s.Push(1)
	expectEq(t, s.Len(), 1)
	expectEq(t, s.Cap(), 10)
}

func Test_Stack_Clear(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Clear()
	expectEq(t, s.Len(), 0)
	expectTrue(t, s.IsEmpty())
}

func Test_Stack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	expectEq(t, s.Len(), 1)
	s.Push(2)
	expectEq(t, s.Len(), 2)
}

func Test_Stack_TryPop(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.TryPop()
	expectFalse(t, ok)
	s.Push(1)
	v, ok := s.TryPop()
	expectTrue(t, ok)
	expectEq(t, v, 1)
}

func Test_Stack_Pop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	v := s.Pop()
	expectEq(t, v, 1)
	expectPanic(t, func() { s.Pop() })
}

func Test_Stack_Top(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	v := s.Top()
	expectEq(t, v, 1)
	s.Pop()
	expectPanic(t, func() { s.Top() })
}
