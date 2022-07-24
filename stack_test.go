package contalgo

import (
	"testing"
)

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
