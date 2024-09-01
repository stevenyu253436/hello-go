func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		// 将 result 左移一位，为当前位腾出位置
		result <<= 1
		// 将 num 的最低有效位加到 result
		result |= num & 1
		// 将 num 右移一位，为下一次处理准备
		num >>= 1
	}
	return result
}
