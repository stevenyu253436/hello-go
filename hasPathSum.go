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

func hasPathSum(root *TreeNode, targetSum int) bool {
	// 如果根节点为空，返回false
	if root == nil {
		return false
	}

	// 如果到达叶子节点，检查是否满足路径和为目标和
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// 递归检查左子树和右子树
	remainingSum := targetSum - root.Val
	return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}

func main() {
	// 创建一个测试树： [5,4,8,11,null,13,4,7,2,null,null,null,1]
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	targetSum := 22
	result := hasPathSum(root, targetSum)
	fmt.Println(result) // 输出: true
}
