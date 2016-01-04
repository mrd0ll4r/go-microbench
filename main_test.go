package go_microbench

import "testing"

func BenchmarkEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}
