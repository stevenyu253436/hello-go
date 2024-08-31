package main

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左子树和右子树的最大贡献值
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 计算通过当前节点的路径和
		priceNewPath := node.Val + leftGain + rightGain

		// 更新最大路径和
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 示例 1
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println(maxPathSum(root1)) // 输出: 6

	// 示例 2
	root2 := &TreeNode{Val: -10}
	root2.Left = &TreeNode{Val: 9}
	root2.Right = &TreeNode{Val: 20}
	root2.Right.Left = &TreeNode{Val: 15}
	root2.Right.Right = &TreeNode{Val: 7}
	fmt.Println(maxPathSum(root2)) // 输出: 42
}
