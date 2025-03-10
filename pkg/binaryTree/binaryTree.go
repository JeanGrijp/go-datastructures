package binarytree

// Node represents a node in a binary tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree represents a binary tree
type BinaryTree struct {
	Root *Node
}

// NewBinaryTree creates a new binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

// Insert inserts a value into the binary tree
func (t *BinaryTree) Insert(value int) {

}
