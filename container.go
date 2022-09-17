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
	Insert(K) bool          // Inserts a element in to the container or replace existing value.
	InsertN(...K) int       // Inserts multiple elements in to the container or replace existing value.
	Remove(K) bool          // Remove specific element, return true if element was in the container.
	RemoveN(...K) int       // Remove multiple elements, return the number of removed elements.
	ForEach(func(K))        // Iterate the container.
	ForEachIf(func(K) bool) // Iterate the container, stops when the callback returns false.
}

// SortedMap is a Map that provides a total ordering on its keys.
type SortedMap[K any, V any] interface {
	Map[K, V]
	// LowerBound returns an iterator to the first element in the container that
	// does not satisfy element.key < value (i.e. greater or equal to),
	// or a end iterator if no such element is found.
	LowerBound(K) MutableMapIterator[K, V]

	// UpperBound returns an iterator to the first element in the container that
	// does not satisfy value < element.key (i.e. strictly greater),
	// or a end iterator if no such element is found.
	UpperBound(K) MutableMapIterator[K, V]

	// FindRange returns an iterator in range [first, last) (last is not included).
	FindRange(K, K) MutableMapIterator[K, V]
}

// SortedSet is a Set that provides a total ordering on its elements.
type SortedSet[K any] interface {
	Set[K]

	// LowerBound returns an iterator to the first element in the container that
	// does not satisfy element < value (i.e. greater or equal to),
	// or a end iterator if no such element is found.
	LowerBound(K) Iterator[K]

	// UpperBound returns an iterator to the first element in the container that
	// does not satisfy value < element (i.e. strictly greater),
	// or a end iterator if no such element is found.
	UpperBound(K) Iterator[K]

	// FindRange returns an iterator in range [first, last) (last is not included).
	FindRange(K, K) Iterator[K]
}

// Queue is a container that can add elements to one end and remove elements from the other end.
type Queue[T any] interface {
	Container
	Front()            // Front returns the first element in the container.
	Back()             // Back returns the last element in the container.
	Push(T)            // Push pushes an element at the back of the container.
	Pop() T            // Pop popups a front from the back of the container.
	TryPop() (T, bool) // TryPop tries to popup a element from the front of the container.
}

// Deque is a container that can add and remove elements from both ends.
type Deque[T any] interface {
	Container
	Front() T               // Front returns the first element in the container.
	Back() T                // Back returns the last element in the container.
	PushFront(T)            // PushBack pushes an element at the front of the container.
	PushBack(T)             // PushBack pushes an element at the back of the container.
	PopFront() T            // PopBack popups a front from the back of the container.
	PopBack() T             // PopBack popups a element from the back of the container.
	TryPopFront() (T, bool) // TryPopFront tries to popup a element from the front of the container.
	TryPopBack() (T, bool)  // TryPopBack tries to popup a element from the back of the container.
}
