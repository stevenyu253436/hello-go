package main

import "fmt"

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowZero := false // 用于标记第一行是否有零
	colZero := false // 用于标记第一列是否有零

	// 检查第一行是否有零
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			rowZero = true
			break
		}
	}

	// 检查第一列是否有零
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			colZero = true
			break
		}
	}

	// 用第一行和第一列作为标记
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 根据标记设置对应行列为零
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 如果第一行有零，则将第一行清零
	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 如果第一列有零，则将第一列清零
	if colZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix) // 输出：[[1 0 1] [0 0 0] [1 0 1]]

	matrix2 := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix2)
	fmt.Println(matrix2) // 输出：[[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}
