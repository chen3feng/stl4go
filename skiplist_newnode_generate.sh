#!/bin/bash

# The go generics sucks!
# We have to generate skiplist_newnode.go by myself.

generate() {
	# NOTE the tab chars, they are used to follow the go style guides.
    echo "package $GOPACKAGE"
    echo "
// newSkipListNode creates a new node initialized with specified key, value and next slice.
func newSkipListNode[K any, V any](level int, key K, value V) *skipListNode[K, V] {
	// For nodes with each levels, point their next slice to the nexts array allocated together,
	// which can reduce 1 memory allocation and improve performance.
	//
	// The generics of the golang doesn't support non-type parameters like in C++,
	// so we have to generate it manually.
	switch level {"

    for i in $(seq 1 $1); do
        echo "\
	case $i:
		n := struct {
			head  skipListNode[K, V]
			nexts [$i]*skipListNode[K, V]
		}{head: skipListNode[K, V]{key, value, nil}}
		n.head.next = n.nexts[:]
		return &n.head"
    done
    echo "\
	}

	panic(\"should not reach here\")
}"
}

max_level=$(grep -E "^\s*$1\s*=\s*\d+" $GOFILE | grep -Eo "\d+")
output_file=$2

echo -e "// AUTO GENERATED CODE, DON'T EDIT!!!\n// EDIT $(basename $0) accordingly.\n" > $output_file
generate $max_level >> $output_file
