package stl4go

import (
	"fmt"
)

// BuiltinSet is an associative container that contains a unordered set of unique objects of type K.
type BuiltinSet[K comparable] map[K]bool

// MakeBuiltinSetOf creates a new BuiltinSet object with the initial content from ks.
func MakeBuiltinSetOf[K comparable](ks ...K) BuiltinSet[K] {
	s := make(BuiltinSet[K])
	s.InsertN(ks...)
	return s
}

// IsEmpty implements the Container interface.
func (s *BuiltinSet[K]) IsEmpty() bool {
	return len(*s) == 0
}

// Len implements the Container interface.
func (s *BuiltinSet[K]) Len() int {
	return len(*s)
}

// Clear implements the Container interface.
func (s *BuiltinSet[K]) Clear() {
	for k := range *s {
		delete(*s, k)
	}
}

// Has implements the Set interface.
func (s *BuiltinSet[K]) Has(k K) bool {
	_, ok := (*s)[k]
	return ok
}

// Insert implements the Set interface.
func (s *BuiltinSet[K]) Insert(k K) {
	(*s)[k] = true
}

// InsertN implements the Set interface.
func (s *BuiltinSet[K]) InsertN(ks ...K) {
	for _, key := range ks {
		(*s)[key] = true
	}
}

// Remove implements the Set interface.
func (s *BuiltinSet[K]) Remove(k K) bool {
	_, ok := (*s)[k]
	delete(*s, k)
	return ok
}

// RemoveN implements the Set interface.
func (s *BuiltinSet[K]) RemoveN(ks ...K) {
	for _, k := range ks {
		delete(*s, k)
	}
}

// Keys return a copy of all keys as a slice.
func (s *BuiltinSet[K]) Keys() []K {
	keys := make([]K, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
	return keys
}

// ForEach implements the Set interface.
func (s *BuiltinSet[K]) ForEach(cb func(k K)) {
	for k := range *s {
		cb(k)
	}
}

// ForEachIf implements the Container interface.
func (s *BuiltinSet[K]) ForEachIf(cb func(k K) bool) {
	for k := range *s {
		if !cb(k) {
			break
		}
	}
}

// String implements the fmt.Stringer interface
func (s BuiltinSet[K]) String() string {
	return fmt.Sprintf("BuiltinSet[%s]%v", nameOfType[K](), s.Keys())
}
