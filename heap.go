// Fast generic heap algorithms, faster than [container/heap]

package stl4go

// MakeMinHeap build a min-heap on slice array.
//
// Complexity: O(len(array))
func MakeMinHeap[T Ordered](array []T) {
	// heapify
	n := len(array)
	for i := n/2 - 1; i >= 0; i-- {
		heapDown(array, i, n)
	}
}

// IsMinHeap checks whether the elements in slice array are a min heap.
//
// Complexity: O(len(array)).
func IsMinHeap[T Ordered](array []T) bool {
	parent := 0
	for child := 1; child < len(array); child++ {
		if array[parent] > array[child] {
			return false
		}

		if (child & 1) == 0 {
			parent++
		}
	}
	return true
}

// PushMinHeap pushes a element v into the min heap.
//
// Complexity: O(log(len(*heap))).
func PushMinHeap[T Ordered](heap *[]T, v T) {
	*heap = append(*heap, v)
	heapUp(*heap, len(*heap)-1)
}

// PopMinHeap removes and returns the minimum element from the heap.
//
// Complexity: O(log n) where n = len(*heap).
func PopMinHeap[T Ordered](heap *[]T) T {
	h := *heap
	n := len(h) - 1
	heapSwap(h, 0, n)
	heapDown(h, 0, n)
	*heap = h[0:n]
	return h[n]
}

// RemoveMinHeap removes and returns the element at index i from the min heap.
//
// Complexity: is O(log(n)) where n = len(*heap).
func RemoveMinHeap[T Ordered](heap *[]T, i int) T {
	h := *heap
	n := len(h) - 1
	if n != i {
		heapSwap(h, i, n)
		if !heapDown(h, i, n) {
			heapUp(h, i)
		}
	}
	*heap = h[0:n]
	return h[n]
}

func heapSwap[T any](heap []T, i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func heapUp[T Ordered](heap []T, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(heap[j] < heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		j = i
	}
}

func heapDown[T Ordered](heap []T, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && heap[j2] < heap[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		if !(heap[j] < heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		i = j
	}
	return i > i0
}

// MakeHeapFunc build a min-heap on slice array with compare function less.
//
// Complexity: O(len(array))
func MakeHeapFunc[T any](array []T, less LessFn[T]) {
	// heapify
	n := len(array)
	for i := n/2 - 1; i >= 0; i-- {
		heapDownFunc(array, i, n, less)
	}
}

// IsHeapFunc checks whether the elements in slice array are a min heap (accord to less).
//
// Complexity: O(len(array)).
func IsHeapFunc[T any](array []T, less LessFn[T]) bool {
	parent := 0
	for child := 1; child < len(array); child++ {
		if !less(array[parent], array[child]) {
			return false
		}

		if (child & 1) == 0 {
			parent++
		}

	}
	return true
}

// PushHeapFunc pushes a element v into the heap.
//
// Complexity: O(log(len(*heap))).
func PushHeapFunc[T any](heap *[]T, v T, less LessFn[T]) {
	*heap = append(*heap, v)
	heapUpFunc(*heap, len(*heap)-1, less)
}

// PopHeapFunc removes and returns the minimum (according to less) element from the heap.
//
// Complexity: O(log n) where n = len(*heap).
func PopHeapFunc[T any](heap *[]T, less LessFn[T]) T {
	h := *heap
	n := len(h) - 1
	heapSwap(h, 0, n)
	heapDownFunc(h, 0, n, less)
	*heap = h[0:n]
	return h[n]
}

// RemoveHeapFunc removes and returns the element at index i from the heap.
//
// Complexity: is O(log(n)) where n = len(*heap).
func RemoveHeapFunc[T any](heap *[]T, i int, less LessFn[T]) T {
	h := *heap
	n := len(h) - 1
	if n != i {
		heapSwap(h, i, n)
		if !heapDownFunc(h, i, n, less) {
			heapUpFunc(h, i, less)
		}
	}
	*heap = h[0:n]
	return h[n]
}

func heapUpFunc[T any](heap []T, j int, less LessFn[T]) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !less(heap[j], heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		j = i
	}
}

func heapDownFunc[T any](heap []T, i0, n int, less LessFn[T]) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && less(heap[j2], heap[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !less(heap[j], heap[i]) {
			break
		}
		heapSwap(heap, i, j)
		i = j
	}
	return i > i0
}
