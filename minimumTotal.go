func minimumTotal(triangle [][]int) int {
	// 如果三角形是空的，直接返回0
	if len(triangle) == 0 {
		return 0
	}

	// 建立一個一維dp數組，初始值設置為最底層的元素
	dp := make([]int, len(triangle))
	copy(dp, triangle[len(triangle)-1])

	// 從倒數第二層開始往上計算
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// dp[j] 表示從當前點向下移動的最小路徑和
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	// dp[0] 最終將包含從頂部到底部的最小路徑和
	return dp[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
