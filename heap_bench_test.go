package stl4go

import (
	"container/heap"
	"testing"
)

// An intHeap is a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkHeapInit(b *testing.B) {
	data := Range(1, 1000)
	Shuffle(data)
	b.Run("container/heap.Init", func(b *testing.B) {
		data = Copy(data)
		h := intHeap(data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			heap.Init(&h)
		}
	})
	b.Run("stlgo.MakeMinHeap", func(b *testing.B) {
		data = Copy(data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MakeMinHeap(data)
		}
	})
	b.Run("stlgo.MakeHeapFunc", func(b *testing.B) {
		data = Copy(data)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			MakeHeapFunc(data, Less[int])
		}
	})
}

func BenchmarkHeapPush(b *testing.B) {
	const count = 1000
	b.Run("container/heap.Push", func(b *testing.B) {
		h := intHeap([]int{})
		for i := 0; i < b.N; i++ {
			h = h[0:0]
			for n := 0; n < count; n++ {
				heap.Push(&h, n)
			}
		}
	})
	b.Run("stlgo.PushMinHeap", func(b *testing.B) {
		h := []int{}
		for i := 0; i < b.N; i++ {
			h = h[0:0]
			for n := 0; n < count; n++ {
				PushMinHeap(&h, n)
			}
		}
	})
	b.Run("stlgo.PushHeapFunc", func(b *testing.B) {
		h := []int{}
		for i := 0; i < b.N; i++ {
			h = h[0:0]
			for n := 0; n < count; n++ {
				PushHeapFunc(&h, n, Less[int])
			}
		}
	})
}

func BenchmarkHeapPop(b *testing.B) {
	s := Range(1, 1000)
	b.Run("container/heap.Pop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			h := intHeap(Copy(s))
			heap.Init(&h)
			b.StartTimer()
			for len(h) > 0 {
				_ = heap.Pop(&h).(int)
			}
		}
	})
	b.Run("stlgo.PopMinHeap", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			h := Copy(s)
			MakeMinHeap(h)
			b.StartTimer()
			for len(h) > 0 {
				_ = PopMinHeap(&h)
			}
		}
	})
	b.Run("stlgo.PopHeapFunc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			h := Copy(s)
			MakeMinHeap(h)
			b.StartTimer()
			for len(h) > 0 {
				_ = PopHeapFunc(&h, Less[int])
			}
		}
	})
}
