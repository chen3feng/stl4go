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
	Has(K) bool                 // Checks whether the container contains element with specific key.
	Find(K) *V                  // Finds element with specific key.
	Insert(K, V)                // Inserts a key-value pair in to the container or replace existing value.
	Remove(K) bool              // Remove element with specific key.
	ForEach(func(K, *V))        // Iterate the container.
	ForEachIf(func(K, *V) bool) // Iterate the container, stops when the callback returns false.
}
