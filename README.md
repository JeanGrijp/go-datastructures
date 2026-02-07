<!-- markdownlint-disable MD041 MD033 MD010 -->
<div align="center">
	<h1>Go Data Structures</h1>
	<p>
		A comprehensive collection of data structures and algorithms implemented in Go, including classic implementations and fundamental computer science algorithms.
	</p>
	<p>
		<img alt="Go Version" src="https://img.shields.io/badge/Go-1.23.2-00ADD8?logo=go&logoColor=white" />
		<img alt="Frameworks" src="https://img.shields.io/badge/Frameworks-Standard%20Library-blue" />
		<img alt="License" src="https://img.shields.io/badge/License-MIT-green.svg" />
	</p>
</div>
<!-- markdownlint-enable MD041 MD033 MD010 -->

## 📦 Project Structure

```bash
go-datastructures/
├── cmd/
│   └── main.go                    # Main demo of packages
├── examples/
│   └── farm_problem/
│       └── main.go                # Farm problem example (Euclidean)
├── pkg/
│   ├── b-tree/
│   │   └── b-tree.go              # B-Tree self-balancing tree
│   ├── binaryTree/
│   │   └── binaryTree.go          # Basic binary tree structure
│   ├── euclidean/                 # → Euclidean Algorithm (GCD/LCM)
│   ├── fatorial/                  # → Factorial calculation with big.Int
│   ├── fibonacci/                 # → Fibonacci algorithms (multiple implementations)
│   ├── graph/                     # → Basic graph structure
│   ├── hashTable/                 # → Hash Table (hash map)
│   ├── sort/                      # → Sorting algorithms
│   └── stack/                     # → Stack (LIFO)
├── go.mod                         # Go module
└── README.md                      # This file
```

## 🧰 Tech Stack & Versions

<!-- markdownlint-disable MD033 MD010 -->
<p align="center">
	<img alt="Language" src="https://img.shields.io/badge/Language-Go%20(1.23.2)-00ADD8?logo=go&logoColor=white" />
	<img alt="Build" src="https://img.shields.io/badge/Build-Go%20Modules-informational" />
	<img alt="Testing" src="https://img.shields.io/badge/Testing-go%20test-9cf" />
	<img alt="StdLib" src="https://img.shields.io/badge/Frameworks-Standard%20Library%20Only-lightgrey" />
</p>
<!-- markdownlint-enable MD033 MD010 -->

Notes:

- This repository uses only the Go Standard Library (no external web frameworks).
- Some packages use big integers via the standard library `math/big`.

## 🚀 Implemented Algorithms and Data Structures

| Package | Description | Status | Documentation |
|---------|-------------|--------|---------------|
| **[b-tree](pkg/b-tree/)** | B-Tree self-balancing search tree for databases | ✅ Complete | [📖 README](pkg/b-tree/README.md) |
| **[euclidean](pkg/euclidean/)** | Euclidean Algorithm - GCD, LCM, farm problem | ✅ Complete | [📖 README](pkg/euclidean/README.md) |
| **[fatorial](pkg/fatorial/)** | Factorial calculation with big.Int for large numbers | ✅ Complete | [📖 README](pkg/fatorial/README.md) |
| **[fibonacci](pkg/fibonacci/)** | Fibonacci sequence - Multiple algorithm implementations | ✅ Complete | [📖 README](pkg/fibonacci/README.md) |
| **[hashTable](pkg/hashTable/)** | Hash Table with separate chaining collision handling | ✅ Complete | [📖 README](pkg/hashTable/README.md) |
| **[stack](pkg/stack/)** | Stack implementation (LIFO) with linked list | ✅ Complete | [📖 README](pkg/stack/README.md) |
| **[sort](pkg/sort/)** | Sorting algorithms (QuickSort, MergeSort, etc.) | ✅ Complete | [📖 README](pkg/sort/README.md) |
| **[binaryTree](pkg/binaryTree/)** | Basic binary tree structure | 🚧 In development | - |
| **[graph](pkg/graph/)** | Graph structures and algorithms | 🚧 Planned | - |

