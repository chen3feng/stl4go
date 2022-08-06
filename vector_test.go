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
	expectEq(t, 0, v.Len())
	expectEq(t, 10, v.Cap())
}

func Test_MakeVectorOf(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	expectEq(t, 3, v.Len())
	expectEq(t, 3, v.Cap())
	expectEq(t, v.At(0), 1)
	expectEq(t, v.At(1), 2)
	expectEq(t, v.At(2), 3)
}

func Test_VectorCap(t *testing.T) {
	v := MakeVectorCap[int](10)
	v.PushBack(1)
	expectEq(t, 1, v.Len())
	expectFalse(t, v.IsEmpty())
	expectEq(t, 10, v.Cap())
}

func Test_Vector_Clear(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Clear()
	expectEq(t, 0, v.Len())
	expectTrue(t, v.IsEmpty())
	expectGt(t, v.Cap(), 0)
}

func Test_Vector_Reserve(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Reserve(1)
	expectEq(t, 3, v.Cap())
	v.Reserve(5)
	expectEq(t, 5, v.Cap())
	expectEq(t, 3, v.Len())
}

func Test_Vector_Shrink(t *testing.T) {
	v := MakeVectorCap[int](10)
	v.Append(1, 2, 3)
	expectEq(t, 10, v.Cap())
	v.Shrink()
	expectEq(t, v.Len(), v.Cap())
}

func Test_Vector_At_Set(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	expectEq(t, v.At(0), 1)
	expectEq(t, v[0], 1)
	v[0] = 2
	expectEq(t, v[0], 2)
	expactPanic(t, func() { v.Set(3, 2) })
}

func Test_Vector_Insert(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Insert(0, 1, 2, 3)
	expectTrue(t, Equal(v, []int{1, 2, 3, 1, 2, 3}))
}

func Test_Vector_Insert_Tail(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Insert(3, 1, 2, 3)
	expectTrue(t, Equal(v, []int{1, 2, 3, 1, 2, 3}))
}

func Test_Vector_Insert_Mid(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Insert(2, 1, 2)
	expectTrue(t, Equal(v, []int{1, 2, 1, 2, 3}))
}

func Test_Vector_Insert_Cap(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Reserve(8)
	v.Insert(2, 1, 2)
	expectTrue(t, Equal(v, []int{1, 2, 1, 2, 3}))
}

func Test_Vector_Remove(t *testing.T) {
	v := MakeVectorOf(1, 2, 3)
	v.Remove(1)
	expectEq(t, 2, v.Len())
	expectEq(t, v.Cap(), 3)
	expectEq(t, 1, v[0])
	expectEq(t, 3, v[1])
}

func Test_Vector_RemoveRange(t *testing.T) {
	v := MakeVectorOf(1, 2, 3, 4)
	v.RemoveRange(1, 3)
	expectEq(t, 2, v.Len())
	expectEq(t, v.Cap(), 4)
	expectEq(t, 1, v[0])
	expectEq(t, 4, v[1])
}

func Test_Vector_RemoveLength(t *testing.T) {
	v := MakeVectorOf(1, 2, 3, 4)
	v.RemoveLength(1, 2)
	expectEq(t, 2, v.Len())
	expectEq(t, v.Cap(), 4)
	expectEq(t, 1, v[0])
	expectEq(t, 4, v[1])
}
