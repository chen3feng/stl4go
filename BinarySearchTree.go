package stl4go

type BST[T any] struct {
	node   *Node[T]
	root   *Node[T]
	length int
}
type Node[T any] struct {
	left, right *Node[T]
	value       T
}
