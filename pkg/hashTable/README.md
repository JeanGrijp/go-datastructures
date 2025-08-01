# Hash Table Package

This implementation provides a hash table (also known as hash map) data structure with collision handling using separate chaining. A hash table is one of the most important and widely used data structures in computer science, providing fast average-case performance for insertions, deletions, and lookups.

## Table of Contents

- [About Hash Tables](#about-hash-tables)
- [Implementation Details](#implementation-details)
- [Available Operations](#available-operations)
- [Usage Examples](#usage-examples)
- [Performance Analysis](#performance-analysis)
- [Collision Handling](#collision-handling)
- [Use Cases](#use-cases)

## About Hash Tables

A hash table is a data structure that implements an associative array abstract data type, mapping keys to values. It uses a hash function to compute an index into an array of buckets or slots, from which the desired value can be found.

### Key Concepts

- **Hash Function**: Converts keys into array indices
- **Buckets**: Array slots that store key-value pairs
- **Collision**: When two different keys hash to the same index
- **Load Factor**: Ratio of stored elements to total buckets
- **Separate Chaining**: Collision resolution using linked lists/slices

### Mathematical Properties

- **Hash Function**: Uses FNV-1a (Fowler-Noll-Vo) hash algorithm
- **Load Factor**: λ = n/m (where n = elements, m = buckets)
- **Expected Bucket Size**: λ for uniform distribution
- **Collision Probability**: Increases with load factor

## Implementation Details

### Data Structure

```go
type HashTable struct {
    buckets  [][]KeyValuePair // Array of buckets (slices)
    size     int              // Number of stored elements
    capacity int              // Number of buckets
}

type KeyValuePair struct {
    key   string // Key for hashing and lookup
    value any    // Associated value
}
```

### Hash Function

This implementation uses the FNV-1a hash function, which provides:
- Fast computation
- Good distribution properties
- Low collision rates for typical data

### Collision Resolution

Uses **separate chaining** where:
- Each bucket contains a slice of key-value pairs
- Collisions are handled by appending to the bucket's slice
- Search within a bucket is linear

## Available Operations

### 1. Creation - `New(capacity int) *HashTable`

Creates a new hash table with specified capacity.

```go
ht := hashtable.New(16) // Creates hash table with 16 buckets
```

**Complexity**: O(1) time, O(capacity) space

### 2. Insert/Update - `Put(key string, value any) bool`

Inserts a new key-value pair or updates an existing one.

```go
isNew := ht.Put("name", "Alice")    // true (new key)
isNew = ht.Put("name", "Bob")       // false (updated existing)
```

**Complexity**: O(1) average, O(n) worst case

### 3. Lookup - `Get(key string) (any, bool)`

Retrieves the value associated with a key.

```go
value, found := ht.Get("name")
if found {
    fmt.Printf("Name: %v\n", value)
}
```

**Complexity**: O(1) average, O(n) worst case

### 4. Delete - `Delete(key string) bool`

Removes a key-value pair from the hash table.

```go
deleted := ht.Delete("name")
if deleted {
    fmt.Println("Name was removed")
}
```

**Complexity**: O(1) average, O(n) worst case

### 5. Membership Test - `Contains(key string) bool`

Checks if a key exists in the hash table.

```go
if ht.Contains("email") {
    fmt.Println("Email key exists")
}
```

**Complexity**: O(1) average, O(n) worst case

### 6. Size Operations

```go
size := ht.Size()        // Get number of elements
empty := ht.IsEmpty()    // Check if empty
ht.Clear()              // Remove all elements
```

**Complexity**: All O(1) except Clear which is O(capacity)

### 7. Iteration Operations

```go
keys := ht.Keys()       // Get all keys
values := ht.Values()   // Get all values
pairs := ht.GetPairs()  // Get all key-value pairs
```

**Complexity**: O(capacity + size) time, O(size) space

### 8. Analysis Operations

```go
loadFactor := ht.LoadFactor()           // Get load factor
distribution := ht.BucketDistribution() // Analyze bucket sizes
fmt.Println(ht.String())               // Print hash table structure
```

## Usage Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "yourmodule/pkg/hashtable"
)

func main() {
    // Create a new hash table
    ht := hashtable.New(10)
    
    // Insert some data
    ht.Put("name", "Alice")
    ht.Put("age", 30)
    ht.Put("city", "New York")
    
    // Retrieve data
    if name, found := ht.Get("name"); found {
        fmt.Printf("Name: %s\n", name)
    }
    
    // Check existence
    if ht.Contains("age") {
        fmt.Println("Age is stored")
    }
    
    // Update existing value
    ht.Put("age", 31)
    
    // Delete a key
    ht.Delete("city")
    
    // Print statistics
    fmt.Printf("Size: %d\n", ht.Size())
    fmt.Printf("Load Factor: %.2f\n", ht.LoadFactor())
}
```

### Working with Different Data Types

```go
package main

import (
    "fmt"
    "yourmodule/pkg/hashtable"
)

func main() {
    ht := hashtable.New(8)
    
    // Store different types
    ht.Put("string", "Hello World")
    ht.Put("integer", 42)
    ht.Put("float", 3.14159)
    ht.Put("boolean", true)
    ht.Put("slice", []int{1, 2, 3, 4, 5})
    ht.Put("map", map[string]int{"a": 1, "b": 2})
    
    // Retrieve and type assert
    if value, found := ht.Get("slice"); found {
        if slice, ok := value.([]int); ok {
            fmt.Printf("Slice: %v\n", slice)
        }
    }
    
    // Iterate over all pairs
    pairs := ht.GetPairs()
    for _, pair := range pairs {
        fmt.Printf("Key: %s, Value: %v, Type: %T\n", 
            pair.Key(), pair.Value(), pair.Value())
    }
}
```

### Performance Analysis

```go
package main

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
    "yourmodule/pkg/hashtable"
)

func main() {
    ht := hashtable.New(100)
    
    // Insert many items
    start := time.Now()
    for i := 0; i < 1000; i++ {
        key := "key" + strconv.Itoa(i)
        ht.Put(key, rand.Intn(1000))
    }
    insertTime := time.Since(start)
    
    // Analyze distribution
    distribution := ht.BucketDistribution()
    fmt.Printf("Bucket Distribution:\n")
    for size, count := range distribution {
        fmt.Printf("  %d buckets have %d items\n", count, size)
    }
    
    fmt.Printf("Insert Time: %v\n", insertTime)
    fmt.Printf("Load Factor: %.2f\n", ht.LoadFactor())
    fmt.Printf("Total Size: %d\n", ht.Size())
    
    // Test lookup performance
    start = time.Now()
    found := 0
    for i := 0; i < 1000; i++ {
        key := "key" + strconv.Itoa(i)
        if ht.Contains(key) {
            found++
        }
    }
    lookupTime := time.Since(start)
    
    fmt.Printf("Lookup Time: %v\n", lookupTime)
    fmt.Printf("Found: %d/1000\n", found)
}
```

### Building a Cache

```go
package main

import (
    "fmt"
    "yourmodule/pkg/hashtable"
)

type Cache struct {
    data *hashtable.HashTable
    hits int
    misses int
}

func NewCache(capacity int) *Cache {
    return &Cache{
        data: hashtable.New(capacity),
    }
}

func (c *Cache) Get(key string) (any, bool) {
    value, found := c.data.Get(key)
    if found {
        c.hits++
    } else {
        c.misses++
    }
    return value, found
}

func (c *Cache) Put(key string, value any) {
    c.data.Put(key, value)
}

func (c *Cache) Stats() (int, int, float64) {
    total := c.hits + c.misses
    hitRate := 0.0
    if total > 0 {
        hitRate = float64(c.hits) / float64(total)
    }
    return c.hits, c.misses, hitRate
}

func main() {
    cache := NewCache(50)
    
    // Simulate cache usage
    cache.Put("user:1", "Alice")
    cache.Put("user:2", "Bob")
    
    // Cache hits
    cache.Get("user:1")
    cache.Get("user:2")
    
    // Cache miss
    cache.Get("user:3")
    
    hits, misses, hitRate := cache.Stats()
    fmt.Printf("Hits: %d, Misses: %d, Hit Rate: %.2f%%\n", 
        hits, misses, hitRate*100)
}
```

## Performance Analysis

### Time Complexity

| Operation | Average Case | Worst Case | Best Case |
|-----------|-------------|------------|-----------|
| Put | O(1) | O(n) | O(1) |
| Get | O(1) | O(n) | O(1) |
| Delete | O(1) | O(n) | O(1) |
| Contains | O(1) | O(n) | O(1) |
| Size | O(1) | O(1) | O(1) |
| Clear | O(m) | O(m) | O(m) |
| Keys/Values | O(m + n) | O(m + n) | O(m + n) |

*Where n = number of elements, m = number of buckets*

### Space Complexity

- **Storage**: O(n) for the key-value pairs
- **Overhead**: O(m) for the bucket array
- **Total**: O(n + m)

### Load Factor Impact

```
Load Factor Range | Performance | Recommendation
0.0 - 0.5        | Excellent   | Optimal for critical applications
0.5 - 0.75       | Good        | Recommended for most use cases
0.75 - 1.0       | Fair        | Consider resizing
> 1.0            | Poor        | Resize recommended
```

### Benchmarks (Approximate)

```
BenchmarkPut-8           5000000    300 ns/op
BenchmarkGet-8          10000000    150 ns/op
BenchmarkDelete-8        5000000    250 ns/op
BenchmarkContains-8     10000000    150 ns/op
```

## Collision Handling

### Separate Chaining

This implementation uses separate chaining where:

1. **Each bucket is a slice** of key-value pairs
2. **Collisions append** new items to the bucket
3. **Search is linear** within each bucket
4. **Memory efficient** compared to open addressing

### Advantages

- **Simple implementation**
- **Handles high load factors well**
- **Deletion is straightforward**
- **No clustering problems**

### Disadvantages

- **Extra memory for pointers/slices**
- **Cache performance may be worse**
- **Linear search within buckets**

## Use Cases

### 1. Database Indexing

```go
// Create an index for fast lookups
userIndex := hashtable.New(1000)
userIndex.Put("alice@email.com", "user_id_123")
userIndex.Put("bob@email.com", "user_id_456")

// Fast user lookup
if userID, found := userIndex.Get("alice@email.com"); found {
    fmt.Printf("User ID: %s\n", userID)
}
```

### 2. Caching System

```go
// Simple web cache
cache := hashtable.New(500)
cache.Put("/api/users", userData)
cache.Put("/api/products", productData)

// Fast response serving
if cachedData, found := cache.Get("/api/users"); found {
    return cachedData // Serve from cache
}
```

### 3. Configuration Management

```go
// Application configuration
config := hashtable.New(50)
config.Put("database.host", "localhost")
config.Put("database.port", 5432)
config.Put("api.rate_limit", 1000)

// Quick config access
if host, found := config.Get("database.host"); found {
    connectToDatabase(host.(string))
}
```

### 4. Frequency Counting

```go
// Word frequency counter
counter := hashtable.New(1000)
words := []string{"hello", "world", "hello", "go", "world", "hello"}

for _, word := range words {
    if count, found := counter.Get(word); found {
        counter.Put(word, count.(int)+1)
    } else {
        counter.Put(word, 1)
    }
}

// Print frequencies
pairs := counter.GetPairs()
for _, pair := range pairs {
    fmt.Printf("%s: %d\n", pair.Key(), pair.Value())
}
```

### 5. Symbol Tables

```go
// Compiler symbol table
symbols := hashtable.New(200)
symbols.Put("variable_x", VariableInfo{Type: "int", Scope: "local"})
symbols.Put("function_foo", FunctionInfo{ReturnType: "string", Params: []string{"int", "bool"}})

// Symbol lookup during compilation
if symbol, found := symbols.Get("variable_x"); found {
    // Process variable information
    processVariable(symbol.(VariableInfo))
}
```

## Testing and Validation

### Running Tests

```bash
# Run all tests
go test ./pkg/hashtable

# Run tests with verbose output
go test -v ./pkg/hashtable

# Run benchmarks
go test -bench=. ./pkg/hashtable

# Run coverage tests
go test -cover ./pkg/hashtable
```

### Performance Testing

```bash
# Run performance benchmarks
go test -bench=BenchmarkPut ./pkg/hashtable
go test -bench=BenchmarkGet ./pkg/hashtable
go test -bench=BenchmarkDelete ./pkg/hashtable

# Memory profiling
go test -bench=. -memprofile=mem.prof ./pkg/hashtable
```

## Limitations and Considerations

### 1. Memory Usage

- **Overhead**: Each bucket requires slice metadata
- **Fragmentation**: Multiple small slices may cause fragmentation
- **Growth**: Slices may over-allocate during growth

### 2. Performance Characteristics

- **Worst Case**: O(n) when all keys hash to the same bucket
- **Load Factor**: Performance degrades with high load factors
- **Hash Quality**: Performance depends on hash function distribution

### 3. Key Limitations

- **Key Type**: Currently only supports string keys
- **Hash Collisions**: Birthday paradox affects collision probability
- **No Resizing**: Fixed capacity may become a bottleneck

### 4. Concurrency

- **Not Thread-Safe**: Requires external synchronization for concurrent access
- **Race Conditions**: Simultaneous reads and writes can cause data races

## Mathematical Analysis

### Expected Bucket Size

For a hash table with `n` elements and `m` buckets:
- Expected bucket size: `n/m` (load factor)
- Standard deviation: `√(n/m)`

### Collision Probability

For `n` elements and `m` buckets:
- Probability of collision: `1 - (m!/((m-n)!m^n))`
- Approximation for large m: `1 - e^(-n²/2m)`

### Load Factor Recommendations

- **α < 0.75**: Good performance
- **α > 1.0**: Consider resizing
- **Optimal range**: 0.5 ≤ α ≤ 0.75

## References

- [Hash Table - Wikipedia](https://en.wikipedia.org/wiki/Hash_table)
- [FNV Hash Function](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function)
- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-third-edition)
- [The Art of Computer Programming, Volume 3](https://www-cs-faculty.stanford.edu/~knuth/taocp.html)
