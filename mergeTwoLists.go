package main

import (
	"fmt"
)

// ListNode 是链表节点的定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists 函数合并两个已排序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个虚拟头节点，以方便处理边界情况
	dummy := &ListNode{}
	current := dummy

	// 遍历两个链表，直到一个链表为空
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 如果其中一个链表还有剩余，直接连接到当前节点后
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// 返回合并后的链表的头节点
	return dummy.Next
}

// 打印链表
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}
		node = node.Next
	}
	fmt.Println()
}

func main() {
	// 示例 1
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList) // 输出: 1 -> 1 -> 2 -> 3 -> 4 -> 4

	// 示例 2
	list1 = nil
	list2 = nil
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList) // 输出: (空链表)

	// 示例 3
	list1 = nil
	list2 = &ListNode{0, nil}
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList) // 输出: 0
}
