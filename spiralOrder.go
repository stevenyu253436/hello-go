package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	result := []int{}

	for top <= bottom && left <= right {
		// Traverse from left to right along the top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom along the right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Make sure we are still within valid bounds before continuing
		if top <= bottom {
			// Traverse from right to left along the bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Make sure we are still within valid bounds before continuing
		if left <= right {
			// Traverse from bottom to top along the left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(spiralOrder(matrix)) // 输出: [1, 2, 3, 6, 9, 8, 7, 4, 5]
}
