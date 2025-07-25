package fatorial

import (
	"fmt"
	"testing"
)

// TestFactorial tests the Factorial function with various inputs
func TestFactorial(t *testing.T) {
	testCases := []struct {
		input    int
		expected string // Using string to handle large numbers
	}{
		{0, "1"},
		{1, "1"},
		{5, "120"},
		{10, "3628800"},
		{20, "2432902008176640000"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Factorial(%d)", tc.input), func(t *testing.T) {
			result := Factorial(tc.input)
			if result.String() != tc.expected {
				t.Errorf("Factorial(%d) = %s; expected %s", tc.input, result.String(), tc.expected)
			}
		})
	}
}

// TestFatorial tests the Portuguese version function
func TestFatorial(t *testing.T) {
	result1 := Factorial(10)
	result2 := Fatorial(10)

	if result1.Cmp(result2) != 0 {
		t.Errorf("Factorial and Fatorial should return the same result")
	}
}

// BenchmarkFactorial benchmarks the Factorial function
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(100)
	}
}

// ExampleFactorial demonstrates basic usage of the Factorial function
func ExampleFactorial() {
	result := Factorial(5)
	fmt.Println("5! =", result)
	// Output: 5! = 120
}

// ExampleFactorial_large demonstrates factorial calculation for large numbers
func ExampleFactorial_large() {
	result := Factorial(100)
	fmt.Printf("100! has %d digits\n", len(result.String()))
	// Output: 100! has 158 digits
}

// ExampleFatorial demonstrates the Portuguese version
func ExampleFatorial() {
	resultado := Fatorial(6)
	fmt.Println("6! =", resultado)
	// Output: 6! = 720
}
