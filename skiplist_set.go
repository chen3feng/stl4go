package stl4go

// SkipListSet is a SortedSet implemented with skiplist.
type SkipListSet[K any] SkipList[K, struct{}]

// NewSkipListSet creates a new SkipListSet object.
func NewSkipListSet[K Ordered]() *SkipListSet[K] {
	return (*SkipListSet[K])(NewSkipList[K, struct{}]())
}

// NewSkipListSetFunc creates a new SkipListSet object with specified compare function.
func NewSkipListSetFunc[K any](cmp CompareFn[K]) *SkipListSet[K] {
	return (*SkipListSet[K])(NewSkipListFunc[K, struct{}](cmp))
}

// NewSkipListSetOf creates a new SkipListSet object with specified elements.
func NewSkipListSetOf[K Ordered](elements ...K) *SkipListSet[K] {
	s := NewSkipListSet[K]()
	for i := range elements {
		s.Insert(elements[i])
	}
	return s
}

// IsEmpty implements the SortedSet interface.
func (s *SkipListSet[K]) IsEmpty() bool {
	return s.asMap().IsEmpty()
}

// Len implements the SortedSet interface.
func (s *SkipListSet[K]) Len() int {
	return s.asMap().Len()
}

// Clear implements the SortedSet interface.
func (s *SkipListSet[K]) Clear() {
	s.asMap().Clear()
}

// Has implements the SortedSet interface.
func (s *SkipListSet[K]) Has(key K) bool {
	return s.asMap().Has(key)
}

// Insert implements the SortedSet interface.
func (s *SkipListSet[K]) Insert(key K) bool {
	oldLen := s.Len()
	s.asMap().Insert(key, struct{}{})
	return s.Len() > oldLen
}

// InsertN implements the SortedSet interface.
func (s *SkipListSet[K]) InsertN(keys ...K) int {
	oldLen := s.Len()
	for i := range keys {
		s.Insert(keys[i])
	}
	return s.Len() - oldLen
}

// Remove implements the SortedSet interface.
func (s *SkipListSet[K]) Remove(key K) bool {
	return s.asMap().Remove(key)
}

// RemoveN implements the SortedSet interface.
func (s *SkipListSet[K]) RemoveN(keys ...K) int {
	oldLen := s.Len()
	for i := range keys {
		s.Remove(keys[i])
	}
	return oldLen - s.Len()
}

// Keys return a copy of sorted keys as slice.
func (s *SkipListSet[K]) Keys() []K {
	keys := make([]K, 0, s.Len())
	s.ForEach(func(k K) { keys = append(keys, k) })
	return keys
}

// ForEach implements the SortedSet interface.
func (s *SkipListSet[K]) ForEach(f func(K)) {
	s.asMap().ForEach(func(k K, s struct{}) { f(k) })
}

// ForEachIf implements the SortedSet interface.
func (s *SkipListSet[K]) ForEachIf(f func(K) bool) {
	s.asMap().ForEachIf(func(k K, s struct{}) bool { return f(k) })
}

// LowerBound implements the SortedSet interface.
func (s *SkipListSet[K]) LowerBound(key K) Iterator[K] {
	return &skipListSetIterator[K]{s.asMap().impl.lowerBound(key), nil}
}

// UpperBound implements the SortedSet interface.
func (s *SkipListSet[K]) UpperBound(key K) Iterator[K] {
	return &skipListSetIterator[K]{s.asMap().impl.upperBound(key), nil}
}

// FindRange implements the SortedSet interface.
func (s *SkipListSet[K]) FindRange(first, last K) Iterator[K] {
	m := s.asMap().impl
	return &skipListSetIterator[K]{m.lowerBound(first), m.upperBound(last)}
}

func (s *SkipListSet[K]) asMap() *SkipList[K, struct{}] {
	return (*SkipList[K, struct{}])(s)
}

type skipListSetIterator[K any] skipListIterator[K, struct{}]

func (it *skipListSetIterator[K]) IsNotEnd() bool {
	return it.node != it.end
}

func (it *skipListSetIterator[K]) MoveToNext() {
	it.node = it.node.next[0]
}

func (it *skipListSetIterator[K]) Value() K {
	return it.node.key
}
