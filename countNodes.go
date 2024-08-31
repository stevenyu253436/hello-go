package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if leftHeight == rightHeight {
		// 左子树是满二叉树，节点数为 2^leftHeight - 1，加上根节点和右子树的节点数
		return (1 << leftHeight) + countNodes(root.Right)
	} else {
		// 右子树是满二叉树，节点数为 2^rightHeight - 1，加上根节点和左子树的节点数
		return (1 << rightHeight) + countNodes(root.Left)
	}
}

func getHeight(root *TreeNode) int {
	height := 0
	for root != nil {
		height++
		root = root.Left
	}
	return height
}

func main() {
	// 示例测试
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}

	fmt.Println(countNodes(root)) // 输出: 6
}
