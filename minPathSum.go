func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// 初始化第一行和第一列的最小路徑和
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	// 計算其餘位置的最小路徑和
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	// 返回右下角的值，即為最小路徑和
	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
