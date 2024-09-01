func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	stack := []*TreeNode{}
	var lastVisited *TreeNode
	current := root

	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			peekNode := stack[len(stack)-1]
			if peekNode.Right != nil && lastVisited != peekNode.Right {
				current = peekNode.Right
			} else {
				result = append(result, peekNode.Val)
				lastVisited = peekNode
				stack = stack[:len(stack)-1]
			}
		}
	}

	return result
}
