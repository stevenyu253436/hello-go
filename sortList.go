/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// sortList sorts the linked list using merge sort algorithm
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Step 1: Split the linked list into two halves
	mid := findMiddle(head)
	left := head
	right := mid.Next
	mid.Next = nil

	// Step 2: Recursively sort the two halves
	left = sortList(left)
	right = sortList(right)

	// Step 3: Merge the two sorted halves
	return mergeTwoLists(left, right)
}

// findMiddle finds the middle of the linked list using the slow and fast pointer technique
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// mergeTwoLists merges two sorted linked lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}
