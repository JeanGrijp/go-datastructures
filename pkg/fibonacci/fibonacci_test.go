package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	if Fibonacci(5) != 5 {
		t.Error("Fibonacci(5) should be 5")
	}
}

func TestFibonacciRecursive(t *testing.T) {
	if FibonacciRecursive(5) != 5 {
		t.Error("FibonacciRecursive(5) should be 5")
	}
}

func TestFibonacciMemoized(t *testing.T) {
	if FibonacciMemoized(10) != 55 {
		t.Error("FibonacciMemoized(10) should be 55")
	}
}

func TestFibonacciMatrix(t *testing.T) {
	if FibonacciMatrix(10) != 55 {
		t.Error("FibonacciMatrix(10) should be 55")
	}
}

func TestFibonacciSequence(t *testing.T) {
	seq := FibonacciSequence(5)
	if len(seq) != 5 {
		t.Error("FibonacciSequence(5) should have 5 elements")
	}
}

func TestIsValidFibonacci(t *testing.T) {
	if !IsValidFibonacci(8) {
		t.Error("8 should be a valid Fibonacci number")
	}
	if IsValidFibonacci(6) {
		t.Error("6 should not be a valid Fibonacci number")
	}
}

func TestFibonacciIndex(t *testing.T) {
	if FibonacciIndex(8) != 6 {
		t.Error("Index of 8 should be 6")
	}
}

func TestFibonacciSum(t *testing.T) {
	if FibonacciSum(5) != 12 {
		t.Error("Sum of first 5 Fibonacci numbers should be 12")
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonacciMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMemoized(20)
	}
}

func BenchmarkFibonacciMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMatrix(20)
	}
}
