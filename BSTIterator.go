package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

// Constructor initializes the iterator with the root node.
func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.pushLeft(root)
	return iterator
}

// Next returns the next smallest number.
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.pushLeft(node.Right)
	return node.Val
}

// HasNext returns whether we have a next smallest number.
func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

// Helper function to push all left children of a node onto the stack.
func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

// Example usage
func main() {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}

	iterator := Constructor(root)
	fmt.Println(iterator.Next())    // return 3
	fmt.Println(iterator.Next())    // return 7
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 9
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 15
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 20
	fmt.Println(iterator.HasNext()) // return false
}
