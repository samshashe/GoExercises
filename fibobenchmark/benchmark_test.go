package fibobenchmark

import (
	"gobook/fibo"
	"testing"
)

func BenchmarkFiboLoop(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fibo.Fiborecursion(50)
	}
}

func BenchmarkFiboRecurssion(b *testing.B) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fibo.Fiboloop(50)
	}
}
