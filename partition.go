package main

import "fmt"

// 定义 ListNode 结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// partition 函数，用于按给定值分割链表
func partition(head *ListNode, x int) *ListNode {
	// 创建两个虚拟头结点
	lesser := &ListNode{}
	greater := &ListNode{}

	lesserHead := lesser
	greaterHead := greater

	// 遍历链表，将节点分到 lesser 和 greater 链表中
	current := head
	for current != nil {
		if current.Val < x {
			lesser.Next = current
			lesser = lesser.Next
		} else {
			greater.Next = current
			greater = greater.Next
		}
		current = current.Next
	}

	// 将 greater 链表的最后一个节点的 Next 置为 nil，防止成环
	greater.Next = nil

	// 将 lesser 链表的末尾与 greater 链表的头部连接
	lesser.Next = greaterHead.Next

	// 返回 lesser 链表的头部
	return lesserHead.Next
}

// 辅助函数，用于创建链表
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// 辅助函数，用于打印链表
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// main 函数，用于测试 partition 函数
func main() {
	// 示例链表 [1,4,3,2,5,2]， x = 3
	head := createList([]int{1, 4, 3, 2, 5, 2})
	fmt.Print("Original List: ")
	printList(head)

	// 调用 partition 函数
	result := partition(head, 3)
	fmt.Print("Partitioned List: ")
	printList(result)
}
