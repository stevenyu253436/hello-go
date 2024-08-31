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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 后序遍历的最后一个元素是当前树的根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	// 划分左子树和右子树的中序和后序遍历序列
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]

	// 递归构造左右子树
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	fmt.Println(root.Val)       // 输出: 3
	fmt.Println(root.Left.Val)  // 输出: 9
	fmt.Println(root.Right.Val) // 输出: 20
}
