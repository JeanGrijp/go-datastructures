// Package stack provides a generic stack data structure implementation
// using a linked list with LIFO (Last In, First Out) behavior.
package stack

// Node represents a single element in the stack linked list.
// Each node contains a value and a pointer to the next node.
type Node struct {
	Value any   // The value stored in this node (can be any type)
	Next  *Node // Pointer to the next node in the stack
}

// Stack represents a LIFO (Last In, First Out) data structure
// implemented using a singly linked list.
type Stack struct {
	first *Node // Pointer to the top element of the stack
}

// Push adds a new element to the top of the stack.
// The new element becomes the first element and the previous
// first element becomes the second element.
//
// Parameters:
//   - value: The value to be added to the stack (can be any type)
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example:
//
//	stack := &Stack{}
//	stack.Push(10)
//	stack.Push("hello")
//	stack.Push([]int{1, 2, 3})
func (s *Stack) Push(value any) {
	s.first = &Node{Value: value, Next: s.first}
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, it returns nil and false.
//
// Returns:
//   - any: The value that was at the top of the stack
//   - bool: true if an element was successfully popped, false if stack was empty
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example:
//
//	stack := &Stack{}
//	stack.Push(42)
//	value, ok := stack.Pop()
//	if ok {
//		fmt.Println("Popped:", value) // Output: Popped: 42
//	}
func (s *Stack) Pop() (any, bool) {
	if s.first == nil {
		return nil, false
	} else {
		value := s.first.Value
		s.first = s.first.Next
		return value, true
	}
}

// Values returns all elements in the stack as a slice.
// The elements are returned in the order they would be popped
// (top to bottom). If the stack is empty, returns nil.
//
// Returns:
//   - any: A slice containing all values in the stack, or nil if empty
//
// Time complexity: O(n) where n is the number of elements
// Space complexity: O(n) for the returned slice
//
// Example:
//
//	stack := &Stack{}
//	stack.Push(1)
//	stack.Push(2)
//	stack.Push(3)
//	values := stack.Values()
//	fmt.Println(values) // Output: [3 2 1]
func (s *Stack) Values() any {
	if s.first == nil {
		return nil
	}
	var values []any
	for aux := s.first; aux != nil; aux = aux.Next {
		values = append(values, aux.Value)
	}
	return values
}

// IsEmpty checks if the stack is empty.
//
// Returns:
//   - bool: true if the stack is empty, false otherwise
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example:
//
//	stack := &Stack{}
//	fmt.Println(stack.IsEmpty()) // Output: true
//	stack.Push(42)
//	fmt.Println(stack.IsEmpty()) // Output: false
func (s *Stack) IsEmpty() bool {
	return s.first == nil
}

// Peek returns the top element without removing it from the stack.
// If the stack is empty, it returns nil and false.
//
// Returns:
//   - any: The value at the top of the stack
//   - bool: true if there is an element, false if stack is empty
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example:
//
//	stack := &Stack{}
//	stack.Push("hello")
//	value, ok := stack.Peek()
//	if ok {
//		fmt.Println("Top element:", value) // Output: Top element: hello
//	}
func (s *Stack) Peek() (any, bool) {
	if s.first == nil {
		return nil, false
	}
	return s.first.Value, true
}

// Size returns the number of elements in the stack.
//
// Returns:
//   - int: The number of elements in the stack
//
// Time complexity: O(n) where n is the number of elements
// Space complexity: O(1)
//
// Example:
//
//	stack := &Stack{}
//	fmt.Println(stack.Size()) // Output: 0
//	stack.Push(1)
//	stack.Push(2)
//	fmt.Println(stack.Size()) // Output: 2
func (s *Stack) Size() int {
	count := 0
	for aux := s.first; aux != nil; aux = aux.Next {
		count++
	}
	return count
}
