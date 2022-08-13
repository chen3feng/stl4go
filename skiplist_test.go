package stl4go

import (
	"math"
	"testing"
)

func TestSkipList_Interface(t *testing.T) {
	sl := NewSkipList[int, int]()
	_ = Map[int, int](sl)
	_ = MapIterator[int, int](sl.Iterate())
}

func TestNewSkipList(t *testing.T) {
	NewSkipList[int, int]()
}

func TestNewSkipListString(t *testing.T) {
	sl := NewSkipList[string, int]()
	sl.Insert("hello", 1)
	expectTrue(t, sl.Has("hello"))
	expectEq(t, *sl.Find("hello"), 1)
	expectFalse(t, sl.Has("world"))
	expectEq(t, sl.Find("world"), nil)
}

func testNewSkipListType[T Numeric](t *testing.T) {
	sl := NewSkipList[T, int]()
	var n T = 1
	sl.Insert(n, 1)
	expectTrue(t, sl.Has(n))
	expectEq(t, *sl.Find(n), 1)
	expectFalse(t, sl.Has(n+1))
	expectEq(t, sl.Find(n+1), nil)
}

func TestNewSkipListInt8(t *testing.T)  { testNewSkipListType[int8](t) }
func TestNewSkipListInt16(t *testing.T) { testNewSkipListType[int16](t) }
func TestNewSkipListInt32(t *testing.T) { testNewSkipListType[int32](t) }
func TestNewSkipListInt64(t *testing.T) { testNewSkipListType[int64](t) }

func TestNewSkipListUInt8(t *testing.T)  { testNewSkipListType[uint8](t) }
func TestNewSkipListUInt16(t *testing.T) { testNewSkipListType[uint16](t) }
func TestNewSkipListUInt32(t *testing.T) { testNewSkipListType[uint32](t) }
func TestNewSkipListUInt64(t *testing.T) { testNewSkipListType[uint64](t) }

func TestNewSkipListUIntPtr(t *testing.T) { testNewSkipListType[uintptr](t) }

func TestNewSkipListFloat32(t *testing.T) { testNewSkipListType[float32](t) }
func TestNewSkipListFloat64(t *testing.T) { testNewSkipListType[float64](t) }

func TestNewSkipListFunc(t *testing.T) {
	type Person struct {
		name string
		age  int
	}
	personCmp := func(a, b Person) int {
		r := OrderedCompare(a.age, b.age)
		if r != 0 {
			return r
		}
		return OrderedCompare(a.name, b.name)
	}
	sl := NewSkipListFunc[Person, int](personCmp)
	sl.Insert(Person{"zhangsan", 20}, 1)
	sl.Insert(Person{"lisi", 20}, 1)
	sl.Insert(Person{"wangwu", 30}, 1)
	expectTrue(t, sl.Has(Person{"zhangsan", 20}))
	expectFalse(t, sl.Has(Person{"zhangsan", 30}))
	expectEq(t, sl.Len(), 3)

	sl.Insert(Person{"zhangsan", 20}, 1)
	expectEq(t, sl.Len(), 3)

	var ps []Person
	sl.ForEach(func(p Person, _ int) {
		ps = append(ps, p)
	})
	expectEq(t, ps[0].name, "lisi")
	expectEq(t, ps[1].name, "zhangsan")
	expectEq(t, ps[2].name, "wangwu")

	sl.Remove(Person{"zhangsan", 20})
	expectEq(t, sl.Len(), 2)

	sl.Remove(Person{"zhaoliu", 40})
	expectEq(t, sl.Len(), 2)
}

func TestNewSkipListFromMap(t *testing.T) {
	m := map[int]int{1: -1, 2: -2, 3: -3}
	sl := NewSkipListFromMap(m)
	for k := range m {
		expectTrue(t, sl.Has(k))
	}
}

func TestNewSkipList_Iterate(t *testing.T) {
	sl := newSkipListN(10)
	i := 0
	for it := sl.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Key(), i)
		expectEq(t, it.Value(), i)
		i++
	}
}

func testNewSkipList_Iterater(t *testing.T, sl *SkipList[int, int]) {
	t.Run("LowerBound", func(t *testing.T) {
		expectEq(t, sl.LowerBound(1).Key(), 1)
		expectEq(t, sl.LowerBound(3).Key(), 4)
		expectEq(t, sl.LowerBound(4).Key(), 4)
		expectFalse(t, sl.LowerBound(5).IsNotEnd())
	})

	t.Run("UpperBound", func(t *testing.T) {
		expectEq(t, sl.UpperBound(0).Key(), 1)
		expectEq(t, sl.UpperBound(1).Key(), 2)
		expectEq(t, sl.UpperBound(3).Key(), 4)
		expectFalse(t, sl.UpperBound(4).IsNotEnd())
		expectFalse(t, sl.UpperBound(5).IsNotEnd())
	})

	t.Run("FindRange", func(t *testing.T) {
		it := sl.FindRange(1, 3)
		expectEq(t, it.Key(), 1)
		it.MoveToNext()
		expectEq(t, it.Key(), 2)
	})
}

func TestNewSkipList_Iterater(t *testing.T) {
	sl := NewSkipListFromMap(map[int]int{1: 1, 2: 2, 4: 4})
	testNewSkipList_Iterater(t, sl)
}

func TestNewSkipList_Func_Iterater(t *testing.T) {
	sl := NewSkipListFunc[int, int](OrderedCompare[int])
	m := map[int]int{1: 1, 2: 2, 4: 4}
	for k, v := range m {
		sl.Insert(k, v)
	}
	testNewSkipList_Iterater(t, sl)
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
	sl := newSkipListN(100)
	expectGe(t, sl.level, 4)
	for i := 0; i < 100; i++ {
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
	sl.ForEach(func(k int, v int) {
		a = append(a, k)
	})
	expectEq(t, len(a), 100)
	expectTrue(t, IsSorted(a))
}

func TestSkipList_ForEachIf(t *testing.T) {
	sl := newSkipListN(100)
	a := []int{}
	sl.ForEachIf(func(k int, v int) bool {
		if k < 50 {
			a = append(a, k)
			return true
		}
		return false
	})
	expectLt(t, MaxN(a...), 50)
}

func TestSkipList_ForEachMutable(t *testing.T) {
	sl := newSkipListN(100)
	sl.ForEachMutable(func(k int, v *int) {
		*v = -*v
	})
	for i := 0; i < sl.Len(); i++ {
		expectEq(t, *sl.Find(i), -i)
	}
}

func TestSkipList_ForEachMutableIf(t *testing.T) {
	sl := newSkipListN(100)
	sl.ForEachMutableIf(func(k int, v *int) bool {
		if k > 50 {
			*v = 0
			return false
		}
		return true
	})

	expectEq(t, *sl.Find(51), 0)
}
