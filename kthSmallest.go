package main

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest 返回 BST 中的第 k 小的元素
func kthSmallest(root *TreeNode, k int) int {
	var inorder func(node *TreeNode)
	count := 0
	result := 0

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中序遍历左子树
		inorder(node.Left)

		// 访问当前节点
		count++
		if count == k {
			result = node.Val
			return
		}

		// 中序遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return result
}

// 主函数用于测试
func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}

	k := 1
	result := kthSmallest(root, k)
	println("The", k, "th smallest element is:", result) // 输出: 1
}
