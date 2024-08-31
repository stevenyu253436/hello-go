package main

import "math"

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// getMinimumDifference 返回二叉搜索树中任意两个不同节点值之间的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt32

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中序遍历左子树
		inorder(node.Left)

		// 计算当前节点与前一个节点值的差值，并更新最小差值
		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node

		// 中序遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

// 主函数用于测试
func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}

	result := getMinimumDifference(root)
	println("Minimum difference:", result) // 输出: 1
}
