package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 根據氣球的結束位置進行排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	currentEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > currentEnd {
			arrows++
			currentEnd = points[i][1]
		}
	}

	return arrows
}

func main() {
	points1 := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points1)) // Output: 2

	points2 := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(findMinArrowShots(points2)) // Output: 4

	points3 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(findMinArrowShots(points3)) // Output: 2
}
