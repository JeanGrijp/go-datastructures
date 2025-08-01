// Package hashtable implements a hash table (hash map) data structure with collision handling.
// A hash table is a data structure that implements an associative array abstract data type,
// a structure that can map keys to values. It uses a hash function to compute an index
// into an array of buckets or slots, from which the desired value can be found.
//
// This implementation uses separate chaining to handle collisions, where each bucket
// contains a slice of key-value pairs. This provides good performance characteristics
// with O(1) average case for insert, delete, and lookup operations.
package hashtable

import (
	"fmt"
	"hash/fnv"
)

// HashTable represents a hash table data structure using separate chaining for collision resolution.
// It maintains an array of buckets where each bucket is a slice of KeyValuePair items.
type HashTable struct {
	buckets  [][]KeyValuePair // Array of buckets, each bucket is a slice of key-value pairs
	size     int              // Current number of key-value pairs stored
	capacity int              // Number of buckets in the hash table
}

// KeyValuePair represents a single key-value pair stored in the hash table.
type KeyValuePair struct {
	key   string // The key used for hashing and lookup
	value any    // The value associated with the key
}

// New creates and returns a new HashTable with the specified initial capacity.
// If capacity is 0 or negative, it defaults to 16.
//
// Parameters:
//   - capacity: The initial number of buckets in the hash table
//
// Returns:
//   - *HashTable: A pointer to the newly created hash table
//
// Time complexity: O(1)
// Space complexity: O(capacity)
//
// Example usage:
//
//	ht := hashtable.New(10)
//	// Creates a hash table with 10 buckets
func New(capacity int) *HashTable {
	if capacity <= 0 {
		capacity = 16 // Default capacity
	}

	return &HashTable{
		buckets:  make([][]KeyValuePair, capacity),
		size:     0,
		capacity: capacity,
	}
}

// hash computes the hash value for a given key using FNV-1a hash function.
// This is an internal method used to determine which bucket a key should be placed in.
//
// Parameters:
//   - key: The string key to hash
//
// Returns:
//   - int: The bucket index (hash value modulo capacity)
//
// Time complexity: O(len(key))
// Space complexity: O(1)
func (ht *HashTable) hash(key string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return int(hasher.Sum32()) % ht.capacity
}

// Put inserts or updates a key-value pair in the hash table.
// If the key already exists, its value is updated. If the key is new, it's added.
//
// Parameters:
//   - key: The key to insert or update
//   - value: The value to associate with the key
//
// Returns:
//   - bool: true if a new key was added, false if an existing key was updated
//
// Time complexity: O(1) average case, O(n) worst case (where n is the number of items in the bucket)
// Space complexity: O(1)
//
// Example usage:
//
//	ht.Put("name", "John")
//	ht.Put("age", 30)
//	isNew := ht.Put("name", "Jane") // Updates existing key, returns false
func (ht *HashTable) Put(key string, value any) bool {
	index := ht.hash(key)
	bucket := ht.buckets[index]

	// Check if key already exists in the bucket
	for i, pair := range bucket {
		if pair.key == key {
			// Update existing key
			ht.buckets[index][i].value = value
			return false
		}
	}

	// Add new key-value pair
	ht.buckets[index] = append(bucket, KeyValuePair{key: key, value: value})
	ht.size++
	return true
}

// Get retrieves the value associated with a given key.
//
// Parameters:
//   - key: The key to look up
//
// Returns:
//   - any: The value associated with the key (nil if key doesn't exist)
//   - bool: true if the key was found, false otherwise
//
// Time complexity: O(1) average case, O(n) worst case (where n is the number of items in the bucket)
// Space complexity: O(1)
//
// Example usage:
//
//	value, found := ht.Get("name")
//	if found {
//		fmt.Printf("Name: %v\n", value)
//	}
func (ht *HashTable) Get(key string) (any, bool) {
	index := ht.hash(key)
	bucket := ht.buckets[index]

	for _, pair := range bucket {
		if pair.key == key {
			return pair.value, true
		}
	}

	return nil, false
}

// Delete removes a key-value pair from the hash table.
//
// Parameters:
//   - key: The key to remove
//
// Returns:
//   - bool: true if the key was found and removed, false if the key didn't exist
//
// Time complexity: O(1) average case, O(n) worst case (where n is the number of items in the bucket)
// Space complexity: O(1)
//
// Example usage:
//
//	deleted := ht.Delete("age")
//	if deleted {
//		fmt.Println("Age was removed")
//	}
func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)
	bucket := ht.buckets[index]

	for i, pair := range bucket {
		if pair.key == key {
			// Remove the pair by slicing
			ht.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			ht.size--
			return true
		}
	}

	return false
}

// Contains checks if a key exists in the hash table.
//
// Parameters:
//   - key: The key to check for
//
// Returns:
//   - bool: true if the key exists, false otherwise
//
// Time complexity: O(1) average case, O(n) worst case (where n is the number of items in the bucket)
// Space complexity: O(1)
//
// Example usage:
//
//	if ht.Contains("email") {
//		fmt.Println("Email key exists")
//	}
func (ht *HashTable) Contains(key string) bool {
	_, found := ht.Get(key)
	return found
}

