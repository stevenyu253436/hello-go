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

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val

	// 如果当前节点是叶子节点，返回当前路径的数字
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// 递归地计算左子树和右子树的路径和
	leftSum := dfs(node.Left, currentSum)
	rightSum := dfs(node.Right, currentSum)

	// 返回左右子树的路径和
	return leftSum + rightSum
}

func main() {
	// 创建一个测试树: [4,9,0,5,1]
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 1}

	result := sumNumbers(root)
	fmt.Println(result) // 输出: 1026
}
