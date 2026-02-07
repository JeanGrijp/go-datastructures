# B-Tree Package

This implementation provides a B-Tree data structure, a self-balancing search tree that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. B-Trees are widely used in databases and file systems due to their ability to minimize disk I/O operations.

## Table of Contents

- [About B-Trees](#about-b-trees)
- [Implementation Details](#implementation-details)
- [Available Operations](#available-operations)
- [Usage Examples](#usage-examples)
- [Performance Analysis](#performance-analysis)
- [Visual Representation](#visual-representation)
- [Use Cases](#use-cases)

## About B-Trees

A B-Tree is a self-balancing tree data structure that maintains sorted data and allows for efficient insertion, deletion, and search operations. Unlike binary search trees, B-Tree nodes can have more than two children, making them particularly well-suited for storage systems that read and write large blocks of data.

### Key Concepts

- **Minimum Degree (t)**: Determines the range of keys in each node
- **Node Capacity**: Each node can have between `t-1` and `2t-1` keys
- **Children**: A non-leaf node with `k` keys has exactly `k+1` children
- **Balance**: All leaves appear at the same level (perfectly balanced)
- **Sorted Keys**: Keys within each node are always in sorted order

### B-Tree Properties

For a B-Tree with minimum degree `t`:

1. Every node has at most `2t - 1` keys
2. Every node (except root) has at least `t - 1` keys
3. The root has at least 1 key (if tree is non-empty)
4. All leaves are at the same depth
5. A non-leaf node with `k` keys has `k + 1` children

### Mathematical Properties

- **Height**: O(log_t(n)) where n is the number of keys
- **Maximum Keys per Node**: 2t - 1
- **Minimum Keys per Node**: t - 1 (except root)
- **Branching Factor**: Between t and 2t

## Implementation Details

### Data Structure

```go
// BTreeNode represents a single node in the B-Tree
type BTreeNode struct {
    leaf   bool         // Indicates whether this node is a leaf
    keys   []int        // Slice of keys stored in sorted order
    childs []*BTreeNode // Slice of pointers to child nodes
}

// BTree represents the B-Tree data structure
type BTree struct {
    root *BTreeNode // Pointer to the root node
    t    int        // Minimum degree of the tree
}
```

### Node Structure Visualization

```
         [key1 | key2 | key3]
        /   |      |      \
    child0 child1 child2 child3
```

Each key `k[i]` satisfies:
- All keys in `child[i]` are less than `k[i]`
- All keys in `child[i+1]` are greater than `k[i]`

## Available Operations

### 1. Creation - `NewBTree(t int) *BTree`

Creates a new empty B-Tree with the specified minimum degree.

```go
bt := btree.NewBTree(3) // Creates B-Tree with minimum degree 3
```

**Complexity**: O(1) time, O(1) space

### 2. Insert - `Insert(k int)`

Inserts a new key into the B-Tree while maintaining all B-Tree properties.

```go
bt := btree.NewBTree(2)
bt.Insert(10)
bt.Insert(20)
bt.Insert(5)
bt.Insert(15)
// Tree now contains: 5, 10, 15, 20
```

**Complexity**: O(t × log_t(n)) time, O(log_t(n)) space

### 3. Search - `Search(k int) bool`

Searches for a key in the B-Tree, returning true if found.

```go
bt := btree.NewBTree(2)
bt.Insert(10)
bt.Insert(20)

found := bt.Search(10)  // true
found = bt.Search(15)   // false
```

**Complexity**: O(t × log_t(n)) time, O(log_t(n)) space (for recursion)

### 4. Remove - `Remove(k int)`

Removes a key from the B-Tree while maintaining all B-Tree properties.

```go
bt := btree.NewBTree(2)
bt.Insert(10)
bt.Insert(20)
bt.Insert(5)
bt.Remove(10)
// Tree now contains: 5, 20
```

**Complexity**: O(t × log_t(n)) time, O(log_t(n)) space

## Usage Examples

### Basic Usage

```go
package main

import (
    "fmt"
    btree "github.com/JeanGrijp/go-datastructures/pkg/b-tree"
)

func main() {
    // Create a B-Tree with minimum degree 3
    bt := btree.NewBTree(3)
    
    // Insert keys
    keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
    for _, k := range keys {
        bt.Insert(k)
    }
    
    // Search for keys
    if bt.Search(12) {
        fmt.Println("Key 12 found!")
    }
    
    if !bt.Search(15) {
        fmt.Println("Key 15 not found!")
    }
    
    // Remove a key
    bt.Remove(6)
}
```

### Building an Index

```go
package main

import (
    btree "github.com/JeanGrijp/go-datastructures/pkg/b-tree"
)

func main() {
    // Create an index for database records
    // Using minimum degree 50 for efficient disk access
    index := btree.NewBTree(50)
    
    // Index some record IDs
    recordIDs := []int{1001, 2045, 3089, 4023, 5067, 6011}
    for _, id := range recordIDs {
        index.Insert(id)
    }
    
    // Check if a record exists
    exists := index.Search(3089) // true
    
    // Remove a deleted record from index
    index.Remove(2045)
}
```

### Sequential Insertions

```go
package main

import (
    btree "github.com/JeanGrijp/go-datastructures/pkg/b-tree"
)

func main() {
    bt := btree.NewBTree(3)
    
    // Insert 1000 sequential keys
    // B-Tree handles this efficiently due to splits
    for i := 1; i <= 1000; i++ {
        bt.Insert(i)
    }
    
    // Tree remains balanced with O(log n) height
}
```

## Performance Analysis

### Time Complexity

| Operation | Average Case | Worst Case |
|-----------|--------------|------------|
| Search    | O(log n)     | O(log n)   |
| Insert    | O(log n)     | O(log n)   |
| Delete    | O(log n)     | O(log n)   |

*Note: The base of the logarithm is t (minimum degree), making operations faster than binary trees for large t.*

### Space Complexity

| Aspect | Complexity |
|--------|------------|
| Storage | O(n) |
| Search Stack | O(log_t(n)) |
| Insert Stack | O(log_t(n)) |
| Delete Stack | O(log_t(n)) |

### Comparison with Other Trees

| Feature | B-Tree | Binary Search Tree | AVL Tree |
|---------|--------|-------------------|----------|
| Height | O(log_t(n)) | O(n) worst | O(log n) |
| Disk I/O | Optimized | Poor | Moderate |
| Node Size | Variable (t to 2t-1 keys) | 1 key | 1 key |
| Balance | Always | Not guaranteed | Always |
| Best For | Databases, File Systems | In-memory | In-memory |

## Visual Representation

### Example: B-Tree with t=2 (2-3-4 Tree)

After inserting keys: 10, 20, 30, 40, 50, 60, 70

```
Initial insertions (10, 20, 30):
        [10 | 20 | 30]

After inserting 40 (split occurs):
           [20]
          /    \
     [10]      [30 | 40]

After inserting 50, 60, 70:
           [20 | 40]
          /    |    \
     [10]  [30]  [50 | 60 | 70]

After another split:
              [20 | 40 | 60]
            /    |    |    \
       [10]  [30]  [50]  [70]
```

### Split Operation

When a node becomes full (2t-1 keys), it splits:

```
Before split (node has 2t-1 = 3 keys):
    Parent: [...X...]
              |
    Full:   [A | B | C]

After split (middle key B moves up):
    Parent: [...X | B ...]
              |     |
    Left:   [A]   [C] :Right
```

## Use Cases

### 1. Database Indexing
B-Trees are the foundation of most database index structures (e.g., MySQL InnoDB, PostgreSQL).

```go
// Primary key index
primaryIndex := btree.NewBTree(100) // High degree for disk blocks

// Insert record IDs
primaryIndex.Insert(recordID)

// Fast lookup
exists := primaryIndex.root.Search(recordID)
```

### 2. File System Directories
File systems like NTFS, HFS+, and ext4 use B-Trees for directory indexing.

### 3. Key-Value Stores
Many NoSQL databases use B-Tree variants for sorted key storage.

### 4. Range Queries
B-Trees efficiently support range queries due to sorted key storage.

### 5. In-Memory Caching
With smaller degree values, B-Trees can be used for sorted in-memory caches.

## Internal Operations

### Insertion Algorithm

1. If tree is empty, create root with the new key
2. If root is full, split it and create new root
3. Traverse down, splitting full nodes proactively
4. Insert key in the appropriate leaf

### Deletion Algorithm

1. Find the key to delete
2. If in leaf, simply remove it
3. If in internal node:
   - Replace with predecessor/successor
   - Or merge children if needed
4. Ensure minimum key requirement via borrowing or merging

### Key Helper Operations

- **splitChild**: Divides a full node into two
- **merge**: Combines two sibling nodes
- **borrowFromPrev/Next**: Redistributes keys from siblings
- **fill**: Ensures a child has enough keys before descent

## API Reference

```go
// Create a new B-Tree
func NewBTree(t int) *BTree

// Insert a key
func (bt *BTree) Insert(k int)

// Search for a key
func (bt *BTree) Search(k int) bool

// Remove a key
func (bt *BTree) Remove(k int)
```

## References

- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms-third-edition) - Chapter 18: B-Trees
- [The Art of Computer Programming, Volume 3 (Knuth)](https://www-cs-faculty.stanford.edu/~knuth/taocp.html)
- [B-Trees - Wikipedia](https://en.wikipedia.org/wiki/B-tree)
- [Database Internals (Alex Petrov)](https://www.databass.dev/) - B-Tree implementations in databases
