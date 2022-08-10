// This implementation is based on https://github.com/liyue201/gostl/tree/master/ds/skiplist
// (many thanks), added many optimizations, such as:
//
//  - adaptive level
//  - lesser search for prevs when key already exists.
//  - reduce memory allocations
//  - richer interface.
//
// etc.

package stl4go

import (
	"math/rand"
	"time"
)

const (
	skipListMaxLevel = 40
)

// SkipList is a probabilistic data structure that seem likely to supplant balanced trees as the
// implementation method of choice for many applications. Skip list algorithms have the same
// asymptotic expected time bounds as balanced trees and are simpler, faster and use less space.
//
// See https://en.wikipedia.org/wiki/Skip_list for more details.
type SkipList[K any, V any] struct {
	keyCmp         CompareFn[K]
	builtinCompare bool // Whether the key type is Ordered, so we can use builtin compare to improve performance.
	level          int  // Current level, may increase dynamically during insertion
	len            int  // Total elements numner in the skiplist.
	head           skipListNode[K, V]
	// This cache is used to save the previous nodes when modifying the skip list to avoid
	// allocating memory each time it is called.
	prevsCache []*skipListNode[K, V] // Cache to avoid memory allocation.
	rander     *rand.Rand
}

type skipListNode[K any, V any] struct {
	key   K
	value V
	next  []*skipListNode[K, V]
}

//go:generate bash ./skiplist_newnode_generate.sh skipListMaxLevel skiplist_newnode.go
// func newSkipListNode[K Ordered, V any](level int, key K, value V) *skipListNode[K, V]

// NewSkipList creates a new Skiplist.
func NewSkipList[K Ordered, V any]() *SkipList[K, V] {
	sl := NewSkipListFunc[K, V](OrderedCompare[K])
	sl.builtinCompare = true
	return sl
}

// NewSkipListFromMap create a new Skiplist from a map.
func NewSkipListFromMap[K Ordered, V any](m map[K]V) *SkipList[K, V] {
	sl := NewSkipList[K, V]()
	for k, v := range m {
		sl.Insert(k, v)
	}
	return sl
}

// NewSkipListFunc creates a new Skiplist with specified compare function keyCmp.
func NewSkipListFunc[K any, V any](keyCmp CompareFn[K]) *SkipList[K, V] {
	l := &SkipList[K, V]{
		level:  1,
		keyCmp: keyCmp,
		// #nosec G404 -- This is not a security condition
		rander:     rand.New(rand.NewSource(time.Now().Unix())),
		prevsCache: make([]*skipListNode[K, V], skipListMaxLevel),
	}
	l.head.next = make([]*skipListNode[K, V], skipListMaxLevel)
	return l
}

func (sl *SkipList[K, V]) IsEmpty() bool {
	return sl.len == 0
}

func (sl *SkipList[K, V]) Len() int {
	return sl.len
}

func (sl *SkipList[K, V]) Clear() {
	for i := range sl.head.next {
		sl.head.next[i] = nil
	}
	sl.level = 1
	sl.len = 0
}

// Insert inserts a key-value pair into the skiplist.
// If the key is already in the skip list, it's value will be updated.
func (sl *SkipList[K, V]) Insert(key K, value V) {
	node, prevs := sl.findInsertPoint(key)

	if node != nil {
		// Already exist, update the value
		node.value = value
		return
	}

	level := sl.randomLevel()
	node = newSkipListNode(level, key, value)

	for i := 0; i < Min(level, sl.level); i++ {
		node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = node
	}

	if level > sl.level {
		for i := sl.level; i < level; i++ {
			sl.head.next[i] = node
		}
		sl.level = level
	}

	sl.len++
}

// Find returns the value associated with the passed key if the key is in the skiplist, otherwise
// returns nil.
func (sl *SkipList[K, V]) Find(key K) *V {
	node := sl.findNode(key)
	if node != nil {
		return &node.value
	}
	return nil
}

func (sl *SkipList[K, V]) Has(key K) bool {
	return sl.findNode(key) != nil
}

