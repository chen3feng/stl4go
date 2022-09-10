package stl4go

import "errors"

// SList is a singly linked list.
type SList[T any] struct {
	head   *sListNode[T]
	tail   *sListNode[T] // To support Back and PushBack
	length int
}

type sListNode[T any] struct {
	next  *sListNode[T]
	value T
}

// SListOf return a SList that contains values.
func SListOf[T any](values ...T) SList[T] {
	l := SList[T]{}
	for i := range values {
		l.PushBack(values[i])
	}
	return l
}

// IsEmpty checks if the container has no elements.
func (l *SList[T]) IsEmpty() bool {
	return l.length == 0
}

// Len returns the number of elements in the container.
func (l *SList[T]) Len() int {
	return l.length
}

// Clear erases all elements from the container. After this call, Len() returns zero.
func (l *SList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Front returns the first element in the list.
func (l *SList[T]) Front() T {
	return l.head.value
}

// Back returns the last element in the list.
func (l *SList[T]) Back() T {
	return l.tail.value
}

// PushFront pushed an element to the front of the list.
func (l *SList[T]) PushFront(v T) {
	node := sListNode[T]{l.head, v}
	l.head = &node
	if l.tail == nil {
		l.tail = &node
	}
	l.length++
}

// PushBack pushed an element to the tail of the list.
func (l *SList[T]) PushBack(v T) {
	node := sListNode[T]{nil, v}
	if l.tail != nil {
		l.tail.next = &node
	}
	l.tail = &node
	if l.head == nil {
		l.head = &node
	}
	l.length++
}

// PopFront popups an element from the front of the list.
// The list must be non-empty!
func (l *SList[T]) PopFront() T {
	node := l.head
	if node != nil {
		l.head = node.next
		if l.head == nil {
			l.tail = nil
		}
		l.length--
	}
	return node.value
}

// Reverse reverses the order of all elements in the container.
func (l *SList[T]) Reverse() {
	var head, tail *sListNode[T]
	for node := l.head; node != nil; {
		next := node.next
		node.next = head
		head = node
		if tail == nil {
			tail = node
		}
		node = next
	}
	l.head = head
	l.tail = tail
}

// Values copies all elements in the container to a slice and return it.
func (l *SList[T]) Values() []T {
	s := make([]T, l.Len())
	for node, i := l.head, 0; node != nil; node = node.next {
		s[i] = node.value
		i++
	}
	return s
}

func (l *SList[T]) Remove(index int) error {
	if index < 0 || index >= l.length {
		return errors.New("index out of range")
	}
	var i T
	dummy := &sListNode[T]{l.head, i}
	for i := 0; i < index; i++ {
		dummy = dummy.next
	}
	//var retNode *sListNode[T]
	retNode := dummy.next
	dummy.next = retNode.next
	l.length--
	return nil
}

// InsertAfter inserts an element after the iterator into the list,
// return an iterator to the inserted element.
// func (l *SList[T]) InsertAfter(it Iterator[T], value T) MutableIterator[T] {
// 	// cause internal compiler error: panic: runtime error: invalid memory address or nil pointer dereference
// 	itp := it.(*sListIterator[T])
// 	node := itp.node
// 	newNode := sListNode[T]{node.next, value}
// 	node.next = &newNode
// 	return &sListIterator[T]{&newNode}
// }

// ForEach iterate the list, apply each element to the cb callback function.
func (l *SList[T]) ForEach(cb func(T)) {
	for node := l.head; node != nil; node = node.next {
		cb(node.value)
	}
}

// ForEachIf iterate the container, apply each element to the cb callback function,
// stop if cb returns false.
func (l *SList[T]) ForEachIf(cb func(T) bool) {
	for node := l.head; node != nil; node = node.next {
		if !cb(node.value) {
			break
		}
	}
}

// ForEachMutable iterate the container, apply pointer of each element to the cb callback function.
func (l *SList[T]) ForEachMutable(cb func(*T)) {
	for node := l.head; node != nil; node = node.next {
		cb(&node.value)
	}
}

// ForEachMutableIf iterate the container, apply pointer of each element to the cb callback function,
// stop if cb returns false.
func (l *SList[T]) ForEachMutableIf(cb func(*T) bool) {
	for node := l.head; node != nil; node = node.next {
		if !cb(&node.value) {
			break
		}
	}
}

// Iterate returns an iterator to the whole container.
func (l *SList[T]) Iterate() MutableIterator[T] {
	it := sListIterator[T]{l.head}
	return &it
}

type sListIterator[T any] struct {
	node *sListNode[T]
}

func (it sListIterator[T]) IsNotEnd() bool {
	return it.node != nil
}

func (it *sListIterator[T]) MoveToNext() {
	it.node = it.node.next
}

func (it sListIterator[T]) Value() T {
	return it.node.value
}

func (it sListIterator[T]) Pointer() *T {
	return &it.node.value
}

// TODO: Sort
