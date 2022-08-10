// AUTO GENERATED CODE, DON'T EDIT!!!
// EDIT skiplist_findnode_generate.sh accordingly.

package stl4go

// findNode find the node which has specified key.
func (sl *SkipList[K, V]) findNode(key K) *skipListNode[K, V] {
	if !sl.builtinCompare {
		return sl.findNodeSlow(key)
	}
	// For knowned Ordered types, use findNodeFast to improve performance.
	var iface interface{} = key
	switch iface.(type) {
	case int:
		tsl := pointerCast[*SkipList[int, V]](sl)
		v, _ := iface.(int)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case string:
		tsl := pointerCast[*SkipList[string, V]](sl)
		v, _ := iface.(string)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case int32:
		tsl := pointerCast[*SkipList[int32, V]](sl)
		v, _ := iface.(int32)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case uint32:
		tsl := pointerCast[*SkipList[uint32, V]](sl)
		v, _ := iface.(uint32)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case int64:
		tsl := pointerCast[*SkipList[int64, V]](sl)
		v, _ := iface.(int64)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case uint64:
		tsl := pointerCast[*SkipList[uint64, V]](sl)
		v, _ := iface.(uint64)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case int8:
		tsl := pointerCast[*SkipList[int8, V]](sl)
		v, _ := iface.(int8)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case uint8:
		tsl := pointerCast[*SkipList[uint8, V]](sl)
		v, _ := iface.(uint8)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case int16:
		tsl := pointerCast[*SkipList[int16, V]](sl)
		v, _ := iface.(int16)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case uint16:
		tsl := pointerCast[*SkipList[uint16, V]](sl)
		v, _ := iface.(uint16)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case uintptr:
		tsl := pointerCast[*SkipList[uintptr, V]](sl)
		v, _ := iface.(uintptr)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case float32:
		tsl := pointerCast[*SkipList[float32, V]](sl)
		v, _ := iface.(float32)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	case float64:
		tsl := pointerCast[*SkipList[float64, V]](sl)
		v, _ := iface.(float64)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))
	}
	return sl.findNodeSlow(key)
}
