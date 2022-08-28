package stl4go

// PriorityQueue is an queue with priority.
// The elements of the priority queue are ordered according to their natural ordering,
// or by a less function provided at construction time, depending on which constructor is used.
type PriorityQueue[T any] struct {
	heap []T
	impl pqImpl[T]
}

// NewPriorityQueue creates an empty priority object.
func NewPriorityQueue[T Ordered]() *PriorityQueue[T] {
	pq := pqOrdered[T]{}
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

// NewPriorityQueueOn creates a new priority object on the specified slices.
// The slice become a heap after the call.
func NewPriorityQueueOn[T Ordered](slice []T) *PriorityQueue[T] {
	MakeMinHeap(slice)
	pq := pqOrdered[T]{}
	pq.heap = slice
	pq.impl = pqImpl[T](&pq)
	return &pq.PriorityQueue
}

// NewPriorityQueueOf creates a new priority object with specified initial elements.
func NewPriorityQueueOf[T Ordered](elements ...T) *PriorityQueue[T] {
	return NewPriorityQueueOn(elements)
}

// NewPriorityQueueFunc creates an empty priority object with specified compare function less.
func NewPriorityQueueFunc[T any](less LessFn[T]) *PriorityQueue[T] {
	pq := pqFunc[T]{}
	pq.less = less
	pq.impl = (pqImpl[T])(&pq)
	return &pq.PriorityQueue
}

// Len returns the number of elements in the priority queue.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.heap)
}

// IsEmpty checks whether priority queue has no elements.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

// Clear checks whether priority queue has no elements.
func (pq *PriorityQueue[T]) Clear() {
	pq.heap = pq.heap[0:0]
}

// Top returns the top element in the priority queue.
func (pq *PriorityQueue[T]) Top() T {
	return pq.heap[0]
}

// Push pushes the given element v to the priority queue.
func (pq *PriorityQueue[T]) Push(v T) {
	pq.impl.Push(v)
}

// Pop removes the top element in the priority queue.
func (pq *PriorityQueue[T]) Pop() T {
	return pq.impl.Pop()
}

type pqImpl[T any] interface {
	Push(v T)
	Pop() T
}

type pqOrdered[T Ordered] struct {
	PriorityQueue[T]
}

func (pq *pqOrdered[T]) Push(v T) {
	PushMinHeap(&pq.heap, v)
}

func (pq *pqOrdered[T]) Pop() T {
	return PopMinHeap(&pq.heap)
}

// funcHeap is a min-heap of T compared with less.
type pqFunc[T any] struct {
	PriorityQueue[T]
	less LessFn[T]
}

func (pq *pqFunc[T]) Push(v T) {
	PushHeapFunc(&pq.heap, v, pq.less)
}

func (pq *pqFunc[T]) Pop() T {
	return PopHeapFunc(&pq.heap, pq.less)
}
