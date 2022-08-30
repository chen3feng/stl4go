package stl4go

// Iterator is the interface for container's iterator.
type Iterator[T any] interface {
	IsNotEnd() bool // Whether it is point to the end of the range.
	MoveToNext()    // Let it point to the next element.
	Value() T       // Return the value of current element.
}

// MutableIterator is the interface for container's mutable iterator.
type MutableIterator[T any] interface {
	Iterator[T]
	Pointer() *T // Return the pointer to the value of current element.
}

// MapIterator is the interface for map's iterator.
type MapIterator[K any, V any] interface {
	Iterator[V]
	Key() K // The key of the element
}

// MutableMapIterator is the interface for map's mutable iterator.
type MutableMapIterator[K any, V any] interface {
	MutableIterator[V]
	Key() K // The key of the element
}
