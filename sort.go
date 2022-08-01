package contalgo

import (
	"sort"
)

// IsSorted returns whether the slice a is sorted in ascending order.
//
// Complexity: O(len(a)).
func IsSorted[T Ordered](a []T) bool {
	if len(a) == 0 {
		return true
	}
	prev := a[0]
	for _, v := range a[1:] {
		if v < prev {
			return false
		}
		prev = v
	}
	return true
}

// IsDescSorted returns whether the slice a is sorted in descending order.
//
// Complexity: O(len(a)).
func IsDescSorted[T Ordered](a []T) bool {
	if len(a) == 0 {
		return true
	}
	prev := a[0]
	for _, v := range a[1:] {
		if v > prev {
			return false
		}
		prev = v
	}
	return true
}

type ascSlice[T Ordered] []T

func (x ascSlice[T]) Len() int           { return len(x) }
func (x ascSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x ascSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort sorts data in ascending order.
// The order of equal elements is not guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func Sort[T Ordered](a []T) {
	sort.Sort(ascSlice[T](a))
}

// StableSort sorts data in ascending order stably.
// The order of equivalent elements is guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func StableSort[T Ordered](a []T) {
	sort.Stable(ascSlice[T](a))
}

type descSlice[T Ordered] []T

func (x descSlice[T]) Len() int           { return len(x) }
func (x descSlice[T]) Less(i, j int) bool { return x[i] > x[j] }
func (x descSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// DescSort sorts data in descending order.
// The order of equal elements is not guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func DescSort[T Ordered](a []T) {
	sort.Sort(descSlice[T](a))
}

// DescStableSort sorts data in descending order stably.
// The order of equivalent elements is guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func DescStableSort[T Ordered](a []T) {
	sort.Stable(descSlice[T](a))
}

type funcSortable[T any] struct {
	e    []T
	less LessFn[T]
}

func (x funcSortable[T]) Len() int           { return len(x.e) }
func (x funcSortable[T]) Less(i, j int) bool { return x.less(x.e[i], x.e[j]) }
func (x funcSortable[T]) Swap(i, j int)      { x.e[i], x.e[j] = x.e[j], x.e[i] }

// SortFunc sorts data in ascending order with compare func less.
// The order of equal elements is not guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func SortFunc[T any](a []T, less func(x, y T) bool) {
	sort.Sort(funcSortable[T]{a, less})
}

// StableSortFunc sorts data in ascending order with compare func less stably.
// The order of equivalent elements is guaranteed to be preserved.
//
// Complexity: O(N*log(N)), N=len(a).
func StableSortFunc[T any](a []T, less func(x, y T) bool) {
	sort.Stable(funcSortable[T]{a, less})
}
