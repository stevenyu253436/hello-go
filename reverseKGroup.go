package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	// Create a dummy node to handle edge cases easily
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		// Check if there are at least k nodes left to reverse
		kth := getKthNode(prevGroupEnd, k)
		if kth == nil {
			break
		}
		nextGroupStart := kth.Next

		// Reverse the group of k nodes
		prev, curr := kth.Next, prevGroupEnd.Next
		for curr != nextGroupStart {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}

		// Connect the reversed group with the previous part of the list
		tmp := prevGroupEnd.Next
		prevGroupEnd.Next = kth
		prevGroupEnd = tmp
	}

	return dummy.Next
}

func getKthNode(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}
	return curr
}

// Helper function to print the linked list (for testing purposes)
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Example 1: head = [1,2,3,4,5], k = 2, Output: [2,1,4,3,5]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	k := 2
	result := reverseKGroup(head, k)
	printList(result) // Output should be: 2 -> 1 -> 4 -> 3 -> 5 -> nil

	// Example 2: head = [1,2,3,4,5], k = 3, Output: [3,2,1,4,5]
	head = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	k = 3
	result = reverseKGroup(head, k)
	printList(result) // Output should be: 3 -> 2 -> 1 -> 4 -> 5 -> nil
}
