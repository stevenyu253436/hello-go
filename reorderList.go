func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Use two pointers to find the middle of the list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the list
	var prev *ListNode
	curr := slow
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	// Step 3: Merge the two halves
	first, second := head, prev
	for second.Next != nil {
		temp1 := first.Next
		temp2 := second.Next

		first.Next = second
		second.Next = temp1

		first = temp1
		second = temp2
	}
}