// Size returns the number of key-value pairs currently stored in the hash table.
//
// Returns:
//   - int: The number of key-value pairs
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example usage:
//
//	count := ht.Size()
//	fmt.Printf("Hash table contains %d items\n", count)
func (ht *HashTable) Size() int {
	return ht.size
}

// IsEmpty checks if the hash table contains no key-value pairs.
//
// Returns:
//   - bool: true if the hash table is empty, false otherwise
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example usage:
//
//	if ht.IsEmpty() {
//		fmt.Println("Hash table is empty")
//	}
func (ht *HashTable) IsEmpty() bool {
	return ht.size == 0
}

// Clear removes all key-value pairs from the hash table.
//
// Time complexity: O(capacity)
// Space complexity: O(1)
//
// Example usage:
//
//	ht.Clear()
//	fmt.Printf("Hash table cleared, size is now: %d\n", ht.Size())
func (ht *HashTable) Clear() {
	for i := range ht.buckets {
		ht.buckets[i] = nil
	}
	ht.size = 0
}

// Keys returns a slice containing all keys in the hash table.
// The order of keys is not guaranteed.
//
// Returns:
//   - []string: A slice containing all keys
//
// Time complexity: O(capacity + size)
// Space complexity: O(size)
//
// Example usage:
//
//	keys := ht.Keys()
//	fmt.Printf("All keys: %v\n", keys)
func (ht *HashTable) Keys() []string {
	keys := make([]string, 0, ht.size)

	for _, bucket := range ht.buckets {
		for _, pair := range bucket {
			keys = append(keys, pair.key)
		}
	}

	return keys
}

// Values returns a slice containing all values in the hash table.
// The order of values is not guaranteed and corresponds to the order of Keys().
//
// Returns:
//   - []any: A slice containing all values
//
// Time complexity: O(capacity + size)
// Space complexity: O(size)
//
// Example usage:
//
//	values := ht.Values()
//	fmt.Printf("All values: %v\n", values)
func (ht *HashTable) Values() []any {
	values := make([]any, 0, ht.size)

	for _, bucket := range ht.buckets {
		for _, pair := range bucket {
			values = append(values, pair.value)
		}
	}

	return values
}

// GetPairs returns a slice containing all key-value pairs in the hash table.
// The order of pairs is not guaranteed.
//
// Returns:
//   - []KeyValuePair: A slice containing all key-value pairs
//
// Time complexity: O(capacity + size)
// Space complexity: O(size)
//
// Example usage:
//
//	pairs := ht.GetPairs()
//	for _, pair := range pairs {
//		fmt.Printf("Key: %s, Value: %v\n", pair.Key(), pair.Value())
//	}
func (ht *HashTable) GetPairs() []KeyValuePair {
	pairs := make([]KeyValuePair, 0, ht.size)

	for _, bucket := range ht.buckets {
		pairs = append(pairs, bucket...)
	}

	return pairs
} // LoadFactor calculates the load factor of the hash table.
// Load factor = number of elements / number of buckets.
// A higher load factor means more collisions are likely.
//
// Returns:
//   - float64: The current load factor
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example usage:
//
//	loadFactor := ht.LoadFactor()
//	if loadFactor > 0.75 {
//		fmt.Println("Consider resizing the hash table")
//	}
func (ht *HashTable) LoadFactor() float64 {
	if ht.capacity == 0 {
		return 0
	}
	return float64(ht.size) / float64(ht.capacity)
}

// BucketDistribution returns information about how items are distributed across buckets.
// This is useful for analyzing the performance characteristics of the hash function.
//
// Returns:
//   - map[int]int: A map where key is bucket size and value is count of buckets with that size
//
// Time complexity: O(capacity)
// Space complexity: O(capacity) worst case
//
// Example usage:
//
//	distribution := ht.BucketDistribution()
//	for bucketSize, count := range distribution {
//		fmt.Printf("%d buckets have %d items\n", count, bucketSize)
//	}
func (ht *HashTable) BucketDistribution() map[int]int {
	distribution := make(map[int]int)

	for _, bucket := range ht.buckets {
		bucketSize := len(bucket)
		distribution[bucketSize]++
	}

	return distribution
}

// String returns a string representation of the hash table.
// This is useful for debugging and visualization.
//
// Returns:
//   - string: A string representation of the hash table
//
// Time complexity: O(capacity + size)
// Space complexity: O(size)
//
// Example usage:
//
//	fmt.Println(ht.String())
func (ht *HashTable) String() string {
	result := fmt.Sprintf("HashTable{size: %d, capacity: %d, loadFactor: %.2f}\n",
		ht.size, ht.capacity, ht.LoadFactor())

	for i, bucket := range ht.buckets {
		if len(bucket) > 0 {
			result += fmt.Sprintf("Bucket %d: ", i)
			for j, pair := range bucket {
				if j > 0 {
					result += " -> "
				}
				result += fmt.Sprintf("[%s: %v]", pair.key, pair.value)
			}
			result += "\n"
		}
	}

	return result
}

// Key returns the key of a KeyValuePair.
// This is a getter method for accessing the private key field.
//
// Returns:
//   - string: The key
func (kvp KeyValuePair) Key() string {
	return kvp.key
}

// Value returns the value of a KeyValuePair.
// This is a getter method for accessing the private value field.
//
// Returns:
//   - any: The value
func (kvp KeyValuePair) Value() any {
	return kvp.value
}
