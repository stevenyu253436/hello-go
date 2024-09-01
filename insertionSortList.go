/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	// 哨兵节点，帮助处理边界情况
	dummy := &ListNode{Next: head}
	current := head

	for current != nil && current.Next != nil {
		// 如果当前节点已经有序，继续下一个
		if current.Val <= current.Next.Val {
			current = current.Next
			continue
		}

		// 需要插入的节点
		toInsert := current.Next
		current.Next = toInsert.Next // 将 toInsert 从原位置移除

		// 从 dummy 开始找到插入位置
		prev := dummy
		for prev.Next.Val < toInsert.Val {
			prev = prev.Next
		}

		// 将 toInsert 插入到 prev 和 prev.Next 之间
		toInsert.Next = prev.Next
		prev.Next = toInsert
	}

	return dummy.Next
}
