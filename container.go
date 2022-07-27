package contalgo

// Container is a holder object that stores a collection of other objects.
type Container[T any] interface {
	IsEmpty() bool // IsEmpty checks if the container has no elements.
	Len() int      // Len returns the number of elements in the container.
	Clean()        // Clean erases all elements from the container. After this call, Len() returns zero.
}
