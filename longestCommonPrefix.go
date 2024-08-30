package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 初始化最長公共前綴為第一個字符串
	prefix := strs[0]

	for _, str := range strs[1:] {
		// 當前字符串與前綴逐字符比較，縮短前綴
		for len(str) < len(prefix) || str[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	// 測試範例
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))           // 輸出: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))              // 輸出: ""
	fmt.Println(longestCommonPrefix([]string{"interview", "inter", "intermediate"})) // 輸出: "inter"
}
