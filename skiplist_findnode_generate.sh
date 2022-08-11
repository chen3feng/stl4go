#!/bin/bash

# The go generics sucks!
# We have to generate skiplist_newnode.go by myself.

generate() {
	# NOTE the tab chars, they are used to follow the go style guides.
    echo "package $GOPACKAGE"

	echo "
// findNode find the node which has specified key.
func (sl *SkipList[K, V]) findNode(key K) *skipListNode[K, V] {
	if !sl.builtinCompare {
		return sl.findNodeSlow(key)
	}
	// For knowned Ordered types, use findNodeFast to improve performance.
	var iface interface{} = key
	switch v := iface.(type) {"

	for type in int string int32 uint32 int64 uint64 int8 uint8 int16 uint16 uintptr float32 float64; do
		echo "\
	case $type:
		tsl := pointerCast[*SkipList[$type, V]](sl)
		return pointerCast[*skipListNode[K, V]](findNodeFast(tsl, v))"
	done
	echo "\
	}
	return sl.findNodeSlow(key)
}"

}

output_file=$1

echo -e "// AUTO GENERATED CODE, DON'T EDIT!!!\n// EDIT $(basename $0) accordingly.\n" > $output_file
generate >> $output_file
