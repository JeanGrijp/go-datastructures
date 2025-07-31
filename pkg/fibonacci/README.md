# Fibonacci Package

This implementation provides various approaches for calculating Fibonacci numbers, each with its own performance characteristics and use cases.

## Table of Contents

- [About the Fibonacci Sequence](#about-the-fibonacci-sequence)
- [Available Implementations](#available-implementations)
- [Usage Examples](#usage-examples)
- [Educational Tools](#educational-tools)
- [Performance](#performance)
- [Use Cases](#use-cases)

## About the Fibonacci Sequence

The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding numbers. The sequence traditionally starts with 0 and 1:

```
F(0) = 0
F(1) = 1
F(n) = F(n-1) + F(n-2) for n > 1
```

**Sequence:** 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377...

### Mathematical Properties

- **Golden Ratio**: As n increases, F(n+1)/F(n) approaches the golden ratio φ ≈ 1.618033988749
- **Binet's Formula**: F(n) = (φⁿ - ψⁿ)/√5, where ψ = (1-√5)/2
- **Cassini's Identity**: F(n-1)×F(n+1) - F(n)² = (-1)ⁿ

## Available Implementations

### 1. Iterative Implementation - `Fibonacci(n int) int`
- **Complexity**: O(n) time, O(1) space
- **Recommended use**: Small to medium values (n < 50)
- **Characteristics**: Most efficient for most use cases

```go
result := fibonacci.Fibonacci(10) // Returns 55
```

### 2. Recursive Implementation - `FibonacciRecursive(n int) int`
- **Complexity**: O(2ⁿ) time, O(n) space
- **Recommended use**: Educational demonstration only (n < 15)
- **Characteristics**: Most intuitive implementation, but extremely inefficient

```go
result := fibonacci.FibonacciRecursive(8) // Returns 21
```

### 3. Memoized Implementation - `FibonacciMemoized(n int) int`
- **Complexity**: O(n) time, O(n) space
- **Recommended use**: Multiple calls for different values
- **Characteristics**: Combines recursive clarity with efficiency

```go
result := fibonacci.FibonacciMemoized(30) // Returns 832040
```

### 4. Big Numbers Implementation - `FibonacciBig(n int) *big.Int`
- **Complexity**: O(n) time, O(1) space
- **Recommended use**: Very large values (n > 50)
- **Characteristics**: Supports arbitrarily large numbers

```go
result := fibonacci.FibonacciBig(100) // Returns 354224848179261915075
```

### 5. Matrix Implementation - `FibonacciMatrix(n int) int`
- **Complexity**: O(log n) time, O(1) space
- **Recommended use**: Large values when int precision is sufficient
- **Characteristics**: Fastest algorithm for single large values

```go
result := fibonacci.FibonacciMatrix(50) // Returns 12586269025
```

## Usage Examples

### Basic Example
```go
package main

import (
    "fmt"
    "yourmodule/pkg/fibonacci"
)

func main() {
    // Calculate the 10th Fibonacci number
    fmt.Printf("F(10) = %d\n", fibonacci.Fibonacci(10))
    
    // Generate sequence of first 10 numbers
    sequence := fibonacci.FibonacciSequence(10)
    fmt.Printf("First 10: %v\n", sequence)
    
    // Check if a number is Fibonacci
    fmt.Printf("Is 8 Fibonacci? %t\n", fibonacci.IsValidFibonacci(8))
    
    // Calculate very large number
    big := fibonacci.FibonacciBig(200)
    fmt.Printf("F(200) = %s\n", big.String())
}
```

### Performance Example
```go
package main

import (
    "fmt"
    "time"
    "yourmodule/pkg/fibonacci"
)

func main() {
    n := 40
    
    // Compare different implementations
    start := time.Now()
    result1 := fibonacci.Fibonacci(n)
    fmt.Printf("Iterative: F(%d) = %d, Time: %v\n", 
        n, result1, time.Since(start))
    
    start = time.Now()
    result2 := fibonacci.FibonacciMatrix(n)
    fmt.Printf("Matrix: F(%d) = %d, Time: %v\n", 
        n, result2, time.Since(start))
    
    start = time.Now()
    result3 := fibonacci.FibonacciMemoized(n)
    fmt.Printf("Memoized: F(%d) = %d, Time: %v\n", 
        n, result3, time.Since(start))
}
```

## Educational Tools

### Sequence Generation - `FibonacciSequence(n int) []int`
Generates the first n Fibonacci numbers.

```go
sequence := fibonacci.FibonacciSequence(8)
// Returns: [0, 1, 1, 2, 3, 5, 8, 13]
```

### Validation - `IsValidFibonacci(num int) bool`
Checks if a number belongs to the Fibonacci sequence.

```go
fmt.Println(fibonacci.IsValidFibonacci(13)) // true
fmt.Println(fibonacci.IsValidFibonacci(14)) // false
```

### Index Search - `FibonacciIndex(num int) int`
Finds the index of a Fibonacci number in the sequence.

```go
index := fibonacci.FibonacciIndex(55) // Returns 10
```

### Golden Ratio - `GoldenRatio(n int) float64`
Calculates the golden ratio approximation using F(n+1)/F(n).

```go
ratio := fibonacci.GoldenRatio(20) // Approximately 1.618034
```

### Sequence Sum - `FibonacciSum(n int) int`
Calculates the sum of the first n Fibonacci numbers.

```go
sum := fibonacci.FibonacciSum(6) // 0+1+1+2+3+5 = 12
```

### Visualization - `Visualize(n int) int`
Calculates and displays the calculation process (useful for debugging and education).

```go
result := fibonacci.Visualize(5) // Shows the calculation process
```

## Performance

### Time Complexity
| Implementation | Complexity | Best For |
|---------------|------------|----------|
| Iterative | O(n) | General use (n < 50) |
| Recursive | O(2ⁿ) | Education only (n < 15) |
| Memoized | O(n) | Multiple calls |
| Matrix | O(log n) | Single large values |
| Big Int | O(n) | Very large numbers |

### Benchmarks (approximate)
```
BenchmarkFibonacci-8               30000000        50.0 ns/op
BenchmarkFibonacciMemoized-8      10000000       150.0 ns/op  
BenchmarkFibonacciMatrix-8         5000000       300.0 ns/op
BenchmarkFibonacciBig-8            1000000      1500.0 ns/op
BenchmarkFibonacciRecursive-8           100  15000000.0 ns/op
```

## Use Cases

### 1. Education and Demonstration
- Use `FibonacciRecursive()` to show the basic concept
- Use `Visualize()` to demonstrate the calculation process
- Use `FibonacciSequence()` to show the progression

### 2. Performance Applications
- Use `Fibonacci()` for small to medium values
- Use `FibonacciMatrix()` for single calculations of large values
- Use `FibonacciMemoized()` when calculating multiple values

### 3. Scientific Computing
- Use `FibonacciBig()` for mathematical research with large numbers
- Use `GoldenRatio()` for golden ratio approximations
- Use `FibonacciSum()` for statistical analysis

### 4. Validation and Analysis
- Use `IsValidFibonacci()` to verify if data follows Fibonacci patterns
- Use `FibonacciIndex()` to find positions in the sequence

## Running Tests

```bash
# Run all tests
go test ./pkg/fibonacci

# Run tests with verbose output
go test -v ./pkg/fibonacci

# Run benchmarks
go test -bench=. ./pkg/fibonacci

# Run coverage tests
go test -cover ./pkg/fibonacci
```

## Limitations Notes

1. **Integer Overflow**: Implementations with `int` overflow after F(46) = 1836311903
2. **Recursive Performance**: Recursive implementation is impractical for n > 15
3. **Memory**: Memoized implementation uses O(n) additional memory
4. **Precision**: For very large numbers, always use `FibonacciBig()`

## Mathematical References

- [Fibonacci Sequence - Wikipedia](https://en.wikipedia.org/wiki/Fibonacci_number)
- [The Golden Ratio](https://en.wikipedia.org/wiki/Golden_ratio)
- [Matrix Form of Fibonacci Sequence](https://www.mathsisfun.com/algebra/matrix-fibonacci.html)
