package contalgo

import "fmt"

// DList is a doubly linked list.
type DList[T any] struct {
	head   dListNode[T]
	length int
}

type dListNode[T any] struct {
	prev, next *dListNode[T]
	val        T
}

// NewDList make a new DList object
func NewDList[T any]() *DList[T] {
	l := DList[T]{}
	l.Clean()
	return &l
}

// NewDListOf make a new DList from a serial of values
func NewDListOf[T any](vs ...T) *DList[T] {
	l := DList[T]{}
	l.Clean()
	for _, v := range vs {
		l.PushBack(v)
	}
	return &l
}

// Clean cleanup the list
func (l *DList[T]) Clean() {
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
	val = l.head.next.val
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
	val = l.head.prev.val
	l.head.prev = l.head.prev.prev
	l.head.prev.next = &l.head
	l.length--
	return val, true
}

// ForEach iterate the list, apply each element to the cb callback function
func (l *DList[T]) ForEach(cb func(val T)) {
	for n := l.head.next; n != &l.head; n = n.next {
		cb(n.val)
	}
}

// ForEach iterate the list, apply each element to the cb callback function, stop if cb returns false.
func (l *DList[T]) ForEachIf(cb func(val T) bool) {
	for n := l.head.next; n != &l.head; n = n.next {
		if !cb(n.val) {
			break
		}
	}
}
