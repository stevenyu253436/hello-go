package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0

	// 将所有在 newInterval 之前的区间添加到结果中
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合并与 newInterval 重叠的区间
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	result = append(result, newInterval)

	// 将剩余的区间添加到结果中
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals1 := [][]int{{1, 3}, {6, 9}}
	newInterval1 := []int{2, 5}
	fmt.Println(insert(intervals1, newInterval1)) // [[1,5],[6,9]]

	intervals2 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval2 := []int{4, 8}
	fmt.Println(insert(intervals2, newInterval2)) // [[1,2],[3,10],[12,16]]
}
