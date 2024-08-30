package main

import "fmt"

// 函數將羅馬數字轉換為整數
func romanToInt(s string) int {
	// 創建一個映射，將羅馬字符映射到其對應的整數值
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(s)

	for i := 0; i < n; i++ {
		// 如果當前字符比下個字符小，則執行減法操作
		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}

	return result
}

func main() {
	// 測試用例
	fmt.Println(romanToInt("III"))     // 輸出: 3
	fmt.Println(romanToInt("LVIII"))   // 輸出: 58
	fmt.Println(romanToInt("MCMXCIV")) // 輸出: 1994
}
