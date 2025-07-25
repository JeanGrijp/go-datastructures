# Go Data Structures

A comprehensive collection of data structures and algorithms implemented in Go, including classic implementations and fundamental computer science algorithms.

## 📦 Project Structure

```bash
go-datastructures/
├── cmd/
│   └── main.go                    # Main demo of packages
├── examples/
│   └── farm_problem/
│       └── main.go                # Farm problem example (Euclidean)
├── pkg/
│   ├── binaryTree/
│   │   └── binaryTree.go          # Basic binary tree structure
│   ├── euclidean/                 # → Euclidean Algorithm (GCD/LCM)
│   ├── fatorial/                  # → Factorial calculation with big.Int
│   ├── graph/                     # → Basic graph structure
│   ├── sort/                      # → Sorting algorithms
│   └── stack/                     # → Stack (LIFO)
├── go.mod                         # Go module
└── README.md                      # This file
```

## 🚀 Implemented Algorithms and Data Structures

| Package | Description | Status | Documentation |
|---------|-------------|--------|---------------|
| **[euclidean](pkg/euclidean/)** | Euclidean Algorithm - GCD, LCM, farm problem | ✅ Complete | [📖 README](pkg/euclidean/README.md) |
| **[fatorial](pkg/fatorial/)** | Factorial calculation with big.Int for large numbers | ✅ Complete | [📖 README](pkg/fatorial/README.md) |
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

### LIFO Stack

```go
import "github.com/JeanGrijp/go-datastructures/pkg/stack"

stack := stack.NewStack()
stack.Push(42)
value, _ := stack.Pop() // 42
```

## 📖 Documentation

For detailed information about each algorithm, consult the specific documentation:

- **[Euclidean Algorithm](pkg/euclidean/README.md)** - GCD, LCM, farm problem
- **[Factorial](pkg/fatorial/README.md)** - Calculations with big.Int
- **[Stack](pkg/stack/README.md)** - LIFO implementation
- **[Sorting Algorithms](pkg/sort/README.md)** - QuickSort, MergeSort, etc.

### go doc Commands

```bash
go doc ./pkg/euclidean
go doc ./pkg/fatorial
go doc ./pkg/stack
go doc ./pkg/sort
```

## 🛠️ Installation

### Prerequisites

- Go 1.19 or higher

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

