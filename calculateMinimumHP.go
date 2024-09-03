package main

import (
	"fmt"
)

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	// Initialize a DP table with dimensions m x n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Start from the bottom-right corner
	dp[m-1][n-1] = max(1, 1-dungeon[m-1][n-1])

	// Fill the last row and last column
	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(1, dp[i+1][n-1]-dungeon[i][n-1])
	}
	for j := n - 2; j >= 0; j-- {
		dp[m-1][j] = max(1, dp[m-1][j+1]-dungeon[m-1][j])
	}

	// Fill the rest of the DP table
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
		}
	}

	// The answer is in the top-left corner
	return dp[0][0]
}

// Utility functions for min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	dungeon := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	fmt.Println(calculateMinimumHP(dungeon)) // Output: 7
}
