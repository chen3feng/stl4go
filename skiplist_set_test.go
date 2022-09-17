package stl4go

import (
	"testing"
)

func TestNewSkipList_Interface(t *testing.T) {
	s := NewSkipListSet[int]()
	_ = SortedSet[int](s)
}

func TestNewSkipListSet(t *testing.T) {
	s := NewSkipListSet[int]()
	expectTrue(t, s.IsEmpty())
	expectEq(t, s.Len(), 0)
}

func TestNewSkipListSetOf(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 3, 4, 5)
	expectTrue(t, Equal(s.Keys(), []int{1, 2, 3, 4, 5}))
}

func TestNewSkipListSetFunc(t *testing.T) {
	s := NewSkipListSetFunc(OrderedCompare[int])
	s.InsertN(1, 2, 3, 4, 5)
	expectTrue(t, Equal(s.Keys(), []int{1, 2, 3, 4, 5}))
}

func TestNewSkipListSet_Insert_Remove_Has(t *testing.T) {
	s := NewSkipListSet[int]()

	expectFalse(t, s.Remove(1))
	expectFalse(t, s.Has(1))

	expectTrue(t, s.Insert(1))
	expectFalse(t, s.Insert(1))
	expectTrue(t, s.Has(1))
	expectFalse(t, s.IsEmpty())
	expectEq(t, s.Len(), 1)

	expectTrue(t, s.Remove(1))
	expectFalse(t, s.Remove(1))
	expectFalse(t, s.Has(1))

	expectTrue(t, s.IsEmpty())
	expectEq(t, s.Len(), 0)
}

func TestNewSkipListSet_Clear(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 3, 4, 5)
	s.Clear()
	expectTrue(t, s.IsEmpty())
	expectEq(t, s.Len(), 0)
}

func TestNewSkipListSet_InsertN(t *testing.T) {
	s := NewSkipListSet[int]()
	expectEq(t, s.InsertN(1, 2, 3, 1), 3)
	expectFalse(t, s.IsEmpty())
	expectEq(t, s.Len(), 3)
}

func TestNewSkipListSet_RemoveN(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 3, 4)
	expectEq(t, s.RemoveN(2, 4, 6), 2)
	expectFalse(t, s.IsEmpty())
	expectEq(t, s.Len(), 2)
	expectTrue(t, Equal(s.Keys(), []int{1, 3}))
}

func TestNewSkipListSet_ForEach(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 3, 4)
	var a []int
	s.ForEach(func(i int) { a = append(a, i) })
	expectTrue(t, Equal(a, []int{1, 2, 3, 4}))
}

func TestNewSkipListSet_ForEachIf(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 3, 4)
	var a []int
	s.ForEachIf(func(i int) bool { a = append(a, i); return i < 2 })
	expectTrue(t, Equal(a, []int{1, 2}))
}

func TestNewSkipList_Iterate(t *testing.T) {
	sl := newSkipListN(10)
	i := 0
	for it := sl.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Key(), i)
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), i)
		i++
	}
}

func testSkipListSetIterater(t *testing.T, sl *SkipListSet[int]) {

}

func TestSkipListSet_Iterater(t *testing.T) {
	s := NewSkipListSetOf(1, 2, 4)
	t.Run("LowerBound", func(t *testing.T) {
		expectEq(t, s.LowerBound(1).Value(), 1)
		expectEq(t, s.LowerBound(3).Value(), 4)
		expectEq(t, s.LowerBound(4).Value(), 4)
		expectFalse(t, s.LowerBound(5).IsNotEnd())
	})

	t.Run("UpperBound", func(t *testing.T) {
		expectEq(t, s.UpperBound(0).Value(), 1)
		expectEq(t, s.UpperBound(1).Value(), 2)
		expectEq(t, s.UpperBound(3).Value(), 4)
		expectFalse(t, s.UpperBound(4).IsNotEnd())
		expectFalse(t, s.UpperBound(5).IsNotEnd())
	})

	t.Run("FindRange", func(t *testing.T) {
		it := s.FindRange(1, 3)
		expectEq(t, it.Value(), 1)
		it.MoveToNext()
		expectEq(t, it.Value(), 2)
	})
}
