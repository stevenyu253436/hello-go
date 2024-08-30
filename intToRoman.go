package main

import (
	"fmt"
)

func intToRoman(num int) string {
	// 定義羅馬數字和對應的數值
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// 初始化結果
	roman := ""

	// 遍歷每個羅馬數字的值，從最大到最小
	for i := 0; i < len(val); i++ {
		for num >= val[i] { // 當數字大於等於當前值時
			roman += symbols[i] // 添加對應的羅馬數字
			num -= val[i]       // 減去當前的值
		}
	}

	return roman
}

func main() {
	// 測試範例
	fmt.Println(intToRoman(3749)) // 輸出: "MMMDCCXLIX"
	fmt.Println(intToRoman(58))   // 輸出: "LVIII"
	fmt.Println(intToRoman(1994)) // 輸出: "MCMXCIV"
}
