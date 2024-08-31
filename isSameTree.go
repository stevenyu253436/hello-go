package main

import "fmt"

// TreeNode 定义二叉树的节点结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree 判断两个二叉树是否相同
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 如果两个节点都为空，则它们相同
	if p == nil && q == nil {
		return true
	}
	// 如果只有一个节点为空，则它们不同
	if p == nil || q == nil {
		return false
	}
	// 如果节点值不同，则它们不同
	if p.Val != q.Val {
		return false
	}
	// 递归检查左子树和右子树是否相同
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// main 函数是程序的入口点
func main() {
	// 构造示例二叉树
	tree1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	// 检查两个二叉树是否相同
	result := isSameTree(tree1, tree2)

	// 输出结果
	fmt.Println("Are the two trees the same?", result)
}
