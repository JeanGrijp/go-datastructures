// Package sort provides efficient sorting algorithms for integer slices.
// This package implements various sorting algorithms including QuickSort
// with different partitioning strategies and optimizations.
package sort

// swap swaps two elements in the array at positions i and j.
// This is a helper function used by the partitioning algorithm.
//
// Parameters:
//   - array: The slice to modify
//   - i: Index of the first element to swap
//   - j: Index of the second element to swap
//
// Time complexity: O(1)
// Space complexity: O(1)
func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// partition partitions the array using Hoare's partitioning scheme.
// This function rearranges the array so that elements less than or equal
// to the pivot are on the left side, and elements greater than the pivot
// are on the right side.
//
// The Hoare partition scheme is more efficient than the Lomuto scheme
// as it performs fewer swaps on average.
//
// Parameters:
//   - array: The slice to partition
//   - left: The starting index of the partition range
//   - right: The ending index of the partition range
//
// Returns:
//   - int: The final position of the partition (all elements to the left
//     are <= pivot, all elements to the right are >= pivot)
//
// Time complexity: O(n) where n = right - left + 1
// Space complexity: O(1)
func partition(array []int, left, right int) int {
	pivot := array[left]
	i := left - 1
	j := right + 1

	for {
		for {
			j--
			if array[j] <= pivot {
				break
			}
		}
		for {
			i++
			if array[i] >= pivot {
				break
			}
		}
		if i < j {
			swap(array, i, j)
		} else {
			return j
		}
	}
}

// QuickSort implements the QuickSort algorithm recursively using Hoare partitioning.
// QuickSort is a divide-and-conquer algorithm that works by selecting a 'pivot'
// element and partitioning the array around it, then recursively sorting the
// sub-arrays.
//
// This implementation uses the first element as the pivot and Hoare's
// partitioning scheme for better performance.
//
// Parameters:
//   - array: The slice to be sorted (modified in-place)
//   - left: The starting index of the range to sort
//   - right: The ending index of the range to sort
//
// Time complexity:
//   - Best case: O(n log n) - when pivot divides array evenly
//   - Average case: O(n log n) - expected performance
//   - Worst case: O(n²) - when pivot is always the smallest/largest element
//
// Space complexity: O(log n) - due to recursion stack in average case
//
//	O(n) - in worst case due to unbalanced partitions
//
// Example usage:
//
//	array := []int{5, 3, 8, 6, 2, 7, 1, 4}
//	sort.QuickSort(array, 0, len(array)-1)
//	fmt.Println(array) // Output: [1 2 3 4 5 6 7 8]
func QuickSort(array []int, left, right int) {
	if left < right {
		p := partition(array, left, right)
		QuickSort(array, left, p)
		QuickSort(array, p+1, right)
	}
}

// QuickSortSlice sorts an entire integer slice using QuickSort algorithm.
// This is a convenience function that wraps the main QuickSort function
// to sort the entire slice without needing to specify indices.
//
// Parameters:
//   - array: The slice to be sorted (modified in-place)
//
// Example usage:
//
//	array := []int{5, 3, 8, 6, 2}
//	sort.QuickSortSlice(array)
//	fmt.Println(array) // Output: [2 3 5 6 8]
func QuickSortSlice(array []int) {
	if len(array) > 1 {
		QuickSort(array, 0, len(array)-1)
	}
}

// IsSorted checks if an integer slice is sorted in ascending order.
// This function can be used to verify the result of sorting algorithms
// or to check if sorting is needed.
//
// Parameters:
//   - array: The slice to check
//
// Returns:
//   - bool: true if the slice is sorted in ascending order, false otherwise
//
// Time complexity: O(n)
// Space complexity: O(1)
//
// Example usage:
//
//	array := []int{1, 2, 3, 4, 5}
//	fmt.Println(sort.IsSorted(array)) // Output: true
//
//	array2 := []int{5, 3, 8, 1}
//	fmt.Println(sort.IsSorted(array2)) // Output: false
func IsSorted(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			return false
		}
	}
	return true
}
