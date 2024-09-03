func shortestPalindrome(s string) string {
	// 生成新的字符串
	revS := reverse(s)
	newS := s + "#" + revS

	// 构造KMP的部分匹配表
	lps := make([]int, len(newS))
	for i := 1; i < len(newS); i++ {
		j := lps[i-1]
		for j > 0 && newS[i] != newS[j] {
			j = lps[j-1]
		}
		if newS[i] == newS[j] {
			j++
		}
		lps[i] = j
	}

	// 计算出要添加的部分
	return revS[:len(s)-lps[len(newS)-1]] + s
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
