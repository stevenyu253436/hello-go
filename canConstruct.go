package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// 创建一个map来存储magazine中每个字母的出现次数
	charCount := make(map[rune]int)

	// 遍历magazine，统计每个字母的出现次数
	for _, char := range magazine {
		charCount[char]++
	}

	// 遍历ransomNote，检查是否能构造出ransomNote
	for _, char := range ransomNote {
		if charCount[char] > 0 {
			charCount[char]-- // 如果有对应字母，减少计数
		} else {
			return false // 如果没有足够的字母，返回false
		}
	}

	return true // 所有字母都满足条件，返回true
}

func main() {
	ransomNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomNote, magazine)) // 输出 true
}
