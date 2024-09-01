func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // 若最右邊的位元是1，則計數器加1
		n >>= 1        // 右移n，檢查下一個位元
	}
	return count
}
