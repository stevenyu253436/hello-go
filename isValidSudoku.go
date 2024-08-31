package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	// 初始化每一行、每一列和每一個 3x3 子格的集合
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	// 遍歷整個數獨板
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			// 檢查行是否有重複
			if rows[i][num] {
				return false
			}
			rows[i][num] = true

			// 檢查列是否有重複
			if cols[j][num] {
				return false
			}
			cols[j][num] = true

			// 檢查 3x3 子格是否有重複
			boxIndex := (i/3)*3 + j/3
			if boxes[boxIndex][num] {
				return false
			}
			boxes[boxIndex][num] = true
		}
	}

	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	result := isValidSudoku(board)
	fmt.Println(result) // 應輸出 true 或 false，根據數獨板的有效性
}
