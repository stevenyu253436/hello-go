import (
	"container/heap"
)

// ListNode 是链表节点的定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 定义一个实现了 heap.Interface 的最小堆
type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// mergeKLists 函数合并 k 个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// 将每个链表的第一个节点加入最小堆
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	// 创建一个虚拟头节点以便返回结果链表
	dummy := &ListNode{}
	current := dummy

	// 从最小堆中不断取出最小元素，并将其下一个节点（如果有的话）加入最小堆
	for minHeap.Len() > 0 {
		minNode := heap.Pop(minHeap).(*ListNode)
		current.Next = minNode
		current = current.Next
		if minNode.Next != nil {
			heap.Push(minHeap, minNode.Next)
		}
	}

	return dummy.Next
}
