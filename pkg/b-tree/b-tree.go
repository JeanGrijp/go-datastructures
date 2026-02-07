// Package btree implements a B-Tree data structure, a self-balancing search tree
// that maintains sorted data and allows searches, sequential access, insertions,
// and deletions in logarithmic time.
//
// A B-Tree of order t (minimum degree) has the following properties:
//   - Every node has at most 2t-1 keys
//   - Every node (except root) has at least t-1 keys
//   - All leaves appear at the same level
//   - A non-leaf node with k keys has k+1 children
//
// B-Trees are commonly used in databases and file systems where large blocks
// of data are read and written. They minimize disk I/O operations by keeping
// the tree height low.
//
// Time complexity for all operations: O(log n)
// Space complexity: O(n)
package btree

// BTreeNode represents a single node in the B-Tree.
// Each node contains keys in sorted order and pointers to child nodes.
type BTreeNode struct {
	leaf   bool         // Indicates whether this node is a leaf (has no children)
	keys   []int        // Slice of keys stored in this node, always kept in sorted order
	childs []*BTreeNode // Slice of pointers to child nodes (empty if leaf is true)
}

// BTree represents a B-Tree data structure with a specified minimum degree.
// The minimum degree t determines the range of keys each node can hold.
type BTree struct {
	root *BTreeNode // Pointer to the root node of the tree (nil if tree is empty)
	t    int        // Minimum degree: each node can have [t-1, 2t-1] keys (except root)
}

// NewBTree creates and returns a new empty B-Tree with the specified minimum degree t.
// The minimum degree determines the capacity of each node:
//   - Each node can hold between t-1 and 2t-1 keys
//   - Each non-leaf node can have between t and 2t children
//
// Parameters:
//   - t: The minimum degree of the B-Tree (must be >= 2 for a valid B-Tree)
//
// Returns:
//   - *BTree: A pointer to the newly created empty B-Tree
//
// Time complexity: O(1)
// Space complexity: O(1)
//
// Example:
//
//	bt := btree.NewBTree(3)  // Creates a B-Tree with minimum degree 3
//	bt.Insert(10)
//	bt.Insert(20)
//	bt.Insert(5)
func NewBTree(t int) *BTree {
	return &BTree{t: t}
}

// Insert adds a new key to the B-Tree while maintaining all B-Tree properties.
// If the key already exists, it will be inserted as a duplicate.
//
// The insertion process:
//  1. If the tree is empty, create a new root with the key
//  2. If the root is full (has 2t-1 keys), split it and create a new root
//  3. Insert the key into the appropriate position using insertNonFull
//
// Parameters:
//   - k: The integer key to insert into the tree
//
// Time complexity: O(t * log_t(n)) where t is the minimum degree and n is the number of keys
// Space complexity: O(log_t(n)) for the recursion stack
//
// Example:
//
//	bt := btree.NewBTree(2)
//	bt.Insert(10)
//	bt.Insert(20)
//	bt.Insert(5)
//	bt.Insert(15)
//	// Tree now contains: 5, 10, 15, 20
func (bt *BTree) Insert(k int) {
	if bt.root == nil {
		bt.root = &BTreeNode{leaf: true, keys: []int{k}}
		return
	}

	// If the root is full, the tree grows in height
	if len(bt.root.keys) == 2*bt.t-1 {
		oldRoot := bt.root
		bt.root = &BTreeNode{leaf: false}
		bt.root.childs = append(bt.root.childs, oldRoot)
		bt.splitChild(bt.root, 0, oldRoot)
	}
	bt.insertNonFull(bt.root, k)
}

