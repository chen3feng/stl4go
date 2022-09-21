package stl4go

import "sync"

// Pool is a type safed sync.Pool.
type Pool[T any] sync.Pool

// MakePool returns a Pool object
func MakePool[T any]() Pool[T] {
	return Pool[T]{New: func() any { return new(T) }}
}

// MakePoolWithNew returns a Pool object with specified new function.
func MakePoolWithNew[T any](new func() *T) Pool[T] {
	if new != nil {
		return Pool[T]{New: func() any { return new() }}
	}
	return Pool[T]{}
}

// Get selects an arbitrary item from the Pool, removes it from the Pool, and returns it to the caller.
func (pool *Pool[T]) Get() *T {
	i := pool.untyped().Get()
	if i == nil {
		return nil
	}
	return i.(*T)
}

// Put puts x to the pool.
func (pool *Pool[T]) Put(x *T) {
	if x != nil {
		pool.untyped().Put(x)
	}
}

func (pool *Pool[T]) untyped() *sync.Pool {
	return (*sync.Pool)(pool)
}
