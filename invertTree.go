package main

import (
	"fmt"
)

// TreeNode is a definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree function to invert the binary tree
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Recursively invert the left and right subtrees
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// Swap the left and right children
	root.Left = right
	root.Right = left

	return root
}

// Helper function to print the tree in pre-order for testing
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func main() {
	// Create a sample binary tree: [4,2,7,1,3,6,9]
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Println("Original tree pre-order:")
	preOrder(root)
	fmt.Println()

	// Invert the tree
	invertedRoot := invertTree(root)

	fmt.Println("Inverted tree pre-order:")
	preOrder(invertedRoot)
	fmt.Println()
}
