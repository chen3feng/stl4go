package stl4go

import (
	"testing"
)

const (
	benchInitSize  = 1000000
	benchBatchSize = 10
)

func newMapN(n int) map[int]int {
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkSkipList_Insert(b *testing.B) {
	start := benchInitSize
	sl := newSkipListN(start)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			sl.Insert(start+i, i)
		}
		start += benchBatchSize
	}
}

func BenchmarkMap_Insert(b *testing.B) {
	start := benchInitSize
	m := newMapN(start)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			m[start+i] = i
		}
		start += benchBatchSize
	}
}

func BenchmarkSkipList_Insert_Dup(b *testing.B) {
	sl := newSkipListN(benchInitSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			sl.Insert(i, i)
		}
	}
}

func BenchmarkMap_Insert_Dup(b *testing.B) {
	m := newMapN(benchInitSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			m[i] = i
		}
	}
}

func BenchmarkMap_Find(b *testing.B) {
	m := newMapN(benchInitSize)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < benchBatchSize; i++ {
			_, _ = m[i]
		}
	}
}
func BenchmarkSkipList_Find(b *testing.B) {
	sl := newSkipListN(benchInitSize)
	b.ResetTimer()
	b.Run("Find", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.findNode(n)
			}
		}
	})
	b.Run("FindSlow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.findNodeSlow(n)
			}
		}
	})
	b.Run("FindEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.findNode(benchInitSize)
			}
		}
	})
	b.Run("FindEndSlow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.findNodeSlow(benchInitSize)
			}
		}
	})
}
