package stl4go

import "golang.org/x/exp/constraints"

//Comparable only supports operator ==  !=
// constraints.Ordered supports the operators < <= >= >

type BST[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

type Node[T constraints.Ordered] struct {
	left, right *Node[T]
	value       T
}

// Len return the length of the list.
func (b *BST[T]) Len() int {
	return b.size
}

func (b *BST[T]) IsEmpty() bool {
	return b.size == 0
}

func isLess[T constraints.Ordered](a, b T) bool {
	return a < b
}

func isBig[T constraints.Ordered](a, b T) bool {
	return a > b
}
func isEqual[T constraints.Ordered](a, b T) bool {
	return a == b
}

func (b *BST[T]) add(v T) {
	b.root = b._add(b.root, v)
	return
}

func (b *BST[T]) _add(node *Node[T], v T) *Node[T] {
	if node == nil {
		b.size++
		var root *Node[T]
		root.value = v
		return root
	}
	if isLess(v, node.value) {
		node.left = b._add(node.left, v)
	} else if isBig(v, node.value) {
		node.right = b._add(node.right, v)
	}
	return node
}

func (b *BST[T]) Contains(v T) bool {
	return b._Contains(b.root, v)
}

func (b *BST[T]) _Contains(node *Node[T], v T) bool {
	if node == nil {
		return false
	}
	if isEqual(v, node.value) {
		return true
	} else if isLess(v, node.value) {
		return b._Contains(node.left, v)
	} else {
		return b._Contains(node.right, v)
	}
}

func (b *BST[T]) minimum() T {
	if b.size == 0 {
		var i T
		return i
	}
	return b._minimum(b.root).value
}

func (b *BST[T]) _minimum(node *Node[T]) *Node[T] {
	if node.left == nil {
		return node
	}
	return b._minimum(node.left)
}

func (b *BST[T]) maximum() T {
	if b.size == 0 {
		var i T
		return i
	}
	return b._maximum(b.root).value
}

func (b *BST[T]) _maximum(node *Node[T]) *Node[T] {
	if node.left == nil {
		return node
	}
	return b._maximum(node.right)
}
