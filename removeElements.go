package main

import (
	"fmt"
)

// ListNode 定义单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeElements 移除链表中值等于val的节点
func removeElements(head *ListNode, val int) *ListNode {
	// 创建一个虚拟头节点，并将其Next指向当前head
	dummy := &ListNode{Next: head}
	current := dummy

	// 遍历链表
	for current.Next != nil {
		if current.Next.Val == val {
			// 如果下一个节点的值等于val，则删除它
			current.Next = current.Next.Next
		} else {
			// 否则移动到下一个节点
			current = current.Next
		}
	}

	// 返回实际的头节点
	return dummy.Next
}

// helper function to print linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// main function to test removeElements
func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	fmt.Println("Original list:")
	printList(head)

	val := 6
	head = removeElements(head, val)

	fmt.Printf("List after removing %d:\n", val)
	printList(head)
}
