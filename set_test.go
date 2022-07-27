package contalgo

import (
	"fmt"
	"strings"
	"testing"
)

func Test_Set_Interface(t *testing.T) {
	s := make(Set[int])
	_ = Container[int](&s)
}

func Test_MakeSet(t *testing.T) {
	s := make(Set[string])
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func Test_MakeSet2(t *testing.T) {
	s := Set[string]{}
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func Test_MakeSetOf(t *testing.T) {
	s := MakeSetOf("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_Set_IsEmpty(t *testing.T) {
	s := make(Set[string])
	expectEq(t, s.IsEmpty(), true)
	s.Add("hello")
	expectEq(t, s.IsEmpty(), false)
}

func Test_Set_Clean(t *testing.T) {
	s := MakeSetOf("hello", "world")
	s.Clean()
	expectTrue(t, s.IsEmpty())
}

func Test_Set_String(t *testing.T) {
	s := MakeSetOf("hello", "world")
	expectTrue(t, strings.HasPrefix(fmt.Sprintf("%v", s), "Set[string]"))
}

func Test_Set_Has(t *testing.T) {
	s := MakeSetOf("hello", "world")
	expectTrue(t, s.Has("hello"))
	expectTrue(t, s.Has("world"))
	expectFalse(t, s.Has("!"))
}

func Test_Set_Get(t *testing.T) {
	s := MakeSetOf("hello", "world")
	expectTrue(t, s["hello"])
	expectTrue(t, s["world"])
	expectFalse(t, s["!"])
}

func Test_Set_Add(t *testing.T) {
	s := make(Set[string])
	s.Add("hello")
	s.Add("hello")
	expectEq(t, s.Has("world"), false)
	s.Add("world")
	expectEq(t, s.Has("hello"), true)
	expectEq(t, s.Len(), 2)
}

func Test_Set_AddN(t *testing.T) {
	s := make(Set[string])
	s.AddN("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_Set_Del(t *testing.T) {
	s := MakeSetOf("hello", "world")
	s.Del("hello")
	expectEq(t, s.Len(), 1)
	s.Del("hello")
	expectEq(t, s.Len(), 1)
	s.Del("world")
	expectEq(t, s.Len(), 0)
}

func Test_Set_DelN(t *testing.T) {
	s := MakeSetOf("hello", "world")
	s.DelN("hello", "world")
	s.Del("world")
	expectTrue(t, s.IsEmpty())
}

func Test_Set_Keys(t *testing.T) {
	s := MakeSetOf("hello", "world")
	ks := s.Keys()
	expectEq(t, 2, len(ks))
}

func Test_Set_For(t *testing.T) {
	s := MakeSetOf("hello", "world")
	for v := range s {
		expectTrue(t, v == "hello" || v == "world")
	}
}

func Test_Set_ForEach(t *testing.T) {
	s := MakeSetOf("hello", "world")
	c := 0
	s.ForEach(func(string) {
		c++
	})
	expectEq(t, c, 2)
}

func Test_Set_ForEachIf(t *testing.T) {
	s := MakeSetOf("hello", "world")
	c := 0
	s.ForEachIf(func(string) bool {
		c++
		return false
	})
	expectLt(t, c, 2)
}
