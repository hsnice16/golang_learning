package gotest

import "testing"

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { // use b.N for looping
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer()	// call the function to stop the stress test time count

	// Do some initialization work, such as reading file data, database connections and the like,
	// So that our benchmarks reflect the performance of the function itself

	b.StartTimer() // re-start time
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}