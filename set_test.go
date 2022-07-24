package contalgo

import (
	"fmt"
	"testing"
	"strings"
)

func Test_New(t *testing.T) {
	s := NewSet[string]()
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func Test_Of(t *testing.T) {
	s := NewSetOf("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_IsEmpty(t *testing.T) {
	s := NewSet[string]()
	expectEq(t, s.IsEmpty(), true)
	s.Add("hello")
	expectEq(t, s.IsEmpty(), false)
}

func Test_String(t *testing.T) {
	s := NewSetOf("hello", "world")
	expectTrue(t, strings.HasPrefix(fmt.Sprintf("%v", s), "Set[string]"))
}

func Test_Add(t *testing.T) {
	s := NewSet[string]()
	s.Add("hello")
	s.Add("hello")
	expectEq(t, s.Has("world"), false)
	s.Add("world")
	expectEq(t, s.Has("hello"), true)
	expectEq(t, s.Len(), 2)
}

func Test_AddN(t *testing.T) {
	s := NewSet[string]()
	s.AddN("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_Del(t *testing.T) {
	s := NewSetOf("hello", "world")
	s.Del("hello")
	expectEq(t, s.Len(), 1)
	s.Del("hello")
	expectEq(t, s.Len(), 1)
	s.Del("world")
	expectEq(t, s.Len(), 0)
}

func Test_Clear(t *testing.T) {
	s := NewSetOf("hello", "world")
	s.Clear()
	expectEq(t, s.IsEmpty(), true)
}

func Test_Keys(t *testing.T) {
	s := NewSetOf("hello", "world")
	ks := s.Keys()
	expectEq(t, 2, len(ks))
}

func Test_ForEach(t *testing.T) {
	s := NewSetOf("hello", "world")
	c := 0
	s.ForEach(func(string) {
		c++
	})
	expectEq(t, c, 2)
}

func Test_ForEachIf(t *testing.T) {
	s := NewSetOf("hello", "world")
	c := 0
	s.ForEachIf(func(string) bool {
		c++
		return false
	})
	expectLt(t, c, 2)
}
