package stl4go

// Container is a holder object that stores a collection of other objects.
type Container interface {
	IsEmpty() bool // IsEmpty checks if the container has no elements.
	Len() int      // Len returns the number of elements in the container.
	Clear()        // Clear erases all elements from the container. After this call, Len() returns zero.
}

// Map is a associative container that contains key-value pairs with unique keys.
type Map[K any, V any] interface {
	Container
	Has(K) bool                        // Checks whether the container contains element with specific key.
	Find(K) *V                         // Finds element with specific key.
	Insert(K, V)                       // Inserts a key-value pair in to the container or replace existing value.
	Remove(K) bool                     // Remove element with specific key.
	ForEach(func(K, V))                // Iterate the container.
	ForEachIf(func(K, V) bool)         // Iterate the container, stops when the callback returns false.
	ForEachMutable(func(K, *V))        // Iterate the container, *V is mutable.
	ForEachMutableIf(func(K, *V) bool) // Iterate the container, *V is mutable, stops when the callback returns false.
}

// Set is a containers that store unique elements.
type Set[K any] interface {
	Container
	Has(K) bool             // Checks whether the container contains element with specific key.
	Insert(K)               // Inserts a key-value pair in to the container or replace existing value.
	InsertN(...K)           // Inserts multiple key-value pairs in to the container or replace existing value.
	Remove(K) bool          // Remove element with specific key.
	RemoveN(...K)           // Remove multiple elements with specific keys.
	ForEach(func(K))        // Iterate the container.
	ForEachIf(func(K) bool) // Iterate the container, stops when the callback returns false.
}

// Queue is a container that can add elements to one end and remove elements from the other end.
type Queue[T any] interface {
	Container
	Front()
	Back()
	Push(T)
	Pop() T
	TryPop() (T, bool)
}

// Deque is a container that can add and remove elements from both ends.
type Deque[T any] interface {
	Container
	Front() T
	Back() T
	PushFront(T)
	PushBack(T)
	PopFront() T
	PopBack() T
	TryPopFront() (T, bool)
	TryPopBack() (T, bool)
}