// splitChild splits a full child node into two nodes during insertion.
// This is a key operation that maintains the B-Tree balance property.
//
// The split process:
//  1. Create a new node to hold the right half of the full node's keys
//  2. Move the t-1 rightmost keys to the new node
//  3. Move corresponding children if the node is not a leaf
//  4. Promote the middle key to the parent node
//  5. Insert the new node as a child of the parent
//
// Parameters:
//   - parent: The parent node that will receive the promoted middle key
//   - i: The index in parent's children where fullNode is located
//   - fullNode: The full node (has 2t-1 keys) to be split
//
// Time complexity: O(t) where t is the minimum degree
// Space complexity: O(t) for the new node
func (bt *BTree) splitChild(parent *BTreeNode, i int, fullNode *BTreeNode) {
	t := bt.t
	newNode := &BTreeNode{leaf: fullNode.leaf}

	// The new node receives the t-1 rightmost keys from the full node
	newNode.keys = append(newNode.keys, fullNode.keys[t:]...)
	if !fullNode.leaf {
		newNode.childs = append(newNode.childs, fullNode.childs[t:]...)
	}

	// The middle key is promoted to the parent
	middleKey := fullNode.keys[t-1]
	fullNode.keys = fullNode.keys[:t-1]
	if !fullNode.leaf {
		fullNode.childs = fullNode.childs[:t]
	}

	// Insert the new node and the middle key into the parent
	parent.keys = append(parent.keys, 0)
	copy(parent.keys[i+1:], parent.keys[i:])
	parent.keys[i] = middleKey

	parent.childs = append(parent.childs, nil)
	copy(parent.childs[i+2:], parent.childs[i+1:])
	parent.childs[i+1] = newNode
}

