package goroutines

import (
	"testing"
)

func BenchmarkGetUniqueIDGoroutines(b *testing.B) {
	gen := NewIDGenerator(100)
	for i := 0; i < b.N; i++ {
		gen.GetUniqueID()
	}
}

func BenchmarkGetUniqueIDMutex(b *testing.B) {
	gen := MutexNewIDGenerator(100)
	for i := 0; i < b.N; i++ {
		gen.GetUniqueID()
	}
}