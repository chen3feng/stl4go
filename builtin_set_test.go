package stl4go

import (
	"fmt"
	"strings"
	"testing"
)

func Test_BuiltinSet_Interface(t *testing.T) {
	s := make(BuiltinSet[int])
	_ = Set[int](&s)
}

func Test_MakeBuiltinSet(t *testing.T) {
	s := make(BuiltinSet[string])
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func Test_MakeBuiltinSet2(t *testing.T) {
	s := BuiltinSet[string]{}
	expectEq(t, s.Len(), 0)
	expectEq(t, s.IsEmpty(), true)
}

func Test_MakeBuiltinSetOf(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_BuiltinSet_IsEmpty(t *testing.T) {
	s := make(BuiltinSet[string])
	expectEq(t, s.IsEmpty(), true)
	s.Insert("hello")
	expectEq(t, s.IsEmpty(), false)
}

func Test_BuiltinSet_Clear(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	s.Clear()
	expectTrue(t, s.IsEmpty())
}

func Test_BuiltinSet_String(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	expectTrue(t, strings.HasPrefix(fmt.Sprintf("%v", s), "BuiltinSet[string]"))
}

func Test_BuiltinSet_Has(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	expectTrue(t, s.Has("hello"))
	expectTrue(t, s.Has("world"))
	expectFalse(t, s.Has("!"))
}

func Test_BuiltinSet_Get(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	expectTrue(t, s["hello"])
	expectTrue(t, s["world"])
	expectFalse(t, s["!"])
}

func Test_BuiltinSet_Insert(t *testing.T) {
	s := make(BuiltinSet[string])
	s.Insert("hello")
	s.Insert("hello")
	expectEq(t, s.Has("world"), false)
	s.Insert("world")
	expectEq(t, s.Has("hello"), true)
	expectEq(t, s.Len(), 2)
}

func Test_BuiltinSet_InsertN(t *testing.T) {
	s := make(BuiltinSet[string])
	s.InsertN("hello", "world")
	expectEq(t, s.Len(), 2)
}

func Test_BuiltinSet_Remove(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	s.Remove("hello")
	expectEq(t, s.Len(), 1)
	s.Remove("hello")
	expectEq(t, s.Len(), 1)
	s.Remove("world")
	expectEq(t, s.Len(), 0)
}

func Test_BuiltinSet_RemoveN(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	s.RemoveN("hello", "world")
	s.Remove("world")
	expectTrue(t, s.IsEmpty())
}

func Test_BuiltinSet_Keys(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	ks := s.Keys()
	expectEq(t, 2, len(ks))
}

func Test_BuiltinSet_For(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	for v := range s {
		expectTrue(t, v == "hello" || v == "world")
	}
}

func Test_BuiltinSet_ForEach(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	c := 0
	s.ForEach(func(string) {
		c++
	})
	expectEq(t, c, 2)
}

func Test_BuiltinSet_ForEachIf(t *testing.T) {
	s := MakeBuiltinSetOf("hello", "world")
	c := 0
	s.ForEachIf(func(string) bool {
		c++
		return false
	})
	expectLt(t, c, 2)
}
