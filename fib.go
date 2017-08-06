package main

import (
	"fmt"
	"log"
)

func main() {
	n := 10
	if n < 0 {
		log.Fatalf("%#v", fmt.Errorf("Only positive number is allowed"))
	}
	f := Fib{}
	x := f.Calculate(n)
	fmt.Printf("Fibonacci for %d is %d", n, x)
}

type Fib struct{}

func (f Fib) Calculate(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return f.Calculate(n-1) + f.Calculate(n-2)
}
