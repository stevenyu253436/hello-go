package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 递归地展平左子树和右子树
	flatten(root.Left)
	flatten(root.Right)

	// 暂存右子树
	rightSubtree := root.Right

	// 将左子树移到右边
	root.Right = root.Left
	root.Left = nil

	// 找到新的右子树的末尾
	current := root
	for current.Right != nil {
		current = current.Right
	}

	// 将原先的右子树接在新的右子树的末尾
	current.Right = rightSubtree
}

// Helper function to print the tree in a flattened form
func printFlattenedTree(root *TreeNode) {
	for root != nil {
		println(root.Val)
		root = root.Right
	}
}

func main() {
	// Create a test tree: [1,2,5,3,4,null,6]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}

	// Flatten the tree
	flatten(root)

	// Print the flattened tree
	printFlattenedTree(root)
}
