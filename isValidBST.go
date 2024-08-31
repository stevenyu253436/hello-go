package main

import (
	"math"
)

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isValidBST 是验证二叉树是否为有效的二叉搜索树的主函数
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

// validate 是一个辅助函数，用来递归地验证每个节点是否在给定的范围内
func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// 当前节点的值必须在 (min, max) 范围内
	if node.Val <= min || node.Val >= max {
		return false
	}

	// 递归检查左子树和右子树
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

// 主函数用于测试
func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	result := isValidBST(root)
	if result {
		println("The tree is a valid BST")
	} else {
		println("The tree is NOT a valid BST")
	}
}
