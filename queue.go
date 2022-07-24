package contalgo

import (
	"fmt"
)

// Queue is a FIFO container
type Queue[T any] struct {
	list DList[T]
}

func NewQueue[T any]() *Queue[T] {
	q := Queue[T]{}
	q.list.Clean()
	return &q
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Clean() {
	q.list.Clean()
}

func (l *Queue[T]) String() string {
	return fmt.Sprintf("Queue[%v]", nameOfType[T]())
}

func (q *Queue[T]) PushFront(val T) {
	q.list.PushFront(val)
}

func (q *Queue[T]) PushBack(val T) {
	q.list.PushBack(val)
}

func (q *Queue[T]) PopFront() (T, bool) {
	return q.list.PopFront()
}

func (q *Queue[T]) PopBack() (T, bool) {
	return q.list.PopBack()
}
