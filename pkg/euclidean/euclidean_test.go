package euclidean

import (
	"testing"
)

// TestGCD tests the iterative GCD function with various inputs
func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Basic case", 48, 18, 6},
		{"Prime numbers", 17, 13, 1},
		{"One is multiple of another", 15, 5, 5},
		{"Same numbers", 12, 12, 12},
		{"Zero as second argument", 42, 0, 42},
		{"Zero as first argument", 0, 35, 35},
		{"Both zero", 0, 0, 0},
		{"Negative numbers", -48, 18, 6},
		{"Both negative", -48, -18, 6},
		{"Large numbers", 1071, 462, 21},
		{"Farm example from book", 1680, 1050, 210},
		{"Rectangle example", 12, 8, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestGCDRecursive tests the recursive GCD function
func TestGCDRecursive(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Basic case", 48, 18, 6},
		{"Prime numbers", 17, 13, 1},
		{"One is multiple of another", 15, 5, 5},
		{"Same numbers", 12, 12, 12},
		{"Zero as second argument", 42, 0, 42},
		{"Zero as first argument", 0, 35, 35},
		{"Negative numbers", -56, 42, 14},
		{"Both negative", -56, -42, 14},
		{"Large numbers", 1071, 462, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCDRecursive(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCDRecursive(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestGCDConsistency verifies that both GCD implementations give the same result
func TestGCDConsistency(t *testing.T) {
	testCases := [][]int{
		{48, 18},
		{1071, 462},
		{17, 13},
		{0, 35},
		{42, 0},
		{-48, 18},
		{56, -42},
		{-56, -42},
		{1680, 1050},
		{100, 25},
		{7, 11},
	}

	for _, tc := range testCases {
		a, b := tc[0], tc[1]
		iterative := GCD(a, b)
		recursive := GCDRecursive(a, b)

		if iterative != recursive {
			t.Errorf("GCD methods inconsistent for (%d, %d): iterative=%d, recursive=%d",
				a, b, iterative, recursive)
		}
	}
}

// TestLCM tests the Least Common Multiple function
func TestLCM(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Basic case", 12, 18, 36},
		{"Prime numbers", 7, 11, 77},
		{"One is multiple of another", 6, 3, 6},
		{"Same numbers", 5, 5, 5},
		{"With zero", 0, 5, 0},
		{"Both zero", 0, 0, 0},
		{"Negative numbers", -12, 18, 36},
		{"Large numbers", 15, 25, 75},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LCM(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("LCM(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestExtendedGCD tests the Extended Euclidean Algorithm
func TestExtendedGCD(t *testing.T) {
	tests := []struct {
		name string
		a, b int
	}{
		{"Basic case", 30, 18},
		{"Prime numbers", 7, 11},
		{"Large numbers", 1071, 462},
		{"Simple case", 10, 6},
		{"With one", 25, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcd, x, y := ExtendedGCD(tt.a, tt.b)

			// Verify that gcd matches the regular GCD
			expectedGCD := GCD(tt.a, tt.b)
			if gcd != expectedGCD {
				t.Errorf("ExtendedGCD(%d, %d) gcd = %d; expected %d", tt.a, tt.b, gcd, expectedGCD)
			}

			// Verify the Bézout identity: ax + by = gcd(a, b)
			if tt.a*x+tt.b*y != gcd {
				t.Errorf("ExtendedGCD(%d, %d): %d*%d + %d*%d = %d; expected %d",
					tt.a, tt.b, tt.a, x, tt.b, y, tt.a*x+tt.b*y, gcd)
			}
		})
	}
}

// TestLargestSquareSize tests the geometric application of GCD
func TestLargestSquareSize(t *testing.T) {
	tests := []struct {
		name           string
		width, height  int
		expectedSquare int
	}{
		{"Book example", 1680, 1050, 210},
		{"Simple rectangle", 12, 8, 4},
		{"Square input", 10, 10, 10},
		{"Prime dimensions", 7, 11, 1},
		{"One dimension is 1", 15, 1, 1},
		{"Large rectangle", 100, 60, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestSquareSize(tt.width, tt.height)
			if result != tt.expectedSquare {
				t.Errorf("LargestSquareSize(%d, %d) = %d; expected %d",
					tt.width, tt.height, result, tt.expectedSquare)
			}
		})
	}
}

// TestSquareDivision tests the square division calculation
func TestSquareDivision(t *testing.T) {
	tests := []struct {
		name                        string
		width, height               int
		expectedSize, expectedCount int
	}{
		{"Simple rectangle", 12, 8, 4, 6},     // 12*8 = 96, 4*4 = 16, 96/16 = 6
		{"Book example", 1680, 1050, 210, 40}, // 1680*1050 = 1764000, 210*210 = 44100, 1764000/44100 = 40
		{"Square input", 10, 10, 10, 1},       // 10*10 = 100, 10*10 = 100, 100/100 = 1
		{"Small rectangle", 6, 4, 2, 6},       // 6*4 = 24, 2*2 = 4, 24/4 = 6
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, count := SquareDivision(tt.width, tt.height)
			if size != tt.expectedSize {
				t.Errorf("SquareDivision(%d, %d) size = %d; expected %d",
					tt.width, tt.height, size, tt.expectedSize)
			}
			if count != tt.expectedCount {
				t.Errorf("SquareDivision(%d, %d) count = %d; expected %d",
					tt.width, tt.height, count, tt.expectedCount)
			}
		})
	}
}

// TestGCDMultiple tests GCD calculation for multiple numbers
func TestGCDMultiple(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"Empty slice", []int{}, 0},
		{"Single number", []int{42}, 42},
		{"Two numbers", []int{48, 18}, 6},
		{"Three numbers", []int{48, 18, 24}, 6},
		{"Multiple numbers", []int{12, 18, 24, 30}, 6},
		{"Prime numbers", []int{7, 11, 13}, 1},
		{"With zero", []int{0, 12, 18}, 6},
		{"All same", []int{15, 15, 15}, 15},
		{"Negative numbers", []int{-12, 18, 24}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCDMultiple(tt.numbers)
			if result != tt.expected {
				t.Errorf("GCDMultiple(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}

// TestIsCoprime tests the coprime check function
func TestIsCoprime(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected bool
	}{
		{"Coprime numbers", 15, 28, true},
		{"Not coprime", 15, 25, false},
		{"Prime numbers", 7, 11, true},
		{"Same number", 5, 5, false},
		{"One is 1", 1, 42, true},
		{"Both 1", 1, 1, true},
		{"Negative coprime", -15, 28, true},
		{"Large coprime", 101, 103, true},
		{"Large not coprime", 100, 150, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCoprime(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("IsCoprime(%d, %d) = %t; expected %t", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestVisualize tests that the visualize function returns correct GCD
// (We can't easily test the printed output in unit tests)
func TestVisualize(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Basic case", 48, 18, 6},
		{"Large numbers", 1071, 462, 21},
		{"Simple case", 10, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Visualize(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Visualize(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance comparison

// BenchmarkGCD benchmarks the iterative GCD implementation
func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(1071, 462)
	}
}

// BenchmarkGCDRecursive benchmarks the recursive GCD implementation
func BenchmarkGCDRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCDRecursive(1071, 462)
	}
}

// BenchmarkGCDLarge benchmarks GCD with large numbers
func BenchmarkGCDLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(123456789, 987654321)
	}
}

// BenchmarkExtendedGCD benchmarks the Extended Euclidean Algorithm
func BenchmarkExtendedGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExtendedGCD(1071, 462)
	}
}

// BenchmarkLCM benchmarks the LCM calculation
func BenchmarkLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LCM(1071, 462)
	}
}

// BenchmarkGCDMultiple benchmarks GCD calculation for multiple numbers
func BenchmarkGCDMultiple(b *testing.B) {
	numbers := []int{48, 18, 24, 30, 36}
	for i := 0; i < b.N; i++ {
		GCDMultiple(numbers)
	}
}
