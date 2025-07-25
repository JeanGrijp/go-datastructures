// Package fibonacci provides implementations of the Fibonacci sequence algorithm
// using different approaches: recursive, iterative, memoized, and matrix-based.
// The Fibonacci sequence is a famous mathematical sequence where each number
// is the sum of the two preceding ones, usually starting with 0 and 1.
//
// The sequence goes: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
//
// This package includes optimized implementations suitable for different use cases,
// from educational purposes to high-performance computing.
package fibonacci

import (
	"fmt"
	"math/big"
)

// Fibonacci calculates the nth Fibonacci number using an iterative approach.
// This is the most efficient implementation for most use cases.
//
// Parameters:
//   - n: The position in the Fibonacci sequence (0-indexed)
//
// Returns:
//   - int: The nth Fibonacci number
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// Example usage:
//
//	fib := fibonacci.Fibonacci(10)
//	fmt.Println(fib) // Output: 55
//
//	// First 10 Fibonacci numbers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
func Fibonacci(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// FibonacciRecursive calculates the nth Fibonacci number using recursion.
// This is the classic textbook implementation but is inefficient for large n.
// Use only for educational purposes or small values of n.
//
// Parameters:
//   - n: The position in the Fibonacci sequence (0-indexed)
//
// Returns:
//   - int: The nth Fibonacci number
//
// Time complexity: O(2^n) - exponential!
// Space complexity: O(n) - due to recursion stack
//
// Warning: This implementation is very slow for n > 40
//
// Example usage:
//
//	fib := fibonacci.FibonacciRecursive(8)
//	fmt.Println(fib) // Output: 21
func FibonacciRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciMemoized calculates the nth Fibonacci number using memoization.
// This combines the elegance of recursion with the efficiency of dynamic programming.
//
// Parameters:
//   - n: The position in the Fibonacci sequence (0-indexed)
//
// Returns:
//   - int: The nth Fibonacci number
//
// Time complexity: O(n)
// Space complexity: O(n)
//
// Example usage:
//
//	fib := fibonacci.FibonacciMemoized(50)
//	fmt.Println(fib) // Output: 12586269025
func FibonacciMemoized(n int) int {
	if n < 0 {
		return 0
	}

	memo := make(map[int]int)
	return fibMemoHelper(n, memo)
}

// fibMemoHelper is a helper function for memoized Fibonacci calculation
func fibMemoHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = fibMemoHelper(n-1, memo) + fibMemoHelper(n-2, memo)
	return memo[n]
}

// FibonacciBig calculates the nth Fibonacci number using big.Int for very large numbers.
// This allows calculation of Fibonacci numbers that exceed the range of standard integers.
//
// Parameters:
//   - n: The position in the Fibonacci sequence (0-indexed)
//
// Returns:
//   - *big.Int: The nth Fibonacci number as a big integer
//
// Time complexity: O(n)
// Space complexity: O(1) for variables, O(log(result)) for the big.Int storage
//
// Example usage:
//
//	fib := fibonacci.FibonacciBig(1000)
//	fmt.Printf("F(1000) has %d digits\n", len(fib.String()))
func FibonacciBig(n int) *big.Int {
	if n < 0 {
		return big.NewInt(0)
	}
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	temp := big.NewInt(0)

	for i := 2; i <= n; i++ {
		temp.Add(a, b)
		a.Set(b)
		b.Set(temp)
	}

	return b
}

// FibonacciMatrix calculates the nth Fibonacci number using matrix exponentiation.
// This is the fastest method for calculating individual large Fibonacci numbers.
//
// Parameters:
//   - n: The position in the Fibonacci sequence (0-indexed)
//
// Returns:
//   - int: The nth Fibonacci number
//
// Time complexity: O(log n)
// Space complexity: O(log n) - due to recursion in matrix exponentiation
//
// Example usage:
//
//	fib := fibonacci.FibonacciMatrix(100)
//	fmt.Println(fib) // Very fast even for large n
func FibonacciMatrix(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return n
	}

	// Base matrix [[1, 1], [1, 0]]
	result := matrixPower([][]int{{1, 1}, {1, 0}}, n-1)
	return result[0][0]
}

// matrixPower calculates matrix^n using fast exponentiation
func matrixPower(matrix [][]int, n int) [][]int {
	if n == 1 {
		return matrix
	}

	if n%2 == 0 {
		half := matrixPower(matrix, n/2)
		return matrixMultiply(half, half)
	}

	return matrixMultiply(matrix, matrixPower(matrix, n-1))
}

// matrixMultiply multiplies two 2x2 matrices
func matrixMultiply(a, b [][]int) [][]int {
	return [][]int{
		{a[0][0]*b[0][0] + a[0][1]*b[1][0], a[0][0]*b[0][1] + a[0][1]*b[1][1]},
		{a[1][0]*b[0][0] + a[1][1]*b[1][0], a[1][0]*b[0][1] + a[1][1]*b[1][1]},
	}
}

