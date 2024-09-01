package main

import "fmt"

func totalNQueens(n int) int {
	var count int
	// 初始化標記列、主對角線、副對角線的陣列
	cols := make([]bool, n)
	d1 := make([]bool, 2*n) // 主對角線
	d2 := make([]bool, 2*n) // 副對角線

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			count++ // 找到一個有效解
			return
		}
		for col := 0; col < n; col++ {
			id1 := col - row + n // 主對角線索引
			id2 := col + row     // 副對角線索引
			if cols[col] || d1[id1] || d2[id2] {
				continue // 如果列或對角線已經有皇后，則跳過
			}
			// 放置皇后
			cols[col], d1[id1], d2[id2] = true, true, true
			// 進入下一行
			backtrack(row + 1)
			// 撤銷選擇
			cols[col], d1[id1], d2[id2] = false, false, false
		}
	}

	backtrack(0)
	return count
}

func main() {
	n := 4
	fmt.Println(totalNQueens(n)) // 輸出: 2
}
