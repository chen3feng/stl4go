package stl4go

import (
	"strconv"
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
				_ = sl.Find(n)
			}
		}
	})
	b.Run("FindSlow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.Find(n)
			}
		}
	})
	b.Run("FindEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.Find(benchInitSize)
			}
		}
	})
	b.Run("FindEndSlow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				_ = sl.Find(benchInitSize)
			}
		}
	})
}

func BenchmarkSkipListString(b *testing.B) {
	sl := NewSkipList[string, int]()
	var a []string
	for i := 0; i < benchBatchSize; i++ {
		a = append(a, strconv.Itoa(benchInitSize+i))
	}
	end := strconv.Itoa(2 * benchInitSize)
	b.ResetTimer()
	b.Run("Insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				sl.Insert(a[n], n)
			}
		}
	})
	b.Run("Find", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				sl.Find(a[n])
			}
		}
	})
	b.Run("FindEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				sl.Find(end)
			}
		}
	})

	b.Run("RemoveEnd", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				sl.Remove(end)
			}
		}
	})
	b.Run("Remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for n := 0; n < benchBatchSize; n++ {
				sl.Remove(a[n])
			}
		}
	})
}
