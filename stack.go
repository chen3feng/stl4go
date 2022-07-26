package contalgo

// Stack s is a container adaptor that provides the functionality of a stack,
// a LIFO (last-in, first-out) data structure.
type Stack[T any] struct {
	ele []T
}

// NewStack creates a new Stack object.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// NewStackCap creates a new Stack object with the specified capicity.
func NewStackCap[T any](capicity int) *Stack[T] {
	return &Stack[T]{make([]T, 0, capicity)}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.ele) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.ele)
}

func (s *Stack[T]) Cap() int {
	return cap(s.ele)
}

func (s *Stack[T]) Clean() {
	s.ele = s.ele[0:0]
}

func (s *Stack[T]) Push(t T) {
	s.ele = append(s.ele, t)
}

func (s *Stack[T]) Pop() (val T, ok bool) {
	var t T
	if len(s.ele) == 0 {
		return t, false
	}
	t = s.ele[len(s.ele)-1]
	s.ele = s.ele[:len(s.ele)-1]
	return t, true
}

func (s *Stack[T]) MustPop() T {
	t := s.ele[len(s.ele)-1]
	s.ele = s.ele[:len(s.ele)-1]
	return t
}
