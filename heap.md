# Heap

## Easy to use

For ordered types, you can easily use the heap algorithms, for example, with `int` type:

```go
func Example() {
    heap := []int {5, 4, 3, 2, 1}
    stlgo.MakeHeap(heap)
    stl4go.PushMinHeap(&heap, 6)
    stl4go.PopMinHeap(&heap)
}
```

Please compare with [container/heap](https://pkg.go.dev/container/heap#example-package-IntHeap):

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

You must define a new type `IntHeap` and five methods before to use the library,
these boilerplate codes are verbose, tedious, and boring.

## Benchmark

The heap algorithms are also much faster than `container/heap`, even for the `Func` version.:

```console
% go test -bench BenchmarkHeap
goos: darwin
goarch: arm64
pkg: github.com/chen3feng/stl4go
BenchmarkHeapInit/container/heap.Init-10                  362450              3172 ns/op
BenchmarkHeapInit/stlgo.MakeMinHeap-10                    960280              1136 ns/op
BenchmarkHeapInit/stlgo.MakeHeapFunc-10                   483109              2336 ns/op
BenchmarkHeapPush/container/heap.Push-10                  101822             11639 ns/op
BenchmarkHeapPush/stlgo.PushMinHeap-10                    501031              2283 ns/op
BenchmarkHeapPush/stlgo.PushHeapFunc-10                   437288              2634 ns/op
BenchmarkHeapPop/container/heap.Pop-10                     14550             82170 ns/op
BenchmarkHeapPop/stlgo.PopMinHeap-10                       62139             19782 ns/op
BenchmarkHeapPop/stlgo.PopHeapFunc-10                      22693             52900 ns/op
PASS
ok      github.com/chen3feng/stl4go     17.670s
```
