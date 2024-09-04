func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// 計算矩形1的面積
	area1 := (ax2 - ax1) * (ay2 - ay1)

	// 計算矩形2的面積
	area2 := (bx2 - bx1) * (by2 - by1)

	// 計算重疊區域的邊界
	overlapWidth := max(0, min(ax2, bx2)-max(ax1, bx1))
	overlapHeight := max(0, min(ay2, by2)-max(ay1, by1))

	// 計算重疊區域的面積
	overlapArea := overlapWidth * overlapHeight

	// 計算總面積
	return area1 + area2 - overlapArea
}

// 求兩個整數中的較大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 求兩個整數中的較小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}