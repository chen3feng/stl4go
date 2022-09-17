package stl4go

import (
	"testing"
)

func TestMakeMinHeap(t *testing.T) {
	data := []int{5, 4, 3, 2, 1}
	expectFalse(t, IsMinHeap(data))
	MakeMinHeap(data)
	expectTrue(t, IsMinHeap(data))
}

func TestIsMinHeap(t *testing.T) {
	heap := []int{}
	expectTrue(t, IsMinHeap(heap))
	heap = append(heap, 1)
	expectTrue(t, IsMinHeap(heap))
}

func TestMakeHeapFunc(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expectFalse(t, IsHeapFunc(data, Greater[int]))
	MakeHeapFunc(data, Greater[int])
	expectTrue(t, IsHeapFunc(data, Greater[int]))
}

func TestMinHeap_PushPop(t *testing.T) {
	heap := []int{}
	PushMinHeap(&heap, 1)
	expectEq(t, PopMinHeap(&heap), 1)
	expectTrue(t, IsMinHeap(heap))
}

func TestHeapFunc_PushPop(t *testing.T) {
	heap := []int{}
	cmp := Greater[int]
	PushHeapFunc(&heap, 1, cmp)
	expectTrue(t, IsHeapFunc(heap, cmp))
	expectEq(t, PopHeapFunc(&heap, cmp), 1)
	expectTrue(t, IsHeapFunc(heap, cmp))
}

func TestMinHeap_Remove(t *testing.T) {
	heap := []int{5, 4, 3, 2, 1}
	MakeMinHeap(heap)
	RemoveMinHeap(&heap, 1)
	expectTrue(t, IsMinHeap(heap))
}

func TestHeapFunc_Remove(t *testing.T) {
	heap := []int{1, 2, 3, 4, 5}
	cmp := Greater[int]
	MakeHeapFunc(heap, cmp)
	RemoveHeapFunc(&heap, 1, cmp)
	expectTrue(t, IsHeapFunc(heap, cmp))
}
