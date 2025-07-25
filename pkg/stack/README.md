# Stack Package

The `stack` package provides a generic stack data structure implementation using a linked list with LIFO (Last In, First Out) behavior.

## Overview

A stack is a linear data structure that follows the LIFO (Last In, First Out) principle. This implementation uses a singly linked list to provide efficient push and pop operations. The stack can store values of any type using Go's `any` interface.

## Features

- **Generic Type Support**: Can store values of any type using `any` interface
- **Efficient Operations**: O(1) push, pop, peek, and isEmpty operations
- **Memory Efficient**: Uses linked list implementation with dynamic memory allocation
- **Safe Operations**: Pop and Peek return boolean indicators for empty stack handling
- **Complete API**: Includes all standard stack operations plus utility methods

## Installation

```bash
go get github.com/JeanGrijp/go-datastructures/pkg/stack
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/JeanGrijp/go-datastructures/pkg/stack"
)

func main() {
    // Create a new stack
    s := &stack.Stack{}
    
    // Push elements
    s.Push(10)
    s.Push("hello")
    s.Push([]int{1, 2, 3})
    
    // Pop elements
    value, ok := s.Pop()
    if ok {
        fmt.Println("Popped:", value) // Output: Popped: [1 2 3]
    }
    
    // Peek at top element
    top, ok := s.Peek()
    if ok {
        fmt.Println("Top:", top) // Output: Top: hello
    }
}
```

### Working with Different Types

```go
// Integer stack
s := &stack.Stack{}
s.Push(1)
s.Push(2)
s.Push(3)

// String stack
stringStack := &stack.Stack{}
stringStack.Push("first")
stringStack.Push("second")
stringStack.Push("third")

// Mixed type stack (though not recommended for type safety)
mixedStack := &stack.Stack{}
mixedStack.Push(42)
mixedStack.Push("hello")
mixedStack.Push(3.14)
```

### Stack Operations

```go
s := &stack.Stack{}

// Check if empty
fmt.Println("Is empty:", s.IsEmpty()) // Output: Is empty: true

// Push elements
s.Push("A")
s.Push("B")
s.Push("C")

// Get size
fmt.Println("Size:", s.Size()) // Output: Size: 3

// Get all values
values := s.Values()
fmt.Println("Values:", values) // Output: Values: [C B A]

// Peek without removing
top, _ := s.Peek()
fmt.Println("Top:", top) // Output: Top: C

// Pop elements
for !s.IsEmpty() {
    value, _ := s.Pop()
    fmt.Println("Popped:", value)
}
// Output: 
// Popped: C
// Popped: B
// Popped: A
```

## API Reference

### Types

#### `Node`

```go
type Node struct {
    Value any   // The value stored in this node
    Next  *Node // Pointer to the next node
}
```

Represents a single element in the stack's linked list.

#### `Stack`

```go
type Stack struct {
    first *Node // Pointer to the top element
}
```

The main stack structure using a singly linked list.

### Methods

#### `Push(value any)`

Adds a new element to the top of the stack.

**Parameters:**

- `value` (any): The value to be added to the stack

**Time Complexity:** O(1)  
**Space Complexity:** O(1)

#### `Pop() (any, bool)`

Removes and returns the top element from the stack.

**Returns:**

- `any`: The value that was at the top of the stack
- `bool`: true if successful, false if stack was empty

**Time Complexity:** O(1)  
**Space Complexity:** O(1)

#### `Peek() (any, bool)`

Returns the top element without removing it from the stack.

**Returns:**

- `any`: The value at the top of the stack
- `bool`: true if there is an element, false if stack is empty

**Time Complexity:** O(1)  
**Space Complexity:** O(1)

#### `IsEmpty() bool`

Checks if the stack is empty.

**Returns:**

- `bool`: true if the stack is empty, false otherwise

**Time Complexity:** O(1)  
**Space Complexity:** O(1)

#### `Size() int`

Returns the number of elements in the stack.

**Returns:**

- `int`: The number of elements in the stack

**Time Complexity:** O(n)  
**Space Complexity:** O(1)

#### `Values() any`

Returns all elements in the stack as a slice in LIFO order.

**Returns:**

- `any`: A slice containing all values, or nil if empty

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

## Performance Characteristics

| Operation | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Push      | O(1)           | O(1)             |
| Pop       | O(1)           | O(1)             |
| Peek      | O(1)           | O(1)             |
| IsEmpty   | O(1)           | O(1)             |
| Size      | O(n)           | O(1)             |
| Values    | O(n)           | O(n)             |

## Examples

### Reverse a String

```go
func reverseString(s string) string {
    stack := &stack.Stack{}
    
    // Push all characters
    for _, char := range s {
        stack.Push(string(char))
    }
    
    // Pop all characters
    var result string
    for !stack.IsEmpty() {
        char, _ := stack.Pop()
        result += char.(string)
    }
    
    return result
}

fmt.Println(reverseString("hello")) // Output: olleh
```

### Check Balanced Parentheses

```go
func isBalanced(expr string) bool {
    stack := &stack.Stack{}
    pairs := map[rune]rune{')': '(', '}': '{', ']': '['}
    
    for _, char := range expr {
        switch char {
        case '(', '{', '[':
            stack.Push(char)
        case ')', '}', ']':
            if stack.IsEmpty() {
                return false
            }
            top, _ := stack.Pop()
            if top.(rune) != pairs[char] {
                return false
            }
        }
    }
    
    return stack.IsEmpty()
}

fmt.Println(isBalanced("({[]})")) // Output: true
fmt.Println(isBalanced("({[})"))  // Output: false
```

### Undo Functionality

```go
type UndoSystem struct {
    actions *stack.Stack
}

func NewUndoSystem() *UndoSystem {
    return &UndoSystem{actions: &stack.Stack{}}
}

func (u *UndoSystem) Execute(action string) {
    u.actions.Push(action)
    fmt.Println("Executed:", action)
}

func (u *UndoSystem) Undo() {
    if action, ok := u.actions.Pop(); ok {
        fmt.Println("Undoing:", action)
    } else {
        fmt.Println("Nothing to undo")
    }
}

// Usage
undo := NewUndoSystem()
undo.Execute("Create file")
undo.Execute("Edit file")
undo.Undo() // Output: Undoing: Edit file
undo.Undo() // Output: Undoing: Create file
```

## Mathematical Background

### Stack Properties

- **LIFO Principle**: Last element pushed is the first to be popped
- **Access**: Only the top element is directly accessible
- **Dynamic Size**: Can grow and shrink during runtime

### Use Cases

- **Function Call Management**: Call stack in programming languages
- **Expression Evaluation**: Parsing mathematical expressions
- **Undo Operations**: Implementing undo functionality in applications
- **Browser History**: Back button functionality
- **Compiler Design**: Syntax parsing and checking

## Memory Management

This implementation uses Go's garbage collector for automatic memory management:

- **Node Creation**: New nodes are allocated when pushing elements
- **Node Cleanup**: Popped nodes become eligible for garbage collection
- **Memory Efficiency**: Only allocates memory for actual elements

## Type Safety Considerations

While this implementation uses `any` for maximum flexibility, consider creating type-specific stacks for better type safety:

```go
type IntStack struct {
    stack *stack.Stack
}

func (is *IntStack) Push(value int) {
    is.stack.Push(value)
}

func (is *IntStack) Pop() (int, bool) {
    if value, ok := is.stack.Pop(); ok {
        return value.(int), true
    }
    return 0, false
}
```

## License

This package is part of the go-datastructures project. See the main repository for license information.

## Contributing

Contributions are welcome! Please see the main repository for contribution guidelines.
