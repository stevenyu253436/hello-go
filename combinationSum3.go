func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var combination []int
	backtrack(k, n, 1, combination, &result)
	return result
}

func backtrack(k int, target int, start int, combination []int, result *[][]int) {
	// 终止条件：当剩余的数字个数为0时
	if k == 0 {
		if target == 0 {
			// 如果组合的和等于目标值，将组合添加到结果集
			temp := make([]int, len(combination))
			copy(temp, combination)
			*result = append(*result, temp)
		}
		return
	}

	// 递归构建组合
	for i := start; i <= 9; i++ {
		if i > target { // 如果当前数字大于目标和，终止循环
			break
		}
		combination = append(combination, i)
		backtrack(k-1, target-i, i+1, combination, result) // 递归调用
		combination = combination[:len(combination)-1]     // 回溯
	}
}