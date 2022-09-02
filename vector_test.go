package stl4go

import (
	"testing"
)

func Test_Vector_Interface(t *testing.T) {
	v := MakeVector[int]()
	_ = Container(&v)
}

func Test_MakeVector(t *testing.T) {
	MakeVector[int]()
	_ = make(Vector[int], 1)
	_ = make(Vector[int], 1, 2)
	var v Vector[int]
	_ = v
	v = Vector[int]{1, 2, 3}
}

func Test_MakeVectorCap(t *testing.T) {
	v := MakeVectorCap[int](10)
	expectEq(t, v.Len(), 0)
	expectEq(t, v.Cap(), 10)
}

func Test_VectorOf(t *testing.T) {
	v := VectorOf(1, 2, 3)
	expectEq(t, v.Len(), 3)
	expectEq(t, v.Cap(), 3)
	expectEq(t, v[0], 1)
	expectEq(t, v[1], 2)
	expectEq(t, v[2], 3)
}

func Test_AsVector(t *testing.T) {
	s := []int{1, 2, 3}
	v := AsVector(s)
	expectTrue(t, Equal(s, v))
	expectEq(t, &s[0], &v[0])
}

func Test_VectorCap(t *testing.T) {
	v := MakeVectorCap[int](10)
	v.PushBack(1)
	expectEq(t, v.Len(), 1)
	expectFalse(t, v.IsEmpty())
	expectEq(t, v.Cap(), 10)
}

func Test_Vector_Clear(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Clear()
	expectEq(t, v.Len(), 0)
	expectTrue(t, v.IsEmpty())
	expectGt(t, v.Cap(), 0)
}

func Test_Vector_Reserve(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Reserve(1)
	expectEq(t, v.Cap(), 3)
	v.Reserve(5)
	expectEq(t, v.Cap(), 5)
	expectEq(t, v.Len(), 3)
}

func Test_Vector_Shrink(t *testing.T) {
	v := MakeVectorCap[int](10)
	v.Append(1, 2, 3)
	expectEq(t, v.Cap(), 10)
	v.Shrink()
	expectEq(t, v.Len(), v.Cap())
}

func Test_Vector_At_Set(t *testing.T) {
	v := VectorOf(1, 2, 3)
	expectEq(t, v.At(0), 1)
	expectEq(t, v[0], 1)
	v[0] = 2
	expectEq(t, v[0], 2)
	expectPanic(t, func() { v.Set(3, 2) })
}

func Test_Vector_PushBack(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.PushBack(4)
	expectTrue(t, Equal(v, []int{1, 2, 3, 4}))
}

func Test_Vector_PopBack(t *testing.T) {
	v := VectorOf(1, 2)
	expectEq(t, v.PopBack(), 2)
	n, ok := v.TryPopBack()
	expectEq(t, n, 1)
	expectTrue(t, ok)
	n, ok = v.TryPopBack()
	expectEq(t, n, 0)
	expectFalse(t, ok)
	expectPanic(t, func() { v.PopBack() })
}

func Test_Vector_Back(t *testing.T) {
	v := VectorOf(1)
	expectEq(t, v.Back(), 1)
	v.PopBack()
	expectPanic(t, func() { v.Back() })
}

func Test_Vector_Insert(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Insert(0, 1, 2, 3)
	expectTrue(t, Equal(v, []int{1, 2, 3, 1, 2, 3}))
}

func Test_Vector_Insert_Tail(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Insert(3, 1, 2, 3)
	expectTrue(t, Equal(v, []int{1, 2, 3, 1, 2, 3}))
}

func Test_Vector_Insert_Mid(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Insert(2, 1, 2)
	expectTrue(t, Equal(v, []int{1, 2, 1, 2, 3}))
}

func Test_Vector_Insert_Cap(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Reserve(8)
	v.Insert(2, 1, 2)
	expectTrue(t, Equal(v, []int{1, 2, 1, 2, 3}))
}

func Test_Vector_Remove(t *testing.T) {
	v := VectorOf(1, 2, 3)
	v.Remove(1)
	expectEq(t, v.Len(), 2)
	expectEq(t, v.Cap(), 3)
	expectEq(t, v[0], 1)
	expectEq(t, v[1], 3)
}

func Test_Vector_RemoveRange(t *testing.T) {
	v := VectorOf(1, 2, 3, 4)
	v.RemoveRange(1, 3)
	expectEq(t, v.Len(), 2)
	expectEq(t, v.Cap(), 4)
	expectEq(t, v[0], 1)
	expectEq(t, v[1], 4)
}

func Test_Vector_RemoveLength(t *testing.T) {
	v := VectorOf(1, 2, 3, 4)
	v.RemoveLength(1, 2)
	expectEq(t, v.Len(), 2)
	expectEq(t, v.Cap(), 4)
	expectEq(t, v[0], 1)
	expectEq(t, v[1], 4)
}

func Test_Vector_ForEach(t *testing.T) {
	a := []int{1, 2, 3}
	v := VectorOf(a...)
	var b []int
	v.ForEach(func(n int) {
		b = append(b, n)
	})
	expectEq(t, len(b), 3)
	expectTrue(t, Equal(a, b))
}

func Test_Vector_ForEachIf(t *testing.T) {
	v := VectorOf(1, 2, 3)
	c := 0
	v.ForEachIf(func(n int) bool {
		c = n
		return n != 2
	})
	expectEq(t, c, 2)
}

func Test_Vector_ForEachMutable(t *testing.T) {
	a := []int{1, 2, 3}
	v := VectorOf(1, 2, 3)
	v.ForEachMutable(func(n *int) {
		*n = -*n
	})
	expectEq(t, v.Len(), 3)
	for i := range v {
		expectEq(t, a[i], -v[i])
	}
}

func Test_Vector_ForEachMutableIf(t *testing.T) {
	v := VectorOf(1, 2, 3)
	c := 0
	v.ForEachMutableIf(func(n *int) bool {
		c = *n
		return *n != 2
	})
	expectEq(t, c, 2)
}

func Test_Vector_Iterate(t *testing.T) {
	v := VectorOf(1, 2, 3, 4)
	i := 1
	for it := v.Iterate(); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), it.Value())
		i++
	}
	expectEq(t, i, 5)
}

func Test_Vector_IterateRange(t *testing.T) {
	v := VectorOf(1, 2, 3, 4)
	i := 2
	for it := v.IterateRange(1, 3); it.IsNotEnd(); it.MoveToNext() {
		expectEq(t, it.Value(), i)
		expectEq(t, *it.Pointer(), it.Value())
		i++
	}
	expectEq(t, i, 4)
}
