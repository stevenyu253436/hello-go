package main

import (
	"fmt"
)

// ListNode is a definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the nth node from the end of the list and returns its head.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node which points to the head
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy

	// Move first pointer so that there is a gap of n nodes between first and second
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until first reaches the end of the list
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from the end
	second.Next = second.Next.Next

	// Return the head of the modified list
	return dummy.Next
}

// Helper function to print the list (for testing)
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example usage:
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	// Remove the 2nd node from the end: 1 -> 2 -> 3 -> 5
	head = removeNthFromEnd(head, 2)

	// Print the modified list
	printList(head)
}
