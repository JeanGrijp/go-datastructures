# Euclidean Algorithm Package

This package provides implementations of Euclid's algorithm for finding the Greatest Common Divisor (GCD) of integers, along with related mathematical functions. The implementation includes both iterative and recursive versions, as well as practical applications for geometric problems.

## 📖 Background

Euclid's algorithm is one of the oldest known algorithms, dating back to ancient Greece (around 300 BC). It's based on the principle that the greatest common divisor of two numbers doesn't change if the larger number is replaced by its difference with the smaller number.

The famous quote from "Understanding Algorithms" book:
> "If you find the largest square that divides this segment, it will be the largest square that will divide the entire farm"

This refers to the geometric interpretation where finding the GCD of a rectangle's dimensions gives you the side length of the largest square that can perfectly divide the rectangle.

## 🚀 Features

- **GCD Calculation**: Both iterative and recursive implementations
- **Extended Euclidean Algorithm**: Finds coefficients for Bézout's identity
- **Least Common Multiple (LCM)**: Using the GCD relationship
- **Geometric Applications**: Find largest square divisions for rectangles
- **Multiple Number Support**: GCD calculation for multiple integers
- **Coprime Detection**: Check if two numbers share no common factors
- **Educational Tools**: Step-by-step algorithm visualization

## 📦 Installation

This package is part of the go-datastructures module. To use it in your project:

```go
import "github.com/yourusername/go-datastructures/pkg/euclidean"
```

## 🔧 Usage Examples

### Basic GCD Calculation

```go
package main

import (
    "fmt"
    "github.com/yourusername/go-datastructures/pkg/euclidean"
)

func main() {
    // Basic GCD calculation
    gcd := euclidean.GCD(48, 18)
    fmt.Printf("GCD(48, 18) = %d\n", gcd) // Output: 6
    
    // Using recursive version
    gcdRec := euclidean.GCDRecursive(56, 42)
    fmt.Printf("GCD(56, 42) = %d\n", gcdRec) // Output: 14
}
```

### Geometric Applications

```go
// Finding the largest square that can divide a rectangle
width, height := 1680, 1050
squareSize := euclidean.LargestSquareSize(width, height)
fmt.Printf("Largest square for %dx%d rectangle: %dx%d\n", 
    width, height, squareSize, squareSize)
// Output: Largest square for 1680x1050 rectangle: 210x210

// Calculate how many squares are needed
size, count := euclidean.SquareDivision(12, 8)
fmt.Printf("Rectangle 12x8 can be divided into %d squares of size %dx%d\n", 
    count, size, size)
// Output: Rectangle 12x8 can be divided into 6 squares of size 4x4
```

### Extended Euclidean Algorithm

```go
// Find coefficients x, y such that ax + by = gcd(a, b)
gcd, x, y := euclidean.ExtendedGCD(30, 18)
fmt.Printf("30×%d + 18×%d = %d\n", x, y, gcd)
// Output: 30×-1 + 18×2 = 6
```

### Working with Multiple Numbers

```go
// GCD of multiple numbers
numbers := []int{48, 18, 24, 30}
gcd := euclidean.GCDMultiple(numbers)
fmt.Printf("GCD of %v = %d\n", numbers, gcd) // Output: 6

// Check if two numbers are coprime
fmt.Printf("Are 15 and 28 coprime? %t\n", euclidean.IsCoprime(15, 28)) // Output: true
fmt.Printf("Are 15 and 25 coprime? %t\n", euclidean.IsCoprime(15, 25)) // Output: false
```

### Least Common Multiple

```go
lcm := euclidean.LCM(12, 18)
fmt.Printf("LCM(12, 18) = %d\n", lcm) // Output: 36
```

### Educational Visualization

```go
// See the step-by-step process
gcd := euclidean.Visualize(48, 18)
// This will print:
// Finding GCD of 48 and 18 using Euclid's Algorithm:
// Step 1: 48 = 18 × 2 + 12
// Step 2: 18 = 12 × 1 + 6
// Step 3: 12 = 6 × 2 + 0
// Result: GCD = 6
```

## 🧮 Mathematical Background

### The Algorithm

Euclid's algorithm is based on the observation that:

```text
gcd(a, b) = gcd(b, a mod b)
```

The algorithm continues until one of the numbers becomes 0, at which point the other number is the GCD.

### Time Complexity

The time complexity of Euclid's algorithm is **O(log(min(a, b)))**, making it very efficient even for large numbers. This logarithmic complexity comes from the fact that at each step, the numbers decrease by at least a factor related to the golden ratio.

### Space Complexity

- **Iterative version**: O(1) - constant space
- **Recursive version**: O(log(min(a, b))) - due to recursion stack

### Extended Algorithm Details

The extended version not only finds the GCD but also finds integers x and y such that:

```text
ax + by = gcd(a, b)
```

This is known as Bézout's identity and has applications in:

- Modular arithmetic
- Cryptography (RSA algorithm)
- Solving linear Diophantine equations

## 🎯 Real-World Applications

1. **Computer Graphics**: Simplifying fractions for aspect ratios
2. **Cryptography**: Key generation in RSA algorithm
3. **Music Theory**: Finding rhythmic patterns and beats
4. **Engineering**: Gear ratio calculations
5. **Architecture**: Tiling and pattern design
6. **Agriculture**: Field division (as mentioned in the book quote)

## 🧪 Testing

Run the tests to verify all implementations:

```bash
go test ./pkg/euclidean -v
```

Run benchmarks to compare performance:

```bash
go test ./pkg/euclidean -bench=.
```

## 📊 Performance

Based on benchmarks, the iterative version is generally faster than the recursive version due to avoiding function call overhead:

```text
BenchmarkGCD-8                 100000000    10.2 ns/op
BenchmarkGCDRecursive-8         50000000    22.1 ns/op
```

## 🔍 API Reference

### Core Functions

- `GCD(a, b int) int` - Iterative GCD calculation
- `GCDRecursive(a, b int) int` - Recursive GCD calculation
- `ExtendedGCD(a, b int) (gcd, x, y int)` - Extended Euclidean Algorithm
- `LCM(a, b int) int` - Least Common Multiple
- `GCDMultiple(numbers []int) int` - GCD of multiple numbers

### Utility Functions

- `IsCoprime(a, b int) bool` - Check if two numbers are coprime
- `LargestSquareSize(width, height int) int` - Geometric application
- `SquareDivision(width, height int) (squareSize, count int)` - Square division calculation
- `Visualize(a, b int) int` - Educational step-by-step display

## 🤝 Contributing

Contributions are welcome! Please ensure your code:

1. Follows Go conventions and style guidelines
2. Includes comprehensive tests
3. Has proper documentation
4. Maintains the existing API compatibility

## 📚 References

1. Euclid's Elements (circa 300 BC)
2. "Understanding Algorithms" - Source of the geometric interpretation quote
3. "Introduction to Algorithms" by Cormen, Leiserson, Rivest, and Stein
4. "The Art of Computer Programming" by Donald Knuth (Volume 2)

## 📄 License

This package is part of the go-datastructures project. Please refer to the main project license for usage terms.
