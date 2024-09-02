import (
	"math"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	minNum, maxNum := math.MaxInt32, math.MinInt32
	for _, num := range nums {
		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}
	}

	// Calculate the minimum possible gap (bucket size)
	bucketSize := int(math.Max(1, float64(maxNum-minNum)/float64(len(nums)-1)))
	bucketCount := (maxNum-minNum)/bucketSize + 1

	// Initialize buckets
	bucketsMin := make([]int, bucketCount)
	bucketsMax := make([]int, bucketCount)
	for i := range bucketsMin {
		bucketsMin[i] = math.MaxInt32
		bucketsMax[i] = math.MinInt32
	}

	// Place numbers into buckets
	for _, num := range nums {
		idx := (num - minNum) / bucketSize
		if num < bucketsMin[idx] {
			bucketsMin[idx] = num
		}
		if num > bucketsMax[idx] {
			bucketsMax[idx] = num
		}
	}

	// Find the maximum gap
	maxGap := 0
	previousMax := minNum

	for i := 0; i < bucketCount; i++ {
		// Skip empty buckets
		if bucketsMin[i] == math.MaxInt32 && bucketsMax[i] == math.MinInt32 {
			continue
		}
		// Gap between the current min bucket and the previous max bucket
		maxGap = int(math.Max(float64(maxGap), float64(bucketsMin[i]-previousMax)))
		previousMax = bucketsMax[i]
	}

	return maxGap
}
