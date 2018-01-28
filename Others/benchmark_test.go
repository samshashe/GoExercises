package main

import (
	"fmt"
	"gobook/fibo"
	"testing"
)

func BenchmarkFiboLoop(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fibo.Fiborecursion(10)
	}
}

func BenchmarkFiboRecurssion(b *testing.B) {

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Fiboloop(10)
	}
}