// FibonacciSequence generates the first n Fibonacci numbers.
//
// Parameters:
//   - n: The number of Fibonacci numbers to generate
//
// Returns:
//   - []int: A slice containing the first n Fibonacci numbers
//
// Time complexity: O(n)
// Space complexity: O(n)
//
// Example usage:
//
//	sequence := fibonacci.FibonacciSequence(10)
//	fmt.Println(sequence) // Output: [0 1 1 2 3 5 8 13 21 34]
func FibonacciSequence(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}

	sequence := make([]int, n)
	sequence[0] = 0
	sequence[1] = 1

	for i := 2; i < n; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}

// IsValidFibonacci checks if a given number is a Fibonacci number.
// Uses the mathematical property that a number n is Fibonacci if and only if
// one of (5*n^2 + 4) or (5*n^2 - 4) is a perfect square.
//
// Parameters:
//   - num: The number to check
//
// Returns:
//   - bool: true if the number is a Fibonacci number, false otherwise
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example usage:
//
//	fmt.Println(fibonacci.IsValidFibonacci(21)) // Output: true
//	fmt.Println(fibonacci.IsValidFibonacci(22)) // Output: false
func IsValidFibonacci(num int) bool {
	if num < 0 {
		return false
	}

	return isPerfectSquare(5*num*num+4) || isPerfectSquare(5*num*num-4)
}

// isPerfectSquare checks if a number is a perfect square
func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}

	sqrt := int(float64(n) + 0.5) // Add 0.5 for rounding
	for sqrt*sqrt > n {
		sqrt = (sqrt + n/sqrt) / 2 // Newton's method
	}

	return sqrt*sqrt == n
}

// FibonacciIndex finds the index of a Fibonacci number in the sequence.
// Returns -1 if the number is not a Fibonacci number.
//
// Parameters:
//   - num: The Fibonacci number to find the index for
//
// Returns:
//   - int: The index of the number in the Fibonacci sequence, or -1 if not found
//
// Time complexity: O(num)
// Space complexity: O(1)
//
// Example usage:
//
//	index := fibonacci.FibonacciIndex(21)
//	fmt.Println(index) // Output: 8 (21 is the 8th Fibonacci number)
func FibonacciIndex(num int) int {
	if num < 0 {
		return -1
	}
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1 // Could also be 2, but we return the first occurrence
	}

	if !IsValidFibonacci(num) {
		return -1
	}

	// Linear search approach - more reliable for this case
	a, b := 0, 1
	for i := 2; ; i++ {
		next := a + b
		if next == num {
			return i
		}
		if next > num {
			break
		}
		a, b = b, next
	}

	return -1
}

// GoldenRatio calculates an approximation of the golden ratio using Fibonacci numbers.
// As n increases, the ratio F(n+1)/F(n) approaches the golden ratio (φ ≈ 1.618033988749).
//
// Parameters:
//   - n: The position to calculate the ratio at (should be reasonably large for accuracy)
//
// Returns:
//   - float64: An approximation of the golden ratio
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// Example usage:
//
//	ratio := fibonacci.GoldenRatio(50)
//	fmt.Printf("Golden ratio approximation: %.15f\n", ratio)
func GoldenRatio(n int) float64 {
	if n <= 0 {
		return 0
	}

	fn := float64(Fibonacci(n))
	fn1 := float64(Fibonacci(n + 1))

	if fn == 0 {
		return 0
	}

	return fn1 / fn
}

// FibonacciSum calculates the sum of the first n Fibonacci numbers.
// Uses the mathematical property: Sum(F(0) to F(n)) = F(n+2) - 1
//
// Parameters:
//   - n: The number of Fibonacci numbers to sum (starting from F(0))
//
// Returns:
//   - int: The sum of the first n Fibonacci numbers
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// Example usage:
//
//	sum := fibonacci.FibonacciSum(6)
//	fmt.Println(sum) // Output: 12 (0+1+1+2+3+5 = 12)
func FibonacciSum(n int) int {
	if n <= 0 {
		return 0
	}

	return Fibonacci(n+2) - 1
}

// Visualize demonstrates the calculation process of Fibonacci numbers.
// This is useful for educational purposes to understand how the algorithm works.
//
// Parameters:
//   - n: The position in the Fibonacci sequence to calculate
//
// Returns:
//   - int: The nth Fibonacci number
//
// Example usage:
//
//	result := fibonacci.Visualize(7)
//	// This will print the step-by-step calculation
func Visualize(n int) int {
	fmt.Printf("Calculating Fibonacci(%d) step by step:\n", n)

	if n < 0 {
		fmt.Println("Input is negative, returning 0")
		return 0
	}
	if n <= 1 {
		fmt.Printf("Base case: F(%d) = %d\n", n, n)
		return n
	}

	fmt.Printf("F(0) = 0\n")
	fmt.Printf("F(1) = 1\n")

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		next := a + b
		fmt.Printf("F(%d) = F(%d) + F(%d) = %d + %d = %d\n", i, i-2, i-1, a, b, next)
		a, b = b, next
	}

	fmt.Printf("Result: F(%d) = %d\n", n, b)
	return b
}
