package stl4go

import "fmt"

// DList is a doubly linked list.
type DList[T any] struct {
	head   dListNode[T]
	length int
}

type dListNode[T any] struct {
	prev, next *dListNode[T]
	value      T
}

// NewDList make a new DList object
func NewDList[T any]() *DList[T] {
	l := DList[T]{}
	l.Clear()
	return &l
}

// NewDListOf make a new DList from a serial of values
func NewDListOf[T any](vs ...T) *DList[T] {
	l := DList[T]{}
	l.Clear()
	for _, v := range vs {
		l.PushBack(v)
	}
	return &l
}

// Clear cleanup the list
func (l *DList[T]) Clear() {
	l.head.prev = &l.head
	l.head.next = &l.head
	l.length = 0
}

// Len return the length of the list
func (l *DList[T]) Len() int {
	return l.length
}

// IsEmpty return whether the list is empty
func (l *DList[T]) IsEmpty() bool {
	return l.length == 0
}

// String convert the list to string
func (l *DList[T]) String() string {
	return fmt.Sprintf("DList[%v]", nameOfType[T]())
}

type dlistIterator[T any] struct {
	dl   *DList[T]
	node *dListNode[T]
}

func (it *dlistIterator[T]) IsNotEnd() bool {
	return it.node != &it.dl.head
}

func (it *dlistIterator[T]) MoveToNext() {
	it.node = it.node.next
}

func (it *dlistIterator[T]) Value() T {
	return it.node.value
}

// Iterate returns an iterator to the first element in the list.
func (l *DList[T]) Iterate() Iterator[T] {
	return &dlistIterator[T]{l, l.head.next}
}

func (l *DList[T]) PushFront(val T) {
	n := dListNode[T]{&l.head, l.head.next, val}
	l.head.next.prev = &n
	l.head.next = &n
	l.length++
}

func (l *DList[T]) PushBack(val T) {
	n := dListNode[T]{l.head.prev, &l.head, val}
	l.head.prev.next = &n
	l.head.prev = &n
	l.length++
}

func (l *DList[T]) PopFront() (T, bool) {
	var val T
	if l.length == 0 {
		return val, false
	}
	val = l.head.next.value
	l.head.next = l.head.next.next
	l.head.prev = &l.head
	l.length--
	return val, true
}

func (l *DList[T]) PopBack() (T, bool) {
	var val T
	if l.length == 0 {
		return val, false
	}
	val = l.head.prev.value
	l.head.prev = l.head.prev.prev
	l.head.prev.next = &l.head
	l.length--
	return val, true
}

// ForEach iterate the list, apply each element to the cb callback function
func (l *DList[T]) ForEach(cb func(val T)) {
	for n := l.head.next; n != &l.head; n = n.next {
		cb(n.value)
	}
}

// ForEach iterate the list, apply each element to the cb callback function, stop if cb returns false.
func (l *DList[T]) ForEachIf(cb func(val T) bool) {
	for n := l.head.next; n != &l.head; n = n.next {
		if !cb(n.value) {
			break
		}
	}
}
