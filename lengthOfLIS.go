import "sort"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := []int{}

	for _, num := range nums {
		i := sort.Search(len(dp), func(i int) bool {
			return dp[i] >= num
		})
		if i < len(dp) {
			dp[i] = num
		} else {
			dp = append(dp, num)
		}
	}

	return len(dp)
}