// insertNonFull inserts a key into a node that is guaranteed to be non-full.
// This is a helper function called by Insert after ensuring the node has room.
//
// The insertion process:
//  1. If the node is a leaf, insert the key in sorted order
//  2. If the node is internal, find the appropriate child to descend into
//  3. If that child is full, split it first, then recurse
//
// Parameters:
//   - node: The non-full node where the key should be inserted
//   - k: The key to insert
//
// Time complexity: O(t * log_t(n))
// Space complexity: O(log_t(n)) for recursion
func (bt *BTree) insertNonFull(node *BTreeNode, k int) {
	i := len(node.keys) - 1

	if node.leaf {
		// Insert the key at the correct position in the leaf
		node.keys = append(node.keys, 0)
		for i >= 0 && node.keys[i] > k {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = k
	} else {
		// Find the child where the key should be inserted
		for i >= 0 && node.keys[i] > k {
			i--
		}
		i++
		if len(node.childs[i].keys) == 2*bt.t-1 {
			bt.splitChild(node, i, node.childs[i])
			if node.keys[i] < k {
				i++
			}
		}
		bt.insertNonFull(node.childs[i], k)
	}
}

// Search looks for a key in the subtree rooted at this node.
// It traverses the tree from this node downward, comparing the key
// with stored keys to determine the search path.
//
// The search process:
//  1. Find the first key greater than or equal to k
//  2. If the key is found, return true
//  3. If this is a leaf node and key not found, return false
//  4. Otherwise, recursively search in the appropriate child
//
// Parameters:
//   - k: The key to search for
//
// Returns:
//   - bool: true if the key exists in the subtree, false otherwise
//
// Time complexity: O(t * log_t(n)) where t is the minimum degree
// Space complexity: O(log_t(n)) for recursion stack
func (node *BTreeNode) search(k int) bool {
	i := 0
	for i < len(node.keys) && k > node.keys[i] {
		i++
	}

	if i < len(node.keys) && node.keys[i] == k {
		return true
	}

	if node.leaf {
		return false
	}

	return node.childs[i].search(k)
}

// Search looks for a key in the B-Tree.
// Returns true if the key exists in the tree, false otherwise.
//
// Parameters:
//   - k: The key to search for
//
// Returns:
//   - bool: true if the key exists in the tree, false otherwise
//
// Time complexity: O(t * log_t(n)) where t is the minimum degree
// Space complexity: O(log_t(n)) for recursion stack
//
// Example:
//
//	bt := btree.NewBTree(2)
//	bt.Insert(10)
//	bt.Insert(20)
//	found := bt.Search(10)  // returns true
//	found = bt.Search(15)   // returns false
func (bt *BTree) Search(k int) bool {
	if bt.root == nil {
		return false
	}
	return bt.root.search(k)
}

// Remove deletes a key from the B-Tree while maintaining all B-Tree properties.
// If the key does not exist in the tree, the tree remains unchanged.
//
// The removal process handles several cases:
//  1. Key is in a leaf node: simply remove it
//  2. Key is in an internal node: replace with predecessor/successor and delete
//  3. After deletion, if root becomes empty, shrink the tree height
//
// Parameters:
//   - k: The key to remove from the tree
//
// Time complexity: O(t * log_t(n)) where t is the minimum degree
// Space complexity: O(log_t(n)) for recursion stack
//
// Example:
//
//	bt := btree.NewBTree(2)
//	bt.Insert(10)
//	bt.Insert(20)
//	bt.Insert(5)
//	bt.Remove(10)  // Removes 10 from the tree
//	// Tree now contains: 5, 20
func (bt *BTree) Remove(k int) {
	if bt.root == nil {
		return
	}

	bt.removeFromNode(bt.root, k)

	if len(bt.root.keys) == 0 {
		if !bt.root.leaf {
			bt.root = bt.root.childs[0]
		} else {
			bt.root = nil
		}
	}
}

// removeFromNode removes a key from the subtree rooted at the given node.
// This is a helper function that handles the recursive deletion logic.
//
// The removal handles three cases:
//  1. Key found in a leaf: directly remove the key
//  2. Key found in an internal node: delegate to removeFromNonLeaf
//  3. Key not in this node: ensure child has enough keys, then recurse
//
// Parameters:
//   - node: The root of the subtree to remove from
//   - k: The key to remove
//
// Time complexity: O(t * log_t(n))
// Space complexity: O(log_t(n)) for recursion
func (bt *BTree) removeFromNode(node *BTreeNode, k int) {
	idx := 0
	for idx < len(node.keys) && node.keys[idx] < k {
		idx++
	}

	if idx < len(node.keys) && node.keys[idx] == k {
		if node.leaf {
			node.keys = append(node.keys[:idx], node.keys[idx+1:]...)
		} else {
			bt.removeFromNonLeaf(node, idx)
		}
	} else {
		if node.leaf {
			return
		}

		lastChild := idx == len(node.keys)
		if len(node.childs[idx].keys) < bt.t {
			bt.fill(node, idx)
		}

		if lastChild && idx > len(node.keys) {
			bt.removeFromNode(node.childs[idx-1], k)
		} else {
			bt.removeFromNode(node.childs[idx], k)
		}
	}
}

// removeFromNonLeaf removes a key from an internal (non-leaf) node.
// This requires finding a replacement key to maintain B-Tree structure.
//
// Three cases are handled:
//  1. Left child has >= t keys: replace with predecessor and delete predecessor
//  2. Right child has >= t keys: replace with successor and delete successor
//  3. Both children have t-1 keys: merge children and delete from merged node
//
// Parameters:
//   - node: The internal node containing the key to remove
//   - idx: The index of the key to remove in node.keys
//
// Time complexity: O(t * log_t(n))
// Space complexity: O(log_t(n)) for recursion
func (bt *BTree) removeFromNonLeaf(node *BTreeNode, idx int) {
	k := node.keys[idx]

	if len(node.childs[idx].keys) >= bt.t {
		pred := bt.getPredecessor(node, idx)
		node.keys[idx] = pred
		bt.removeFromNode(node.childs[idx], pred)
	} else if len(node.childs[idx+1].keys) >= bt.t {
		succ := bt.getSuccessor(node, idx)
		node.keys[idx] = succ
		bt.removeFromNode(node.childs[idx+1], succ)
	} else {
		bt.merge(node, idx)
		bt.removeFromNode(node.childs[idx], k)
	}
}

// getPredecessor finds the predecessor of the key at index idx in node.
// The predecessor is the largest key in the left subtree of keys[idx].
// It traverses to the rightmost leaf of the left child.
//
// Parameters:
//   - node: The node containing the key whose predecessor is needed
//   - idx: The index of the key in node.keys
//
// Returns:
//   - int: The predecessor key (largest key smaller than node.keys[idx])
//
// Time complexity: O(log_t(n))
// Space complexity: O(1)
func (bt *BTree) getPredecessor(node *BTreeNode, idx int) int {
	cur := node.childs[idx]
	for !cur.leaf {
		cur = cur.childs[len(cur.childs)-1]
	}
	return cur.keys[len(cur.keys)-1]
}

// getSuccessor finds the successor of the key at index idx in node.
// The successor is the smallest key in the right subtree of keys[idx].
// It traverses to the leftmost leaf of the right child.
//
// Parameters:
//   - node: The node containing the key whose successor is needed
//   - idx: The index of the key in node.keys
//
// Returns:
//   - int: The successor key (smallest key larger than node.keys[idx])
//
// Time complexity: O(log_t(n))
// Space complexity: O(1)
func (bt *BTree) getSuccessor(node *BTreeNode, idx int) int {
	cur := node.childs[idx+1]
	for !cur.leaf {
		cur = cur.childs[0]
	}
	return cur.keys[0]
}

// fill ensures that the child at index idx has at least t keys.
// This is called before descending into a child during deletion
// to guarantee we can remove a key without violating B-Tree properties.
//
// The fill operation tries in order:
//  1. Borrow from the previous sibling if it has >= t keys
//  2. Borrow from the next sibling if it has >= t keys
//  3. Merge with a sibling if neither has spare keys
//
// Parameters:
//   - node: The parent node
//   - idx: The index of the child that needs more keys
//
// Time complexity: O(t)
// Space complexity: O(1)
func (bt *BTree) fill(node *BTreeNode, idx int) {
	if idx != 0 && len(node.childs[idx-1].keys) >= bt.t {
		bt.borrowFromPrev(node, idx)
	} else if idx != len(node.keys) && len(node.childs[idx+1].keys) >= bt.t {
		bt.borrowFromNext(node, idx)
	} else {
		if idx != len(node.keys) {
			bt.merge(node, idx)
		} else {
			bt.merge(node, idx-1)
		}
	}
}

// borrowFromPrev borrows a key from the previous (left) sibling.
// This is used when a child has fewer than t keys and needs to be filled.
//
// The borrow process:
//  1. Move the separator key from parent down to the child
//  2. Move the last key from sibling up to parent as new separator
//  3. If not a leaf, move the last child pointer from sibling to child
//
// Parameters:
//   - node: The parent node containing the separator key
//   - idx: The index of the child that needs a key
//
// Time complexity: O(t) due to slice prepend operations
// Space complexity: O(t) for the new slices
func (bt *BTree) borrowFromPrev(node *BTreeNode, idx int) {
	child := node.childs[idx]
	sibling := node.childs[idx-1]

	child.keys = append([]int{node.keys[idx-1]}, child.keys...)
	if !child.leaf {
		child.childs = append([]*BTreeNode{sibling.childs[len(sibling.childs)-1]}, child.childs...)
		sibling.childs = sibling.childs[:len(sibling.childs)-1]
	}

	node.keys[idx-1] = sibling.keys[len(sibling.keys)-1]
	sibling.keys = sibling.keys[:len(sibling.keys)-1]
}

// borrowFromNext borrows a key from the next (right) sibling.
// This is used when a child has fewer than t keys and needs to be filled.
//
// The borrow process:
//  1. Move the separator key from parent down to the child
//  2. Move the first key from sibling up to parent as new separator
//  3. If not a leaf, move the first child pointer from sibling to child
//
// Parameters:
//   - node: The parent node containing the separator key
//   - idx: The index of the child that needs a key
//
// Time complexity: O(t) due to slice operations
// Space complexity: O(1)
func (bt *BTree) borrowFromNext(node *BTreeNode, idx int) {
	child := node.childs[idx]
	sibling := node.childs[idx+1]

	child.keys = append(child.keys, node.keys[idx])
	if !child.leaf {
		child.childs = append(child.childs, sibling.childs[0])
		sibling.childs = sibling.childs[1:]
	}

	node.keys[idx] = sibling.keys[0]
	sibling.keys = sibling.keys[1:]
}

// merge combines two sibling nodes into one, absorbing the separator key from parent.
// This is used when both siblings have exactly t-1 keys and a key needs to be deleted.
//
// The merge process:
//  1. Move the separator key from parent down to the left child
//  2. Append all keys from right sibling to left child
//  3. If not leaves, append all child pointers from right sibling
//  4. Remove the separator key and right child pointer from parent
//
// After merge, the left child will have 2t-1 keys (maximum capacity).
//
// Parameters:
//   - node: The parent node containing the separator key
//   - idx: The index of the separator key (also index of left child)
//
// Time complexity: O(t)
// Space complexity: O(1)
func (bt *BTree) merge(node *BTreeNode, idx int) {
	child := node.childs[idx]
	sibling := node.childs[idx+1]

	child.keys = append(child.keys, node.keys[idx])
	child.keys = append(child.keys, sibling.keys...)

	if !child.leaf {
		child.childs = append(child.childs, sibling.childs...)
	}

	node.keys = append(node.keys[:idx], node.keys[idx+1:]...)
	node.childs = append(node.childs[:idx+1], node.childs[idx+2:]...)
}
