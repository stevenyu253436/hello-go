func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]  // 获取栈顶元素
		stack = stack[:len(stack)-1] // 弹出栈顶元素

		result = append(result, node.Val)

		// 注意：先将右子节点入栈，确保左子节点后出栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
