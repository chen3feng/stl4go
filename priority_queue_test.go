package stl4go

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	_ = (Container)(NewPriorityQueue[int]())
}

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue[int]()
	expectTrue(t, pq.IsEmpty())
	expectEq(t, pq.Len(), 0)
}

func TestNewPriorityQueueOf(t *testing.T) {
	pq := NewPriorityQueueOf(5, 4, 3, 2, 1)
	expectEq(t, pq.Len(), 5)
}

func TestPriorityQueue_PushPop(t *testing.T) {
	less := Less[int]
	pq := NewPriorityQueueFunc(less)
	for i := 5; i > 0; i-- {
		pq.Push(i)
		expectFalse(t, pq.IsEmpty())
	}
	var elements []int
	for !pq.IsEmpty() {
		elements = append(elements, pq.Pop())
	}
	expectTrue(t, pq.IsEmpty())
	expectTrue(t, IsSorted(elements))
	expectEq(t, len(elements), 5)
}

func TestPriorityQueueFunc_PushPop(t *testing.T) {
	pq := NewPriorityQueueFunc(Less[int])
	for i := 5; i > 0; i-- {
		pq.Push(i)
	}
	var elements []int
	for !pq.IsEmpty() {
		elements = append(elements, pq.Pop())
	}
	expectTrue(t, pq.IsEmpty())
	expectTrue(t, IsSorted(elements))
	expectEq(t, len(elements), 5)
}

func TestPriorityQueue_Clear(t *testing.T) {
	pq := NewPriorityQueue[int]()
	pq.Clear()
	expectTrue(t, pq.IsEmpty())
	expectEq(t, pq.Len(), 0)
	pq = NewPriorityQueueOf(1, 2, 3)
	pq.Clear()
	expectTrue(t, pq.IsEmpty())
	expectEq(t, pq.Len(), 0)
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func ExamplePriorityQueue() {
	h := NewPriorityQueue[int]()
	h.Push(3)
	h.Push(2)
	h.Push(1)
	h.Push(5)
	fmt.Printf("minimum: %d\n", h.Top())

	for h.Len() > 0 {
		fmt.Printf("%d ", h.Pop())
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}
