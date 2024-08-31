package main

import (
	"fmt"
)

// TreeNode 定義二叉樹的節點
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder 函數實現Zigzag層次遍歷
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	leftToRight := true // 用來控制遍歷方向

	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 根據方向插入節點值
			if leftToRight {
				currentLevel = append(currentLevel, node.Val)
			} else {
				currentLevel = append([]int{node.Val}, currentLevel...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// 將當前層結果添加到結果中，並切換方向
		result = append(result, currentLevel)
		leftToRight = !leftToRight
	}

	return result
}

func main() {
	// 測試範例
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(zigzagLevelOrder(root)) // 輸出: [[3], [20, 9], [15, 7]]
}
