# Sort Package

The `sort` package provides efficient sorting algorithms for integer slices, with a focus on the QuickSort algorithm implementation using Hoare partitioning.

## Overview

This package implements the QuickSort algorithm, one of the most efficient comparison-based sorting algorithms. QuickSort is a divide-and-conquer algorithm that works by selecting a 'pivot' element and partitioning the array around it, then recursively sorting the sub-arrays.

The implementation uses Hoare's partitioning scheme, which is more efficient than Lomuto partitioning as it performs fewer swaps on average.

## Features

- **QuickSort Algorithm**: Efficient O(n log n) average-case sorting
- **Hoare Partitioning**: More efficient partitioning scheme with fewer swaps
- **In-place Sorting**: Sorts arrays without additional memory allocation
- **Flexible Interface**: Sort entire slices or specific ranges
- **Utility Functions**: Helper functions for validation and convenience
- **Well Documented**: Comprehensive documentation with complexity analysis

## Installation

```bash
go get github.com/JeanGrijp/go-datastructures/pkg/sort
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/JeanGrijp/go-datastructures/pkg/sort"
)

func main() {
    // Sort an entire slice
    array := []int{5, 3, 8, 6, 2, 7, 1, 4}
    sort.QuickSortSlice(array)
    fmt.Println("Sorted:", array) // Output: Sorted: [1 2 3 4 5 6 7 8]
    
    // Sort a specific range
    array2 := []int{9, 5, 3, 8, 6, 2, 7, 1, 4, 0}
    sort.QuickSort(array2, 2, 7) // Sort indices 2 to 7
    fmt.Println("Partially sorted:", array2) // Output: [9 5 1 2 3 6 7 8 4 0]
}
```

### Working with Different Data Sets

```go
// Small arrays
small := []int{3, 1, 2}
sort.QuickSortSlice(small)
fmt.Println(small) // Output: [1 2 3]

// Large arrays
large := make([]int, 10000)
for i := range large {
    large[i] = 10000 - i // Reverse order (worst case)
}
sort.QuickSortSlice(large)
fmt.Println("First 5:", large[:5]) // Output: First 5: [1 2 3 4 5]

// Already sorted arrays
sorted := []int{1, 2, 3, 4, 5}
fmt.Println("Is sorted:", sort.IsSorted(sorted)) // Output: Is sorted: true
sort.QuickSortSlice(sorted) // Still works efficiently
```

### Validation and Testing

```go
// Check if sorting is needed
data := []int{5, 2, 8, 1, 9}
if !sort.IsSorted(data) {
    fmt.Println("Array needs sorting")
    sort.QuickSortSlice(data)
    fmt.Println("Sorted:", data)
}

// Verify sorting worked
if sort.IsSorted(data) {
    fmt.Println("Sorting successful!")
}
```

## API Reference

### Functions

#### `QuickSort(array []int, left, right int)`

Sorts a portion of an integer slice using the QuickSort algorithm.

**Parameters:**

- `array` ([]int): The slice to be sorted (modified in-place)
- `left` (int): The starting index of the range to sort
- `right` (int): The ending index of the range to sort

**Time Complexity:**
- Best case: O(n log n)
- Average case: O(n log n)
- Worst case: O(n²)

**Space Complexity:** O(log n) average, O(n) worst case

#### `QuickSortSlice(array []int)`

Convenience function to sort an entire integer slice.

**Parameters:**

- `array` ([]int): The slice to be sorted (modified in-place)

**Example:**

```go
data := []int{5, 3, 8, 6, 2}
sort.QuickSortSlice(data)
// data is now [2, 3, 5, 6, 8]
```

#### `IsSorted(array []int) bool`

Checks if an integer slice is sorted in ascending order.

**Parameters:**

- `array` ([]int): The slice to check

**Returns:**

- `bool`: true if sorted, false otherwise

**Time Complexity:** O(n)  
**Space Complexity:** O(1)

## Algorithm Details

### QuickSort Overview

QuickSort is a divide-and-conquer algorithm that works as follows:

1. **Choose a pivot**: Select an element from the array as the pivot
2. **Partition**: Rearrange array so elements smaller than pivot come before it, and elements greater come after
3. **Recursively sort**: Apply QuickSort to the sub-arrays on either side of the pivot

### Hoare Partitioning Scheme

This implementation uses Hoare's partitioning scheme:

- **Pivot Selection**: Uses the first element as pivot
- **Two Pointers**: Start from both ends and move toward each other
- **Efficient**: Performs fewer swaps compared to Lomuto partitioning
- **Stable Performance**: Better performance on arrays with many duplicate elements

### Partitioning Process

```
Initial:  [5, 3, 8, 6, 2, 7, 1, 4]
Pivot: 5

Step 1: Find elements to swap
- From right: find element ≤ 5 → 4
- From left: find element ≥ 5 → 8
- Swap: [5, 3, 4, 6, 2, 7, 1, 8]

Step 2: Continue until pointers meet
- Result: [1, 3, 4, 2, 5, 7, 6, 8]
- Partition point: index 4
```

