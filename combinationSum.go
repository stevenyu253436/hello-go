package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	var backtrack func(target int, start int)
	backtrack = func(target int, start int) {
		if target == 0 {
			// 當前組合的總和等於目標，將其加入結果集中
			comb := make([]int, len(current))
			copy(comb, current)
			result = append(result, comb)
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] <= target {
				// 選擇 candidates[i]
				current = append(current, candidates[i])
				// 遞歸呼叫，因為數字可以重複使用，所以參數 start 依然是 i
				backtrack(target-candidates[i], i)
				// 撤銷選擇
				current = current[:len(current)-1]
			}
		}
	}

	backtrack(target, 0)
	return result
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target)) // 输出: [[2,2,3],[7]]
}
