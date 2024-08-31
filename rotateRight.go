package main

import (
	"fmt"
)

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// rotateRight 函数用于旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表长度
	length := 1
	current := head
	for current.Next != nil {
		current = current.Next
		length++
	}

	// 将链表末尾连接到头部，形成一个环
	current.Next = head

	// 计算新的头节点位置
	k = k % length
	stepsToNewHead := length - k

	// 找到新的头节点
	for stepsToNewHead > 0 {
		current = current.Next
		stepsToNewHead--
	}

	// 断开环
	newHead := current.Next
	current.Next = nil

	return newHead
}

// printList 用于打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}

// main 函数为程序的入口
func main() {
	// 创建一个测试用的链表: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original list:")
	printList(head)

	// 旋转链表
	k := 2
	rotatedHead := rotateRight(head, k)

	fmt.Println("Rotated list:")
	printList(rotatedHead)
}
