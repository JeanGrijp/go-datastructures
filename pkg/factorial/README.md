# Factorial Package

The `factorial` package provides arbitrary-precision factorial computation using Go's `math/big`.

## Overview

Factorials grow very quickly and overflow built-in integer types for relatively small inputs. This package avoids overflow by returning results as `*big.Int`.

## Features

- Arbitrary-precision factorials with `big.Int`
- Simple and focused API
- Efficient iterative implementation
- Fully covered by tests and examples

## Installation

```bash
go get github.com/JeanGrijp/go-datastructures/pkg/factorial
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/JeanGrijp/go-datastructures/pkg/factorial"
)

func main() {
    result := factorial.Factorial(100)
    fmt.Println("100! =", result)
}
```

## API Reference

### `Factorial(n int) *big.Int`

Computes `n!` using arbitrary-precision arithmetic.

- Input: non-negative integer `n`
- Output: `*big.Int`

Behavior notes:

- `Factorial(0) == 1`
- `Factorial(1) == 1`
- For negative `n`, current behavior returns `1` because the loop does not execute

## Performance

- Time complexity: `O(n)`
- Extra space complexity: `O(1)` excluding result storage

## Examples

### Small Inputs

```go
fmt.Println("5! =", factorial.Factorial(5))
fmt.Println("10! =", factorial.Factorial(10))
```

### Large Inputs

```go
value := factorial.Factorial(1000)
fmt.Printf("1000! has %d digits\n", len(value.String()))
```

## Testing

```bash
go test ./pkg/factorial
```

Run all repository tests:

```bash
go test ./...
```
