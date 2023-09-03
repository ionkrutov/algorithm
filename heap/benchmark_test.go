package heap

import (
	"math/rand"
	"testing"
)

func benchmarkInsertMultiple(b *testing.B, arraySize int) {
	array := rand.Perm(arraySize)
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	h.Insert(array...)
}

func benchmarkInsertByOne(b *testing.B, arraySize int) {
	array := rand.Perm(arraySize)
	h := Heap[int]{Comparator: func(i1, i2 int) bool { return i1 > i2 }}
	for _, val := range array {
		h.Insert(val)
	}
}

func BenchmarkIM10K(b *testing.B) {
	benchmarkInsertMultiple(b, 10_000)
}

func BenchmarkIM100K(b *testing.B) {
	benchmarkInsertMultiple(b, 100_000)
}

func BenchmarkIM1M(b *testing.B) {
	benchmarkInsertMultiple(b, 1_000_000)
}

func BenchmarkIM10M(b *testing.B) {
	benchmarkInsertMultiple(b, 10_000_000)
}

func BenchmarkIBO10K(b *testing.B) {
	benchmarkInsertByOne(b, 10_000)
}

func BenchmarkIBO100K(b *testing.B) {
	benchmarkInsertByOne(b, 100_000)
}

func BenchmarkIBO1M(b *testing.B) {
	benchmarkInsertByOne(b, 1_000_000)
}

func BenchmarkIBO10M(b *testing.B) {
	benchmarkInsertByOne(b, 10_000_000)
}
