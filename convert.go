package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 初始化 numRows 個字串來存儲每一行的結果
	rows := make([]string, numRows)
	currRow := 0
	goingDown := false

	// 模擬字符串在 Z 字形中的排列
	for _, char := range s {
		rows[currRow] += string(char)
		// 當前行數是否該轉方向（向下或向上）
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}
		// 根據方向決定下一個字符應該放在哪一行
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	// 將所有行的字符串拼接成結果
	return strings.Join(rows, "")
}

func main() {
	// 測試範例
	fmt.Println(convert("PAYPALISHIRING", 3)) // Output: "PAHNAPLSIIGYIR"
	fmt.Println(convert("PAYPALISHIRING", 4)) // Output: "PINALSIGYAHRPI"
	fmt.Println(convert("A", 1))              // Output: "A"
}
