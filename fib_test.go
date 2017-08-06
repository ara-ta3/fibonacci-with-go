package main

import "testing"

func EmptyFib() Fib {
	return Fib{
		cache: map[int]int{},
	}
}

func Test_Fib(t *testing.T) {
	f := EmptyFib()
	actual := f.Calculate(10)
	if actual != 55 {
		t.Fatalf("Fib %d must be 55", 10)
	}
}

func Benchmark_Fib(b *testing.B) {
	f := EmptyFib()
	for i := 0; i < b.N; i++ {
		f.Calculate(10)
	}
}
