package stl4go

import (
	"testing"
)

func Test_SList_Interface(t *testing.T) {
	sl := SList[int]{}
	_ = Container(&sl)
}

func Test_SList_Clean(t *testing.T) {
	sl := SList[int]{}
	sl.PushFront(1)
	sl.Clear()
	expectTrue(t, sl.IsEmpty())
	expectEq(t, sl.Len(), 0)
}

func Test_SList_Front(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sl := SList[int]{}
		expectPanic(t, func() { sl.Front() })
	})

	t.Run("normal", func(t *testing.T) {
		sl := SList[int]{}
		sl.PushFront(1)
		expectEq(t, sl.Front(), 1)

		sl.PushBack(2)
		expectEq(t, sl.Front(), 1)

		sl.PushFront(3)
		expectEq(t, sl.Front(), 3)
	})
}

func Test_SList_Back(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sl := SList[int]{}
		expectPanic(t, func() { sl.Back() })
	})

	t.Run("normal", func(t *testing.T) {
		sl := SList[int]{}
		sl.PushBack(1)
		expectEq(t, sl.Back(), 1)

		sl.PushFront(2)
		expectEq(t, sl.Back(), 1)

		sl.PushBack(3)
		expectEq(t, sl.Back(), 3)
	})
}

func Test_SList_PushFront(t *testing.T) {
	sl := SList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushFront(i)
		expectEq(t, sl.Front(), i)
		expectEq(t, sl.Len(), i)
	}
}

func Test_SList_PushBack(t *testing.T) {
	sl := SList[int]{}
	for i := 1; i < 10; i++ {
		sl.PushBack(i)
		expectEq(t, sl.Back(), i)
		expectEq(t, sl.Len(), i)
		expectFalse(t, sl.IsEmpty())
	}
}

func Test_SList_PopFront(t *testing.T) {
	sl := SList[int]{}
	expectPanic(t, func() { sl.PopFront() })

	sl.PushFront(1)
	sl.PushFront(2)
	expectEq(t, sl.PopFront(), 2)
	expectEq(t, sl.PopFront(), 1)
	expectPanic(t, func() { sl.PopFront() })
}

func Test_SList_Reverse(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	sl.Reverse()
	expectTrue(t, Equal(sl.Values(), []int{4, 3, 2, 1}))
}

func Test_SList_ForEach(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEach(func(v int) {
		i++
		expectEq(t, v, i)
	})
	expectEq(t, i, sl.Len())
}

func Test_SList_ForEachIf(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachIf(func(v int) bool {
		i++
		expectEq(t, v, i)
		return i < 3
	})
	expectEq(t, i, 3)
}

func Test_SList_ForEachMutable(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachMutable(func(v *int) {
		i++
		expectEq(t, *v, i)
		*v = -*v
	})
	expectEq(t, i, sl.Len())
	sl.ForEachMutable(func(v *int) {
		expectLt(t, *v, 0)
	})
}

func Test_SList_ForEachMutableIf(t *testing.T) {
	sl := SListOf(1, 2, 3, 4)
	i := 0
	sl.ForEachMutableIf(func(v *int) bool {
		i++
		expectEq(t, *v, i)
		return i < 3
	})
	expectEq(t, i, 3)
}

func Test_SList_Iterate(t *testing.T) {
	sl := SList[int]{}
	sl.PushBack(1)
	sl.PushBack(2)
	sl.PushBack(3)
	i := 0
	for it := sl.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		i++
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), i)
	}
	expectEq(t, i, 3)
}