## Performance Characteristics

| Case | Time Complexity | Description |
|------|----------------|-------------|
| **Best** | O(n log n) | Pivot always divides array evenly |
| **Average** | O(n log n) | Expected performance with random data |
| **Worst** | O(n²) | Pivot is always smallest/largest (sorted arrays) |

### Space Complexity

- **Average case**: O(log n) - recursion depth
- **Worst case**: O(n) - unbalanced partitions lead to deep recursion

### Comparison with Other Algorithms

| Algorithm | Best | Average | Worst | Space | Stable |
|-----------|------|---------|-------|-------|--------|
| QuickSort | O(n log n) | O(n log n) | O(n²) | O(log n) | No |
| MergeSort | O(n log n) | O(n log n) | O(n log n) | O(n) | Yes |
| HeapSort | O(n log n) | O(n log n) | O(n log n) | O(1) | No |
| BubbleSort | O(n) | O(n²) | O(n²) | O(1) | Yes |

## Examples

### Basic Sorting

```go
func ExampleBasicSort() {
    data := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Original:", data)
    
    sort.QuickSortSlice(data)
    fmt.Println("Sorted:", data)
    // Output:
    // Original: [64 34 25 12 22 11 90]
    // Sorted: [11 12 22 25 34 64 90]
}
```

### Performance Testing

```go
func ExamplePerformanceTest() {
    // Best case: Random data
    random := []int{5, 2, 8, 1, 9, 3}
    start := time.Now()
    sort.QuickSortSlice(random)
    fmt.Printf("Random data: %v\n", time.Since(start))
    
    // Worst case: Already sorted
    sorted := []int{1, 2, 3, 4, 5, 6}
    start = time.Now()
    sort.QuickSortSlice(sorted)
    fmt.Printf("Sorted data: %v\n", time.Since(start))
}
```

### Partial Sorting

```go
func ExamplePartialSort() {
    data := []int{9, 5, 3, 8, 6, 2, 7, 1, 4, 0}
    fmt.Println("Original:", data)
    
    // Sort only the middle portion (indices 2-7)
    sort.QuickSort(data, 2, 7)
    fmt.Println("Partially sorted:", data)
    // Output:
    // Original: [9 5 3 8 6 2 7 1 4 0]
    // Partially sorted: [9 5 1 2 3 6 7 8 4 0]
}
```

### Duplicate Elements

```go
func ExampleDuplicates() {
    data := []int{5, 2, 8, 2, 1, 5, 3, 5}
    fmt.Println("With duplicates:", data)
    
    sort.QuickSortSlice(data)
    fmt.Println("Sorted:", data)
    // Output:
    // With duplicates: [5 2 8 2 1 5 3 5]
    // Sorted: [1 2 2 3 5 5 5 8]
}
```

## Optimization Considerations

### When to Use QuickSort

**Good for:**
- Large datasets where average-case performance matters
- When memory usage should be minimized
- When in-place sorting is required
- Random or unsorted data

**Consider alternatives for:**
- Small arrays (insertion sort might be faster)
- Nearly sorted data (insertion sort or Tim sort)
- When guaranteed O(n log n) is required (merge sort or heap sort)
- When stability is required (merge sort)

### Performance Tips

1. **Pivot Selection**: For production use, consider median-of-three pivot selection
2. **Small Arrays**: Switch to insertion sort for small subarrays (< 10 elements)
3. **Duplicate Elements**: Consider 3-way partitioning for many duplicates
4. **Stack Overflow**: Use iterative version for very large arrays

### Hybrid Approaches

```go
// Example: Switch to insertion sort for small arrays
func HybridQuickSort(array []int, left, right int) {
    if right - left < 10 {
        insertionSort(array, left, right)
    } else {
        // Use QuickSort for larger arrays
        sort.QuickSort(array, left, right)
    }
}
```

## Mathematical Background

### Recurrence Relation

The time complexity of QuickSort can be expressed as:

- **Best/Average case**: T(n) = 2T(n/2) + O(n) = O(n log n)
- **Worst case**: T(n) = T(n-1) + O(n) = O(n²)

### Probability Analysis

- Probability of worst case with random pivot: very low
- Expected number of comparisons: ~1.39 n log n
- Expected number of swaps: ~0.33 n log n

## Implementation Notes

### Memory Usage

- **In-place**: Sorts without additional arrays
- **Recursion**: Uses call stack for recursion
- **Cache-friendly**: Good locality of reference

### Stability

QuickSort is **not stable** - equal elements may change relative order. For stable sorting, consider merge sort.

## License

This package is part of the go-datastructures project. See the main repository for license information.

## Contributing

Contributions are welcome! Please see the main repository for contribution guidelines.
