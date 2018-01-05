package gotest

import "testing"

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func Benchmark_Loops(b *testing.B) {
	var test ForTest
	ptr := &test

	for i := 0; i < b.N; i++ {
		ptr.Loops()
	}
}

func Benchmark_LoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var test ForTest
		ptr := &test
		for pb.Next() {
			ptr.Loops()
		}
	})
}