// Remove removes the key-value pair associated with the passed key and returns true if the key is
// in the skiplist, otherwise returns false.
func (sl *SkipList[K, V]) Remove(key K) bool {
	node, prevs := sl.findRemovePoint(key)
	if node == nil {
		return false
	}
	for i, v := range node.next {
		prevs[i].next[i] = v
	}
	for sl.level > 1 && sl.head.next[sl.level-1] == nil {
		sl.level--
	}
	sl.len--
	return true
}

func (sl *SkipList[K, V]) ForEach(op func(K, V)) {
	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		op(e.key, e.value)
	}
}

func (sl *SkipList[K, V]) ForEachMutable(op func(K, *V)) {
	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		op(e.key, &e.value)
	}
}

func (sl *SkipList[K, V]) ForEachIf(op func(K, V) bool) {
	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		if !op(e.key, e.value) {
			return
		}
	}
}

func (sl *SkipList[K, V]) ForEachMutableIf(op func(K, *V) bool) {
	for e := sl.head.next[0]; e != nil; e = e.next[0] {
		if !op(e.key, &e.value) {
			return
		}
	}
}
func (sl *SkipList[K, V]) randomLevel() int {
	total := uint64(1)<<uint64(skipListMaxLevel) - 1 // 2^n-1
	k := sl.rander.Uint64() % total
	levelN := uint64(1) << (uint64(skipListMaxLevel) - 1)

	level := 1
	for total -= levelN; total > k; level++ {
		levelN >>= 1
		total -= levelN
		// Since levels are randomly generated, most should be less than log2(s.len).
		// Then make a limit according to sl.len to avoid unexpectedly large value.
		if level > 2 && 1<<(level-2) > sl.len {
			break
		}
	}
	return level
}

//go:generate bash ./skiplist_findnode_generate.sh skiplist_findnode.go
// func (sl *SkipList[K, V]) findNode(key K) *skipListNode[K, V]

// findNodeFast find node with builtin comparation, which is faster than function call.
func findNodeFast[K Ordered, V any](sl *SkipList[K, V], key K) *skipListNode[K, V] {
	var pre = &sl.head
	for i := sl.level - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			if cur.key == key {
				return cur
			}
			if cur.key > key {
				break
			}
			pre = cur
		}
	}
	return nil
}

// findNodeSlow use keyCmp to compare key, which is slower.
func (sl *SkipList[K, V]) findNodeSlow(key K) *skipListNode[K, V] {
	var pre = &sl.head
	for i := sl.level - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := sl.keyCmp(cur.key, key)
			if cmpRet == 0 {
				return cur
			}
			if cmpRet > 0 {
				break
			}
			pre = cur
		}
	}
	return nil
}

// findInsertPoint returns (*node, nil) to the existed node if the key exists,
// or (nil, []*node) to the previous nodes if the key doesn't exist
func (sl *SkipList[K, V]) findInsertPoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := sl.prevsCache[0:sl.level]
	prev := &sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			r := sl.keyCmp(next.key, key)
			if r == 0 {
				// The key is already existed, prevs are useless because no new node insertion.
				// stop searching.
				return next, nil
			}
			if r > 0 {
				// All other node in this level must be greater than the key,
				// search the next level.
				break
			}
			prev = next
		}
		prevs[i] = prev
	}
	return nil, prevs
}

// findRemovePoint finds the node which match the key and it's previous nodes.
func (sl *SkipList[K, V]) findRemovePoint(key K) (*skipListNode[K, V], []*skipListNode[K, V]) {
	prevs := sl.findPrevNodes(key)
	node := prevs[0].next[0]
	if node == nil || sl.keyCmp(node.key, key) != 0 {
		return nil, nil
	}
	return node, prevs
}

func (sl *SkipList[K, V]) findPrevNodes(key K) []*skipListNode[K, V] {
	prevs := sl.prevsCache[0:sl.level]
	prev := &sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for next := prev.next[i]; next != nil; next = next.next[i] {
			if sl.keyCmp(next.key, key) >= 0 {
				break
			}
			prev = next
		}
		prevs[i] = prev
	}
	return prevs
}
