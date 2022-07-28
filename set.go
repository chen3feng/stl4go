package contalgo

import (
	"fmt"
)

// Set is an associative container that contains a unordered set of unique objects of type K.
type Set[K comparable] map[K]bool

// MakeSetOf creates a new Set object with the initial content from ks.
func MakeSetOf[K comparable](ks ...K) Set[K] {
	s := make(Set[K])
	s.AddN(ks...)
	return s
}

func (s *Set[K]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Set[K]) Len() int {
	return len(*s)
}

func (s *Set[K]) Clean() {
	for k := range *s {
		delete(*s, k)
	}
}

func (s *Set[K]) Has(k K) bool {
	_, ok := (*s)[k]
	return ok
}

func (s *Set[K]) Add(k K) {
	(*s)[k] = true
}

func (s *Set[K]) AddN(ks ...K) {
	for _, key := range ks {
		(*s)[key] = true
	}
}

func (s *Set[K]) Del(k K) {
	delete(*s, k)
}

func (s *Set[K]) DelN(ks ...K) {
	for _, k := range ks {
		delete(*s, k)
	}
}

func (s *Set[K]) Keys() []K {
	keys := make([]K, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
	return keys
}

func (s *Set[K]) ForEach(cb func(k K)) {
	for k := range *s {
		cb(k)
	}
}

func (s *Set[K]) ForEachIf(cb func(k K) bool) {
	for k := range *s {
		if !cb(k) {
			break
		}
	}
}

func (s Set[K]) String() string {
	return fmt.Sprintf("Set[%s]%v", nameOfType[K](), s.Keys())
}
