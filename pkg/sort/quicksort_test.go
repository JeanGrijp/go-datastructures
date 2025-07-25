package sort

import (
	"fmt"
	"testing"
)

// TestQuickSortSlice tests the convenience function for sorting entire slices
func TestQuickSortSlice(t *testing.T) {
	array := []int{5, 3, 8, 6, 2, 7, 1, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	QuickSortSlice(array)

	for i, v := range array {
		if v != expected[i] {
			t.Errorf("QuickSortSlice failed: got %v, expected %v", array, expected)
			break
		}
	}
}

// TestQuickSort tests the main QuickSort function
func TestQuickSort(t *testing.T) {
	array := []int{5, 3, 8, 6, 2, 7, 1, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	QuickSort(array, 0, len(array)-1)

	for i, v := range array {
		if v != expected[i] {
			t.Errorf("QuickSort failed: got %v, expected %v", array, expected)
			break
		}
	}
}

// TestIsSorted tests the IsSorted function
func TestIsSorted(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5}
	if !IsSorted(sorted) {
		t.Error("IsSorted should return true for sorted array")
	}

	unsorted := []int{5, 3, 8, 1}
	if IsSorted(unsorted) {
		t.Error("IsSorted should return false for unsorted array")
	}
}

// ExampleQuickSortSlice demonstrates basic usage
func ExampleQuickSortSlice() {
	array := []int{5, 3, 8, 6, 2, 7, 1, 4}
	fmt.Println("Original:", array)

	QuickSortSlice(array)
	fmt.Println("Sorted:", array)
	// Output:
	// Original: [5 3 8 6 2 7 1 4]
	// Sorted: [1 2 3 4 5 6 7 8]
}
