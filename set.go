package contalgo

import (
	"fmt"
	"reflect"
)

// Set is an associative container that contains a unordered set of unique objects of type K.
type Set[K comparable] struct {
	m map[K]struct{} // Empty struct takes 0 byte space
}

// NewSet creates a new Set object.
func NewSet[K comparable]() *Set[K] {
	return &Set[K]{m: make(map[K]struct{})}
}

// NewSetOf creates a new Set object with the initial content from ks.
func NewSetOf[K comparable](ks ...K) *Set[K] {
	s := NewSet[K]()
	s.AddN(ks...)
	return s
}

func (s *Set[K]) IsEmpty() bool {
	return len(s.m) == 0
}

func (s *Set[K]) Len() int {
	return len(s.m)
}

func (s *Set[K]) Clean() {
	for k := range s.m {
		delete(s.m, k)
	}
}

func (s *Set[K]) Has(k K) bool {
	_, ok := s.m[k]
	return ok
}

func (s *Set[K]) Add(k K) {
	s.m[k] = struct{}{}
}

func (s *Set[K]) AddN(ks ...K) {
	for _, key := range ks {
		s.m[key] = struct{}{}
	}
}

func (s *Set[K]) Del(k K) {
	delete(s.m, k)
}

func (s *Set[K]) DelN(ks ...K) {
	for _, k := range ks {
		delete(s.m, k)
	}
}

func (s *Set[K]) Keys() []K {
	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set[K]) ForEach(cb func(k K)) {
	for k := range s.m {
		cb(k)
	}
}

func (s *Set[K]) ForEachIf(cb func(k K) bool) {
	for k := range s.m {
		if !cb(k) {
			break
		}
	}
}

func (s *Set[K]) String() string {
	var k *K = nil
	return fmt.Sprintf("Set[%s]%v", reflect.TypeOf(k).String()[1:], s.Keys())
}
