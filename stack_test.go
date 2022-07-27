package contalgo

import (
	"testing"
)

func Test_Stack_Interface(t *testing.T) {
	_ = Container[int](NewStack[int]())
}
func Test_NewStack(t *testing.T) {
	si := NewStack[int]()
	ss := NewStack[string]()
	expectEq(t, 0, si.Len())
	expectEq(t, 0, ss.Len())
}

func Test_NewStackCap(t *testing.T) {
	s := NewStackCap[int](10)
	expectEq(t, 0, s.Len())
	expectEq(t, 10, s.Cap())
}

func Test_StackCap(t *testing.T) {
	s := NewStackCap[int](10)
	s.Push(1)
	expectEq(t, 1, s.Len())
	expectEq(t, 10, s.Cap())
}

func Test_Stack_Clean(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Clean()
	expectEq(t, 0, s.Len())
	expectTrue(t, s.IsEmpty())
}

func Test_Stack_Push(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	expectEq(t, 1, s.Len())
	s.Push(2)
	expectEq(t, 2, s.Len())
}

func Test_Stack_Pop(t *testing.T) {
	s := NewStack[int]()
	_, ok := s.Pop()
	expectFalse(t, ok)
	s.Push(1)
	v, ok := s.Pop()
	expectTrue(t, ok)
	expectEq(t, 1, v)
}

func Test_Stack_Must(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	v := s.MustPop()
	expectEq(t, 1, v)
	expactPanic(t, func() { s.MustPop() })
}
