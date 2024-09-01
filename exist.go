package main

import "fmt"

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])
	var backtrack func(row, col, index int) bool

	backtrack = func(row, col, index int) bool {
		if index == len(word) {
			return true
		}

		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != word[index] {
			return false
		}

		temp := board[row][col]
		board[row][col] = '#' // 標記這個位置已經使用過

		// 上下左右四個方向進行搜索
		found := backtrack(row+1, col, index+1) ||
			backtrack(row-1, col, index+1) ||
			backtrack(row, col+1, index+1) ||
			backtrack(row, col-1, index+1)

		board[row][col] = temp // 恢復這個位置的原始值
		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // 輸出: true
}
