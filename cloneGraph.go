package main

import "fmt"

// Node 定義圖的節點結構
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph 是主要函數，用來克隆圖
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 用來記錄已克隆的節點
	visited := make(map[*Node]*Node)

	// 定義一個遞歸的DFS函數
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}

		// 如果該節點已被克隆，直接返回克隆的節點
		if clone, found := visited[n]; found {
			return clone
		}

		// 創建克隆的節點
		cloneNode := &Node{Val: n.Val}
		visited[n] = cloneNode

		// 克隆所有的鄰居節點
		for _, neighbor := range n.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(neighbor))
		}

		return cloneNode
	}

	// 從輸入節點開始克隆
	return dfs(node)
}

func main() {
	// 測試範例
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	clone := cloneGraph(node1)

	// 檢查克隆的結果
	fmt.Printf("克隆的節點值: %d\n", clone.Val)
	for _, neighbor := range clone.Neighbors {
		fmt.Printf("鄰居節點值: %d\n", neighbor.Val)
	}
}
