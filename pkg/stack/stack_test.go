package stack

import (
	"fmt"
	"testing"
)

// TestPush tests the Push method
func TestPush(t *testing.T) {
	s := &Stack{}

	// Test pushing different types
	s.Push(42)
	s.Push("hello")
	s.Push([]int{1, 2, 3})

	if s.IsEmpty() {
		t.Error("Stack should not be empty after pushing elements")
	}

	if s.Size() != 3 {
		t.Errorf("Expected size 3, got %d", s.Size())
	}
}

// TestPop tests the Pop method
func TestPop(t *testing.T) {
	s := &Stack{}

	// Test pop from empty stack
	value, ok := s.Pop()
	if ok {
		t.Error("Pop should return false for empty stack")
	}
	if value != nil {
		t.Error("Pop should return nil for empty stack")
	}

	// Test pop from non-empty stack
	s.Push(42)
	s.Push("hello")

	value, ok = s.Pop()
	if !ok {
		t.Error("Pop should return true for non-empty stack")
	}
	if value != "hello" {
		t.Errorf("Expected 'hello', got %v", value)
	}

	value, ok = s.Pop()
	if !ok {
		t.Error("Pop should return true for non-empty stack")
	}
	if value != 42 {
		t.Errorf("Expected 42, got %v", value)
	}

	// Stack should be empty now
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

// TestPeek tests the Peek method
func TestPeek(t *testing.T) {
	s := &Stack{}

	// Test peek on empty stack
	value, ok := s.Peek()
	if ok {
		t.Error("Peek should return false for empty stack")
	}
	if value != nil {
		t.Error("Peek should return nil for empty stack")
	}

	// Test peek on non-empty stack
	s.Push(42)
	s.Push("hello")

	value, ok = s.Peek()
	if !ok {
		t.Error("Peek should return true for non-empty stack")
	}
	if value != "hello" {
		t.Errorf("Expected 'hello', got %v", value)
	}

	// Size should remain the same after peek
	if s.Size() != 2 {
		t.Errorf("Size should remain 2 after peek, got %d", s.Size())
	}
}

// TestIsEmpty tests the IsEmpty method
func TestIsEmpty(t *testing.T) {
	s := &Stack{}

	// Test empty stack
	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	// Test non-empty stack
	s.Push(42)
	if s.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}

	// Test empty after pop
	s.Pop()
	if !s.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

// TestSize tests the Size method
func TestSize(t *testing.T) {
	s := &Stack{}

	// Test empty stack
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}

	// Test growing stack
	for i := 1; i <= 5; i++ {
		s.Push(i)
		if s.Size() != i {
			t.Errorf("Expected size %d, got %d", i, s.Size())
		}
	}

	// Test shrinking stack
	for i := 4; i >= 0; i-- {
		s.Pop()
		if s.Size() != i {
			t.Errorf("Expected size %d, got %d", i, s.Size())
		}
	}
}

// TestValues tests the Values method
func TestValues(t *testing.T) {
	s := &Stack{}

	// Test empty stack
	values := s.Values()
	if values != nil {
		t.Error("Values should return nil for empty stack")
	}

	// Test non-empty stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	values = s.Values()
	if values == nil {
		t.Error("Values should not return nil for non-empty stack")
	}

	valuesSlice := values.([]any)
	expected := []any{3, 2, 1}

	if len(valuesSlice) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(valuesSlice))
	}

	for i, v := range valuesSlice {
		if v != expected[i] {
			t.Errorf("Expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}

// TestLIFOBehavior tests the LIFO (Last In, First Out) behavior
func TestLIFOBehavior(t *testing.T) {
	s := &Stack{}
	input := []string{"first", "second", "third", "fourth"}

	// Push all elements
	for _, item := range input {
		s.Push(item)
	}

	// Pop all elements and verify LIFO order
	for i := len(input) - 1; i >= 0; i-- {
		value, ok := s.Pop()
		if !ok {
			t.Error("Pop should succeed")
		}
		if value != input[i] {
			t.Errorf("Expected %s, got %v", input[i], value)
		}
	}
}

// BenchmarkPush benchmarks the Push operation
func BenchmarkPush(b *testing.B) {
	s := &Stack{}
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}

// BenchmarkPop benchmarks the Pop operation
func BenchmarkPop(b *testing.B) {
	s := &Stack{}
	// Pre-populate the stack
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

// BenchmarkPeek benchmarks the Peek operation
func BenchmarkPeek(b *testing.B) {
	s := &Stack{}
	s.Push(42)

	for i := 0; i < b.N; i++ {
		s.Peek()
	}
}

// ExampleStack demonstrates basic stack usage
func ExampleStack() {
	s := &Stack{}

	// Push elements
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Pop elements
	for !s.IsEmpty() {
		value, _ := s.Pop()
		fmt.Println(value)
	}
	// Output:
	// 3
	// 2
	// 1
}

// ExampleStack_Peek demonstrates peek functionality
func ExampleStack_Peek() {
	s := &Stack{}
	s.Push("hello")
	s.Push("world")

	// Peek at top element
	top, ok := s.Peek()
	if ok {
		fmt.Println("Top element:", top)
	}

	// Size remains the same
	fmt.Println("Size:", s.Size())
	// Output:
	// Top element: world
	// Size: 2
}

// ExampleStack_Values demonstrates getting all values
func ExampleStack_Values() {
	s := &Stack{}
	s.Push("A")
	s.Push("B")
	s.Push("C")

	values := s.Values()
	fmt.Println("Stack contents:", values)
	// Output: Stack contents: [C B A]
}
