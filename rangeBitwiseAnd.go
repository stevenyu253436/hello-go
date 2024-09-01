func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	// 找到 left 和 right 的共同前綴
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	// 共同前綴左移回原位
	return left << shift
}
