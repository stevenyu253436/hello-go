package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	// 创建一个虚拟头节点
	dummy := &ListNode{Next: head}
	prev := dummy

	// 当前节点指针
	current := head

	for current != nil {
		// 检查当前节点的值是否与下一个节点相同
		if current.Next != nil && current.Val == current.Next.Val {
			// 跳过所有具有相同值的节点
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			// 将prev的Next指向current的Next，跳过重复节点
			prev.Next = current.Next
		} else {
			// 没有重复的情况，移动prev指针
			prev = prev.Next
		}
		// 移动当前节点指针
		current = current.Next
	}

	return dummy.Next
}

// 用于创建链表的辅助函数
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

// 用于打印链表的辅助函数
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// 测试用例
	list := createList([]int{1, 2, 3, 3, 4, 4, 5})
	fmt.Println("原始链表:")
	printList(list)

	// 删除重复节点
	list = deleteDuplicates(list)
	fmt.Println("删除重复节点后的链表:")
	printList(list)
}
