package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用 map 作為集合來存儲數字
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longestStreak := 0

	// 遍歷數組中的每個數字
	for num := range numSet {
		// 如果該數字是序列的起點（num-1不在集合中）
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentStreak := 1

			// 依次檢查連續的數字
			for {
				_, exists := numSet[currentNum+1]
				if !exists {
					break
				}
				currentNum++
				currentStreak++
			}

			// 更新最長連續序列的長度
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         // 4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
}
