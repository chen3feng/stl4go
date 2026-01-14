package stl4go

import (
	"fmt"
)

// BuiltinSet is an associative container that contains an unordered set of unique objects of type K.
type BuiltinSet[K comparable] map[K]struct{}

var _ Set[int] = (*BuiltinSet[int])(nil)

// SetOf creates a new BuiltinSet object with the initial content from ks.
func SetOf[K comparable](ks ...K) BuiltinSet[K] {
	s := make(BuiltinSet[K])
	s.InsertN(ks...)
	return s
}

// IsEmpty implements the Container interface.
func (s BuiltinSet[K]) IsEmpty() bool {
	return len(s) == 0
}

// Len implements the Container interface.
func (s BuiltinSet[K]) Len() int {
	return len(s)
}

// Clear implements the Container interface.
func (s BuiltinSet[K]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Has implements the Set interface.
func (s BuiltinSet[K]) Has(k K) bool {
	_, ok := s[k]
	return ok
}

// Insert implements the Set interface.
func (s BuiltinSet[K]) Insert(k K) bool {
	oldLen := len(s)
	s[k] = struct{}{}
	return len(s) > oldLen
}

// InsertN implements the Set interface.
func (s BuiltinSet[K]) InsertN(ks ...K) int {
	oldLen := len(s)
	for _, key := range ks {
		s[key] = struct{}{}
	}
	return len(s) - oldLen
}

// Remove implements the Set interface.
func (s BuiltinSet[K]) Remove(k K) bool {
	_, ok := s[k]
	delete(s, k)
	return ok
}

// Delete deletes an element from the set.
// It returns nothing, so it's faster than Remove.
func (s BuiltinSet[K]) Delete(k K) {
	delete(s, k)
}

// RemoveN implements the Set interface.
func (s BuiltinSet[K]) RemoveN(ks ...K) int {
	oldLen := len(s)
	for _, k := range ks {
		delete(s, k)
	}
	return oldLen - len(s)
}

// Keys return a copy of all keys as a slice.
func (s BuiltinSet[K]) Keys() []K {
	keys := make([]K, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

// ForEach implements the Set interface.
func (s BuiltinSet[K]) ForEach(cb func(k K)) {
	for k := range s {
		cb(k)
	}
}

// ForEachIf implements the Container interface.
func (s BuiltinSet[K]) ForEachIf(cb func(k K) bool) {
	for k := range s {
		if !cb(k) {
			break
		}
	}
}

// String implements the fmt.Stringer interface.
func (s BuiltinSet[K]) String() string {
	return fmt.Sprintf("BuiltinSet[%s]%v", nameOfType[K](), s.Keys())
}

// Update adds all elements from other to set. set |= other.
func (s BuiltinSet[K]) Update(other BuiltinSet[K]) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Union returns a new set with elements from the set and other.
func (s BuiltinSet[K]) Union(other BuiltinSet[K]) BuiltinSet[K] {
	result := BuiltinSet[K]{}
	result.Update(s)
	result.Update(other)
	return result
}

func orderSet[K comparable](a, b BuiltinSet[K]) (small, large BuiltinSet[K]) {
	if a.Len() < b.Len() {
		return a, b
	}
	return b, a
}

// Intersection returns a new set with elements common to the set and other.
func (s BuiltinSet[K]) Intersection(other BuiltinSet[K]) BuiltinSet[K] {
	result := BuiltinSet[K]{}
	small, large := orderSet(s, other)
	for k := range small {
		if large.Has(k) {
			result.Insert(k)
		}
	}
	return result
}

// Difference returns a new set with elements in the set that are not in other.
func (s BuiltinSet[K]) Difference(other BuiltinSet[K]) BuiltinSet[K] {
	result := BuiltinSet[K]{}
	for k := range s {
		if !other.Has(k) {
			result.Insert(k)
		}
	}
	return result
}

// IsDisjointOf return True if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (s BuiltinSet[K]) IsDisjointOf(other BuiltinSet[K]) bool {
	small, large := orderSet(s, other)
	for k := range small {
		if large.Has(k) {
			return false
		}
	}
	return true
}

// IsSubsetOf tests whether every element in the set is in other.
func (s BuiltinSet[K]) IsSubsetOf(other BuiltinSet[K]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for k := range s {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

// IsSupersetOf tests whether every element in other is in the set.
func (s BuiltinSet[K]) IsSupersetOf(other BuiltinSet[K]) bool {
	return other.IsSubsetOf(s)
}
