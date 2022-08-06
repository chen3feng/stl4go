package stl4go

import (
	"math"
	"testing"
)

func TestSkipListInterface(t *testing.T) {
	_ = Container(NewSkipList[int, int]())
}

func TestNewSkipList(t *testing.T) {
	NewSkipList[int, int]()
}

func TestNewSkipListFromMap(t *testing.T) {
	NewSkipListFromMap(map[int]int{1: 1, 2: 2, 3: 3})
}

func TestSkipList_Insert(t *testing.T) {
	sl := NewSkipList[int, int]()
	for i := 0; i < 100; i++ {
		expectEq(t, sl.Len(), i)
		sl.Insert(i, i)
		expectEq(t, sl.Len(), i+1)
	}
}

func TestSkipList_Insert_Reverse(t *testing.T) {
	sl := NewSkipList[int, int]()
	for i := 100; i > 0; i-- {
		oldlen := sl.Len()
		sl.Insert(i, i)
		expectEq(t, sl.Len(), oldlen+1)
	}
}

func TestSkipList_Insert_Dup(t *testing.T) {
	sl := NewSkipList[int, int]()
	sl.Insert(1, 1)
	expectEq(t, sl.Len(), 1)
	sl.Insert(1, 2)
	expectEq(t, sl.Len(), 1)
}

func newSkipListN(n int) *SkipList[int, int] {
	sl := NewSkipList[int, int]()
	for i := 0; i < n; i++ {
		sl.Insert(i, i)
	}
	return sl
}

func TestSkipList_Remove(t *testing.T) {
	sl := newSkipListN(100)
	for i := 0; i < 100; i++ {
		sl.Remove(i)
	}
	expectTrue(t, sl.IsEmpty())
	expectEq(t, sl.Len(), 0)
}

func TestSkipList_Remove_Nonexist(t *testing.T) {
	sl := NewSkipList[int, int]()
	sl.Insert(1, 1)
	sl.Insert(2, 2)
	sl.Remove(0)
	sl.Remove(3)
	expectEq(t, sl.Len(), 2)
}

func TestSkipList_Remove_Level(t *testing.T) {
	sl := newSkipListN(1000)
	expectGt(t, sl.level, 8)
	for i := 0; i < 1000; i++ {
		sl.Remove(i)
	}
	expectEq(t, sl.level, 1)
}

func TestSkipList_Clean(t *testing.T) {
	sl := NewSkipList[int, int]()
	for i := 0; i < 100; i++ {
		sl.Insert(i, i)
	}
	sl.Clear()

	expectTrue(t, sl.IsEmpty())
	expectEq(t, sl.Len(), 0)
	expectEq(t, sl.level, 1)
}

func TestSkipList_level(t *testing.T) {
	var diffs []int
	for i := 0; i < 1000; i++ {
		for size := 1; size < 10000; size *= 10 {
			sl := newSkipListN(size)
			log2Len := int(math.Ceil(math.Log2(float64(sl.Len()))))
			diffs = append(diffs, log2Len-sl.level)
		}
	}
	expectLt(t, math.Abs(float64(Average(diffs))), 2)
}

func TestSkipList_newnode(t *testing.T) {
	for level := 1; level <= skipListMaxLevel; level++ {
		node := newSkipListNode(level, 1, 1)
		expectEq(t, len(node.next), level)
	}
	expactPanic(t, func() { newSkipListNode(0, 1, 1) })
	expactPanic(t, func() { newSkipListNode(skipListMaxLevel+1, 1, 1) })
}

func TestSkipList_Find(t *testing.T) {
	sl := newSkipListN(100)
	for i := 0; i < 100; i++ {
		p := sl.Find(i)
		expectEq(t, i, *p)
	}
	expectEq(t, sl.Find(100), nil)
}

func TestSkipList_Has(t *testing.T) {
	sl := NewSkipList[int, int]()
	expectFalse(t, sl.Has(1))
	sl.Insert(1, 2)
	expectTrue(t, sl.Has(1))
}

func TestSkipList_ForEach(t *testing.T) {
	sl := newSkipListN(100)
	a := []int{}
	sl.ForEach(func(k int, v *int) {
		a = append(a, k)
	})
	expectEq(t, len(a), 100)
	expectTrue(t, IsSorted(a))
}

func isEven(n int) bool { return n%2 == 0 }

func TestSkipList_ForEachIf(t *testing.T) {
	sl := newSkipListN(100)
	a := []int{}
	sl.ForEachIf(func(k int, v *int) bool {
		if k < 50 {
			a = append(a, k)
			return true
		}
		return false
	})
	expectLt(t, MaxN(a...), 50)
}
