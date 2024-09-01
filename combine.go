package main

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	var result [][]int
	var combination []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(combination) == k {
			// 當前組合達到 k 個元素，將其加入結果集中
			combinationCopy := make([]int, len(combination))
			copy(combinationCopy, combination)
			result = append(result, combinationCopy)
			return
		}

		for i := start; i <= n; i++ {
			combination = append(combination, i)
			backtrack(i + 1)                               // 遞歸生成後續的組合
			combination = combination[:len(combination)-1] // 回溯
		}
	}

	backtrack(1)
	return result
}

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k)) // 输出: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
}
