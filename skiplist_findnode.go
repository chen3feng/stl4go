// AUTO GENERATED CODE, DON'T EDIT!!!
// EDIT skiplist_findnode_generate.sh accordingly.

package stl4go

import "unsafe"

func forceCast[T any, F any](f *F) *T {
	//#nosec G103
	return (*T)(unsafe.Pointer(f))
}

func (sl *SkipList[K, V]) findNode(key K) *skipListNode[K, V] {
	if !sl.builtinCompare {
		return sl.findNodeSlow(key)
	}
	var iface interface{} = key
	switch iface.(type) {
	case int:
		tsl := forceCast[SkipList[int, V]](sl)
		v, _ := iface.(int)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case string:
		tsl := forceCast[SkipList[string, V]](sl)
		v, _ := iface.(string)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case int32:
		tsl := forceCast[SkipList[int32, V]](sl)
		v, _ := iface.(int32)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case uint32:
		tsl := forceCast[SkipList[uint32, V]](sl)
		v, _ := iface.(uint32)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case int64:
		tsl := forceCast[SkipList[int64, V]](sl)
		v, _ := iface.(int64)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case uint64:
		tsl := forceCast[SkipList[uint64, V]](sl)
		v, _ := iface.(uint64)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case int8:
		tsl := forceCast[SkipList[int8, V]](sl)
		v, _ := iface.(int8)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case uint8:
		tsl := forceCast[SkipList[uint8, V]](sl)
		v, _ := iface.(uint8)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case int16:
		tsl := forceCast[SkipList[int16, V]](sl)
		v, _ := iface.(int16)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case uint16:
		tsl := forceCast[SkipList[uint16, V]](sl)
		v, _ := iface.(uint16)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case uintptr:
		tsl := forceCast[SkipList[uintptr, V]](sl)
		v, _ := iface.(uintptr)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case float32:
		tsl := forceCast[SkipList[float32, V]](sl)
		v, _ := iface.(float32)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	case float64:
		tsl := forceCast[SkipList[float64, V]](sl)
		v, _ := iface.(float64)
		return forceCast[skipListNode[K, V]](findNodeFast(tsl, v))
	}
	return sl.findNodeSlow(key)
}
