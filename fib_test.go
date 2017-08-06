package main

import "testing"

func Test_Fib(t *testing.T) {
	f := Fib{}
	actual := f.Calculate(10)
	if actual != 55 {
		t.Fatalf("Fib %d must be 55", 10)
	}
}

func Bench_Fib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ := f.Calculate(10)
	}
}
