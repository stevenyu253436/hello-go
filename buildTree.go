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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 前序遍历的第一个元素是当前树的根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 使用根节点的位置将中序遍历分为左子树和右子树的部分
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	// 前序遍历中，紧跟根节点的部分属于左子树，之后是右子树
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	// 递归构造左右子树
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	fmt.Println(root.Val)       // 输出: 3
	fmt.Println(root.Left.Val)  // 输出: 9
	fmt.Println(root.Right.Val) // 输出: 20
}
