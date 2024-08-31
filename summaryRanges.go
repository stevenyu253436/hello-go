package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		// 如果当前元素与前一个元素不连续，则结束当前范围
		if nums[i] != nums[i-1]+1 {
			// 将范围添加到结果中
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i-1]))
			}
			start = nums[i] // 开始一个新的范围
		}
	}

	// 处理最后一个范围
	if start == nums[len(nums)-1] {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(nums[len(nums)-1]))
	}

	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // ["0->2", "4->5", "7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0", "2->4", "6", "8->9"]
	fmt.Println(summaryRanges([]int{}))                    // []
	fmt.Println(summaryRanges([]int{-1}))                  // ["-1"]
	fmt.Println(summaryRanges([]int{0}))                   // ["0"]
}
