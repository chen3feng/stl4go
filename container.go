package contalgo

import "golang.org/x/exp/constraints"

type Container[T any] interface {
	Len() int
	IsEmpty() bool
	Clean()
	ForEach(cb func(T))
}

type MutableContainer[T any] interface {
	Container[T]
}

func container_Max[T constraints.Ordered](c Container[T]) T {
	var max T
	c.ForEach(func(v T) {
		if v > max {
			max = v
		}
	})
	return max
}
