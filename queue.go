package stl4go

import (
	"fmt"
)

// Queue is a FIFO container
type Queue[T any] struct {
	list DList[T]
}

// NewQueue create a new Queue object.
func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	q.list.Clear()
	return &q
}

// Len implements the Container interface.
func (q *Queue[T]) Len() int {
	return q.list.Len()
}

// IsEmpty implements the Container interface.
func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Clear implements the Container interface.
func (q *Queue[T]) Clear() {
	q.list.Clear()
}

// Len implements the fmt.Stringer interface.
func (q *Queue[T]) String() string {
	return fmt.Sprintf("Queue[%v]", nameOfType[T]())
}

// PushFront pushed an element to the front of the queue.
func (q *Queue[T]) PushFront(val T) {
	q.list.PushFront(val)
}

// PushBack pushed an element to the back of the queue.
func (q *Queue[T]) PushBack(val T) {
	q.list.PushBack(val)
}

// PopFront popups an element from the front of the queue.
func (q *Queue[T]) PopFront() T {
	return q.list.PopFront()
}

// PopBack popups an element from the back of the queue.
func (q *Queue[T]) PopBack() T {
	return q.list.PopBack()
}

// TryPopFront tries popuping an element from the front of the queue.
func (q *Queue[T]) TryPopFront() (T, bool) {
	return q.list.TryPopFront()
}

// TryPopBack tries popuping an element from the back of the queue.
func (q *Queue[T]) TryPopBack() (T, bool) {
	return q.list.TryPopBack()
}
