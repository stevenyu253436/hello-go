func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// 创建 DP 数组
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 填充 DP 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

// 辅助函数，用于求三个数中的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
