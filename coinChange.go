import "math"

func coinChange(coins []int, amount int) int {
	// 初始化 dp 數組，長度為 amount + 1，所有元素設為最大值
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	// 設置 dp[0] = 0，表示金額為 0 所需硬幣數為 0
	dp[0] = 0

	// 遍歷金額，從 1 到 amount
	for i := 1; i <= amount; i++ {
		// 遍歷所有硬幣
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// 如果 dp[amount] 仍然是最大值，表示無法湊出這個金額
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
