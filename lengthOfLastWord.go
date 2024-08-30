package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	// 初始化最後一個單詞的長度
	length := 0

	// 從字符串的末尾開始遍歷
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' { // 當前字符不是空格時
			length++
		} else if length > 0 { // 當前字符是空格並且已經找到最後一個單詞時
			break
		}
	}

	return length
}

func main() {
	// 測試範例
	fmt.Println(lengthOfLastWord("Hello World"))                 // 輸出: 5
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  ")) // 輸出: 4
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))       // 輸出: 6
}
