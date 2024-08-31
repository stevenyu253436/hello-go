package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	// Step 1: Transpose the matrix (swap rows and columns)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func main() {
	// Example usage
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rotate(matrix)

	fmt.Println("Rotated matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
