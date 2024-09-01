package main

import (
	"fmt"
)

func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}

	maxPointsOnLine := 0

	for i := 0; i < len(points); i++ {
		slopes := make(map[string]int)
		overlap := 0
		currentMax := 0

		for j := i + 1; j < len(points); j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			if dx == 0 && dy == 0 {
				overlap++
				continue
			}

			gcd := gcd(dx, dy)
			dx /= gcd
			dy /= gcd

			slope := fmt.Sprintf("%d_%d", dx, dy)
			slopes[slope]++
			currentMax = max(currentMax, slopes[slope])
		}

		maxPointsOnLine = max(maxPointsOnLine, currentMax+overlap+1)
	}

	return maxPointsOnLine
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	points := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}
	fmt.Println(maxPoints(points)) // Output: 4
}
