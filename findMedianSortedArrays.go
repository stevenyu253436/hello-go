package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	imin, imax, halfLen := 0, m, (m+n+1)/2
	var maxOfLeft, minOfRight float64

	for imin <= imax {
		i := (imin + imax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			// i 太小，必须向右移动
			imin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// i 太大，必须向左移动
			imax = i - 1
		} else {
			// 找到正确的 i
			if i == 0 {
				maxOfLeft = float64(nums2[j-1])
			} else if j == 0 {
				maxOfLeft = float64(nums1[i-1])
			} else {
				maxOfLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			if (m+n)%2 == 1 {
				return maxOfLeft
			}

			if i == m {
				minOfRight = float64(nums2[j])
			} else if j == n {
				minOfRight = float64(nums1[i])
			} else {
				minOfRight = math.Min(float64(nums1[i]), float64(nums2[j]))
			}

			return (maxOfLeft + minOfRight) / 2.0
		}
	}

	return 0.0
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.00000

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2)) // Output: 2.50000
}