## 🎯 Quick Start

### Run Main Demo

```bash
go run ./cmd/main.go
```

### Practical Example - Farm Problem

```bash
go run ./examples/farm_problem/main.go
```

### Run Tests

```bash
go test ./pkg/...
```

## 📚 Test Coverage

| Package | Coverage | Status |
|---------|----------|--------|
| euclidean | 93.5% | ✅ |
| fatorial | 100% | ✅ |
| fibonacci | 52.7% | ✅ |
| stack | 100% | ✅ |
| sort | 96.3% | ✅ |

## 🛠️ Quick Usage

### Euclidean Algorithm

```go
import "github.com/JeanGrijp/go-datastructures/pkg/euclidean"

// Farm problem: largest square that divides 1680x1050
squareSize := euclidean.LargestSquareSize(1680, 1050) // 210
```

### Large Number Factorial

```go
import "github.com/JeanGrijp/go-datastructures/pkg/fatorial"

// Calculate 3000! (9131 digits)
result := fatorial.Factorial(3000)
```

### Fibonacci Sequence

```go
import "github.com/JeanGrijp/go-datastructures/pkg/fibonacci"

// Multiple algorithm implementations
fib := fibonacci.Fibonacci(50)                    // Iterative O(n)
fibBig := fibonacci.FibonacciBig(1000)           // Big numbers
fibMatrix := fibonacci.FibonacciMatrix(100)      // Matrix O(log n)
sequence := fibonacci.FibonacciSequence(10)      // [0 1 1 2 3 5 8 13 21 34]
```

### LIFO Stack

```go
import "github.com/JeanGrijp/go-datastructures/pkg/stack"

stack := stack.NewStack()
stack.Push(42)
value, _ := stack.Pop() // 42
```

### B-Tree

```go
import btree "github.com/JeanGrijp/go-datastructures/pkg/b-tree"

bt := btree.NewBTree(3)  // minimum degree 3
bt.Insert(10)
bt.Insert(20)
bt.Insert(5)
found := bt.Search(10)  // true
bt.Remove(10)
```

## 📖 Documentation

For detailed information about each algorithm, consult the specific documentation:

- **[B-Tree](pkg/b-tree/README.md)** - Self-balancing search tree
- **[Euclidean Algorithm](pkg/euclidean/README.md)** - GCD, LCM, farm problem
- **[Factorial](pkg/fatorial/README.md)** - Calculations with big.Int
- **[Fibonacci Sequence](pkg/fibonacci/README.md)** - Multiple algorithm implementations
- **[Hash Table](pkg/hashTable/README.md)** - Hash map with collision handling
- **[Stack](pkg/stack/README.md)** - LIFO implementation
- **[Sorting Algorithms](pkg/sort/README.md)** - QuickSort, MergeSort, etc.

### go doc Commands

```bash
go doc ./pkg/b-tree
go doc ./pkg/euclidean
go doc ./pkg/fatorial
go doc ./pkg/fibonacci
go doc ./pkg/hashTable
go doc ./pkg/stack
go doc ./pkg/sort
```

## 🛠️ Installation

### Prerequisites

- Go 1.23+ (module targets Go 1.23.2)

### Clone Repository

```bash
git clone https://github.com/JeanGrijp/go-datastructures.git
cd go-datastructures
```

## 🤝 Contributing

Contributions are welcome! For each new algorithm or data structure:

1. Create the package in `pkg/algorithm_name/`
2. Implement with complete tests
3. Add detailed documentation in the package's README.md
4. Update the table in this main README

## 📄 License

This project is licensed under the MIT License.

## 🔗 References

- [Understanding Algorithms](https://www.manning.com/books/grokking-algorithms) - Source of the farm problem
- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-third-edition)
- [The Art of Computer Programming (Knuth)](https://www-cs-faculty.stanford.edu/~knuth/taocp.html)
- [Official Go Documentation](https://golang.org/doc/)

