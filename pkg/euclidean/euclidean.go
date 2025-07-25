// Package euclidean provides implementations of Euclid's algorithm for finding
// the Greatest Common Divisor (GCD) of two integers. This algorithm is famous
// for being one of the oldest known algorithms and has practical applications
// in many areas including finding the largest square that can divide a rectangle.
//
// The quote "If you find the largest square that divides this segment,
// it will be the largest square that will divide the entire farm" refers
// to the geometric interpretation of Euclid's algorithm.
package euclidean

import "fmt"

// GCD calculates the Greatest Common Divisor of two integers using Euclid's algorithm.
// The algorithm works by repeatedly applying the principle that gcd(a, b) = gcd(b, a mod b)
// until one of the numbers becomes 0.
//
// This is the iterative version of Euclid's algorithm, which is more efficient
// in terms of space complexity as it doesn't use recursion.
//
// Parameters:
//   - a: The first integer (must be non-negative)
//   - b: The second integer (must be non-negative)
//
// Returns:
//   - int: The greatest common divisor of a and b
//
// Time complexity: O(log(min(a, b)))
// Space complexity: O(1)
//
// Example usage:
//
//	gcd := euclidean.GCD(48, 18)
//	fmt.Println(gcd) // Output: 6
//
//	// Finding largest square for a 1680x1050 rectangle
//	largest := euclidean.GCD(1680, 1050)
//	fmt.Println(largest) // Output: 210 (210x210 squares)
func GCD(a, b int) int {
	// Handle negative numbers by taking absolute values
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Ensure a >= b for consistency
	if a < b {
		a, b = b, a
	}

	// Apply Euclid's algorithm
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}

	return a
}

// GCDRecursive calculates the Greatest Common Divisor using the recursive version
// of Euclid's algorithm. This is the classic textbook implementation that directly
// follows the mathematical definition.
//
// Parameters:
//   - a: The first integer (must be non-negative)
//   - b: The second integer (must be non-negative)
//
// Returns:
//   - int: The greatest common divisor of a and b
//
// Time complexity: O(log(min(a, b)))
// Space complexity: O(log(min(a, b))) due to recursion stack
//
// Example usage:
//
//	gcd := euclidean.GCDRecursive(56, 42)
//	fmt.Println(gcd) // Output: 14
func GCDRecursive(a, b int) int {
	// Handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	// Base case: if b is 0, then gcd(a, 0) = a
	if b == 0 {
		return a
	}

	// Recursive case: gcd(a, b) = gcd(b, a mod b)
	return GCDRecursive(b, a%b)
}

// LCM calculates the Least Common Multiple of two integers using the relationship:
// lcm(a, b) = (a * b) / gcd(a, b)
//
// Parameters:
//   - a: The first integer (must be non-negative)
//   - b: The second integer (must be non-negative)
//
// Returns:
//   - int: The least common multiple of a and b
//
// Time complexity: O(log(min(a, b)))
// Space complexity: O(1)
//
// Example usage:
//
//	lcm := euclidean.LCM(12, 18)
//	fmt.Println(lcm) // Output: 36
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	// Handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	return (a * b) / GCD(a, b)
}

// ExtendedGCD implements the Extended Euclidean Algorithm, which not only finds
// the GCD of two numbers but also finds the coefficients (x, y) such that:
// ax + by = gcd(a, b)
//
// This is useful in modular arithmetic and cryptography.
//
// Parameters:
//   - a: The first integer
//   - b: The second integer
//
// Returns:
//   - gcd: The greatest common divisor of a and b
//   - x: Coefficient for a in the equation ax + by = gcd(a, b)
//   - y: Coefficient for b in the equation ax + by = gcd(a, b)
//
// Time complexity: O(log(min(a, b)))
// Space complexity: O(1)
//
// Example usage:
//
//	gcd, x, y := euclidean.ExtendedGCD(30, 18)
//	fmt.Printf("gcd=%d, x=%d, y=%d\n", gcd, x, y) // Output: gcd=6, x=-1, y=2
//	// Verification: 30*(-1) + 18*(2) = -30 + 36 = 6
func ExtendedGCD(a, b int) (gcd, x, y int) {
	if b == 0 {
		return a, 1, 0
	}

	gcd1, x1, y1 := ExtendedGCD(b, a%b)
	x = y1
	y = x1 - (a/b)*y1

	return gcd1, x, y
}

