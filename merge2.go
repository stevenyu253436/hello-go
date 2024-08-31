package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 先按起始位置对区间进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用于存储合并后的区间
	var merged [][]int
	current := intervals[0]

	for _, interval := range intervals {
		// 如果当前区间与前一个区间重叠，则合并它们
		if interval[0] <= current[1] {
			current[1] = max(current[1], interval[1])
		} else {
			// 否则，将当前区间添加到结果中，并更新当前区间为新的区间
			merged = append(merged, current)
			current = interval
		}
	}
	// 将最后一个区间添加到结果中
	merged = append(merged, current)

	return merged
}

// 辅助函数，返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // [[1, 6], [8, 10], [15, 18]]

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals)) // [[1, 5]]
}
