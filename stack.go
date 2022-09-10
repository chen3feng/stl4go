package stl4go

// Stack s is a container adaptor that provides the functionality of a stack,
// a LIFO (last-in, first-out) data structure.
type Stack[T any] struct {
	elements []T
}

// NewStack creates a new Stack object.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// NewStackCap creates a new Stack object with the specified capicity.
func NewStackCap[T any](capicity int) *Stack[T] {
	return &Stack[T]{make([]T, 0, capicity)}
}

// IsEmpty implements the Container interface.
func (s Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Len implements the Container interface.
func (s Stack[T]) Len() int {
	return len(s.elements)
}

// Cap returns the capacity of the stack.
func (s Stack[T]) Cap() int {
	return cap(s.elements)
}

// Clear implements the Container interface.
func (s *Stack[T]) Clear() {
	s.elements = s.elements[0:0]
}

// Push pushes the element to the top of the stack.
func (s *Stack[T]) Push(t T) {
	s.elements = append(s.elements, t)
}

// TryPop tries to popup an element from the top of the stack.
func (s *Stack[T]) TryPop() (val T, ok bool) {
	var t T
	if len(s.elements) == 0 {
		return t, false
	}
	t = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return t, true
}

// Pop popups an element from the top of the stack.
// It must be called when IsEmpty() returned false,
// otherwise it will panic.
func (s *Stack[T]) Pop() T {
	t := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return t
}

// Top returns the top element in the stack.
// It must be called when s.IsEmpty() returned false,
// otherwise it will panic.
func (s Stack[T]) Top() T {
	return s.elements[len(s.elements)-1]
}
