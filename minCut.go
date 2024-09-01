func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Step 1: 預計算每個子字符串是否為回文
	palindrome := make([][]bool, n)
	for i := range palindrome {
		palindrome[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || palindrome[i+1][j-1]) {
				palindrome[i][j] = true
			}
		}
	}

	// Step 2: 使用 DP 計算最小分割次數
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if palindrome[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i // 最壞情況是每個字元都要切一刀
			for j := 0; j < i; j++ {
				if palindrome[j+1][i] {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