// LargestSquareSize finds the side length of the largest square that can divide
// a rectangle with given width and height. This is a direct application of
// Euclid's algorithm to the geometric problem mentioned in the quote.
//
// Parameters:
//   - width: The width of the rectangle
//   - height: The height of the rectangle
//
// Returns:
//   - int: The side length of the largest square that divides the rectangle
//
// Example usage:
//
//	// For a 1680x1050 farm/rectangle
//	squareSize := euclidean.LargestSquareSize(1680, 1050)
//	fmt.Printf("Largest square: %dx%d\n", squareSize, squareSize) // Output: 210x210
func LargestSquareSize(width, height int) int {
	return GCD(width, height)
}

// SquareDivision shows how many squares of the largest possible size
// are needed to completely divide a rectangle.
//
// Parameters:
//   - width: The width of the rectangle
//   - height: The height of the rectangle
//
// Returns:
//   - squareSize: The side length of each square
//   - count: The number of squares needed
//
// Example usage:
//
//	size, count := euclidean.SquareDivision(12, 8)
//	fmt.Printf("Rectangle 12x8 can be divided into %d squares of size %dx%d\n", count, size, size)
//	// Output: Rectangle 12x8 can be divided into 6 squares of size 4x4
func SquareDivision(width, height int) (squareSize, count int) {
	squareSize = GCD(width, height)
	count = (width * height) / (squareSize * squareSize)
	return squareSize, count
}

// GCDMultiple calculates the GCD of multiple integers.
//
// Parameters:
//   - numbers: A slice of integers
//
// Returns:
//   - int: The greatest common divisor of all numbers
//
// Example usage:
//
//	gcd := euclidean.GCDMultiple([]int{48, 18, 24})
//	fmt.Println(gcd) // Output: 6
func GCDMultiple(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		if numbers[0] < 0 {
			return -numbers[0]
		}
		return numbers[0]
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = GCD(result, numbers[i])
		if result == 1 {
			break // If GCD becomes 1, it won't get smaller
		}
	}

	return result
}

// IsCoprime checks if two integers are coprime (their GCD is 1).
// Two numbers are coprime if they share no common factors other than 1.
//
// Parameters:
//   - a: The first integer
//   - b: The second integer
//
// Returns:
//   - bool: true if a and b are coprime, false otherwise
//
// Example usage:
//
//	fmt.Println(euclidean.IsCoprime(15, 28)) // Output: true (gcd = 1)
//	fmt.Println(euclidean.IsCoprime(15, 25)) // Output: false (gcd = 5)
func IsCoprime(a, b int) bool {
	return GCD(a, b) == 1
}

// Visualize demonstrates the step-by-step process of Euclid's algorithm
// by printing each step. This is useful for educational purposes.
//
// Parameters:
//   - a: The first integer
//   - b: The second integer
//
// Returns:
//   - int: The greatest common divisor of a and b
//
// Example usage:
//
//	gcd := euclidean.Visualize(48, 18)
//	// This will print the steps of the algorithm
func Visualize(a, b int) int {
	fmt.Printf("Finding GCD of %d and %d using Euclid's Algorithm:\n", a, b)

	// Handle negative numbers
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	step := 1
	for b != 0 {
		quotient := a / b
		remainder := a % b
		fmt.Printf("Step %d: %d = %d × %d + %d\n", step, a, b, quotient, remainder)
		a, b = b, remainder
		step++
	}

	fmt.Printf("Result: GCD = %d\n", a)
	return a
}
