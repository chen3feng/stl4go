# Heap

stl4go provides a group of [heap](https://en.wikipedia.org/wiki/Heap_(data_structure)) algorithms.

## Easy to use

For ordered types, you can easily use the heap algorithms, for example, with `int` type:

```go
func Example() {
    heap := []int {5, 4, 3, 2, 1}
    stl4go.MakeMinHeap(heap)
    stl4go.PushMinHeap(&heap, 6)
    n := stl4go.PopMinHeap(&heap) // get 1
}
```

Please compare it with [container/heap](https://pkg.go.dev/container/heap#example-package-IntHeap):

```go
// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```

You must define a new type `IntHeap` and five methods before to use the standard library,
these boilerplate codes are verbose, tedious, and boring.

## Benchmark

The heap algorithms are also much faster than `container/heap`, even for the `Func` version.:

```console
% go test -bench BenchmarkHeap -benchmem
goos: darwin
goarch: arm64
pkg: github.com/chen3feng/stl4go
BenchmarkHeapInit/container/heap.Init-10                  360126              3195 ns/op               0 B/op          0 allocs/op
BenchmarkHeapInit/stlgo.MakeMinHeap-10                    998325              1127 ns/op               0 B/op          0 allocs/op
BenchmarkHeapInit/stlgo.MakeHeapFunc-10                   488419              2355 ns/op               0 B/op          0 allocs/op
...
BenchmarkHeapPush/container/heap.Push-10                  101260             11630 ns/op            5952 B/op        744 allocs/op
BenchmarkHeapPush/stlgo.PushMinHeap-10                    511680              2261 ns/op               0 B/op          0 allocs/op
BenchmarkHeapPush/stlgo.PushHeapFunc-10                   445850              2625 ns/op               0 B/op          0 allocs/op
...
BenchmarkHeapPop/container/heap.Pop-10                     14709             81153 ns/op            5952 B/op        744 allocs/op
BenchmarkHeapPop/stlgo.PopMinHeap-10                       61608             19516 ns/op               0 B/op          0 allocs/op
BenchmarkHeapPop/stlgo.PopHeapFunc-10                      22887             52440 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/chen3feng/stl4go     19.827s
```
