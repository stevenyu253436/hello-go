func singleNumber(nums []int) int {
	ones, twos := 0, 0

	for _, num := range nums {
		// `twos`记录的是出现两次的位
		twos |= ones & num

		// `ones`记录的是出现一次的位
		ones ^= num

		// 清除那些出现三次的位
		commonMask := ^(ones & twos)
		ones &= commonMask
		twos &= commonMask
	}

	return ones
}
