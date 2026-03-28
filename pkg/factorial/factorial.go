// Package factorial provides functions to calculate factorials of integers
// using arbitrary precision arithmetic to handle very large numbers.
package factorial

import "math/big"

// Factorial calculates the factorial of n (n!) using big.Int for arbitrary precision.
// The factorial of a non-negative integer n is the product of all positive integers
// less than or equal to n.
//
// For example:
//   - Factorial(0) = 1
//   - Factorial(5) = 5 × 4 × 3 × 2 × 1 = 120
//   - Factorial(100) = 100 × 99 × ... × 2 × 1
//
// This function uses big.Int to avoid integer overflow, allowing calculation
// of factorials for very large numbers (limited only by available memory).
//
// Parameters:
//   - n: A non-negative integer for which to calculate the factorial
//
// Returns:
//   - *big.Int: The factorial of n as an arbitrary precision integer
//
// Example usage:
//
//	result := factorial.Factorial(100)
//	fmt.Println("100! =", result)
//
// Time complexity: O(n)
// Space complexity: O(1) excluding the result
func Factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
