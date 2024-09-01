package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			// 當前排列完成，將其加入結果集中
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			// 選擇 nums[i]
			current = append(current, nums[i])
			used[i] = true
			// 繼續生成剩下的排列
			backtrack()
			// 撤銷選擇
			current = current[:len(current)-1]
			used[i] = false
		}
	}

	backtrack()
	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums)) // 输出: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
