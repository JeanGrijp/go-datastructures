# Factorial Package

The `fatorial` package provides functions to calculate factorial of integers using arbitrary precision arithmetic.

## Overview

This package solves the integer overflow problem that occurs when calculating factorials of large numbers using standard integer types. By using Go's `math/big` package, it can calculate factorials of very large numbers limited only by available memory.

## Features

- **Arbitrary Precision**: Uses `big.Int` to handle very large numbers without overflow
- **Bilingual Support**: Provides both English (`Factorial`) and Portuguese (`Fatorial`) function names
- **Efficient Implementation**: Iterative approach with O(n) time complexity
- **Memory Safe**: No risk of integer overflow or data loss

## Installation

```bash
go get github.com/JeanGrijp/go-datastructures/pkg/fatorial
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/JeanGrijp/go-datastructures/pkg/fatorial"
)

func main() {
    // Calculate factorial of 100
    result := fatorial.Factorial(100)
    fmt.Println("100! =", result)
    
    // Using Portuguese function name
    resultado := fatorial.Fatorial(50)
    fmt.Println("50! =", resultado)
}
```

### Working with Results

```go
// Convert to string for display
result := fatorial.Factorial(300)
fmt.Printf("Length of 300!: %d digits\n", len(result.String()))

// Check if result fits in int64
if result.IsInt64() {
    fmt.Println("Result fits in int64:", result.Int64())
} else {
    fmt.Println("Result is too large for int64")
}
```

## API Reference

### Functions

#### `Factorial(n int) *big.Int`

Calculates the factorial of n (n!) using arbitrary precision arithmetic.

**Parameters:**

- `n` (int): A non-negative integer

**Returns:**

- `*big.Int`: The factorial of n

**Example:**

```go
result := fatorial.Factorial(10)  // Returns 3628800
```

#### `Fatorial(n int) *big.Int`

Portuguese version of the Factorial function. Provides the same functionality with Portuguese naming for compatibility.

**Parameters:**

- `n` (int): A non-negative integer

**Returns:**

- `*big.Int`: The factorial of n

## Performance Characteristics

- **Time Complexity**: O(n) - Linear time relative to input size
- **Space Complexity**: O(1) - Constant space (excluding result storage)
- **Memory Usage**: Result size grows exponentially with input

## Examples

### Small Numbers

```go
fmt.Println("5! =", fatorial.Factorial(5))    // Output: 5! = 120
fmt.Println("10! =", fatorial.Factorial(10))  // Output: 10! = 3628800
```

### Large Numbers

```go
// This would overflow with standard int types
result := fatorial.Factorial(100)
fmt.Printf("100! has %d digits\n", len(result.String()))
// Output: 100! has 158 digits

// Even larger numbers
huge := fatorial.Factorial(1000)
fmt.Printf("1000! has %d digits\n", len(huge.String()))
// Output: 1000! has 2568 digits
```

### Comparison with Standard Types

```go
// This will overflow and give incorrect results
func standardFactorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

// Demonstrating overflow
fmt.Println("Standard 30!:", standardFactorial(30))     // Incorrect due to overflow
fmt.Println("Big.Int 30!:", fatorial.Factorial(30))     // Correct result
```

## Mathematical Background

The factorial function is defined as:

- n! = n × (n-1) × (n-2) × ... × 2 × 1
- 0! = 1 (by convention)

### Growth Rate

Factorials grow extremely rapidly:

- 10! = 3,628,800
- 20! = 2,432,902,008,176,640,000
- 30! = 265,252,859,812,191,058,636,308,480,000,000

## Why Use This Package?

### Integer Overflow Problem

Standard integer types in Go have limits:

- `int32`: max value ≈ 2.1 × 10^9 (13! already overflows)
- `int64`: max value ≈ 9.2 × 10^18 (21! already overflows)

### Solution Benefits

- ✅ No overflow - can calculate factorial of any reasonable size
- ✅ Exact precision - no floating-point approximation errors
- ✅ Easy to use - same interface as standard arithmetic
- ✅ Type safe - compile-time guarantees

## License

This package is part of the go-datastructures project. See the main repository for license information.

## Contributing

Contributions are welcome! Please see the main repository for contribution guidelines.
