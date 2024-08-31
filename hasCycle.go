package main

import (
	"fmt"
)

// ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle detects if a linked list has a cycle.
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}

// Helper function to create a linked list from a slice.
func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to create a cycle in the linked list.
func createCycle(head *ListNode, pos int) {
	if head == nil {
		return
	}
	tail := head
	var cycleNode *ListNode
	for i := 0; tail.Next != nil; i++ {
		if i == pos {
			cycleNode = tail
		}
		tail = tail.Next
	}
	tail.Next = cycleNode
}

func main() {
	// Create a linked list [3, 2, 0, -4]
	vals := []int{3, 2, 0, -4}
	head := createLinkedList(vals)

	// Create a cycle where the tail connects to the 1st node (index 1)
	createCycle(head, 1)

	// Test if the linked list has a cycle
	if hasCycle(head) {
		fmt.Println("Cycle detected")
	} else {
		fmt.Println("No cycle detected")
	}
}
