package main

import (
	"fmt"
)

func isHappy(n int) bool {
	// 定義一個輔助函數來計算數字的每位數字平方和
	getNext := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	slow := n
	fast := getNext(n)

	// 使用快慢指針法來檢測循環
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}

	return fast == 1
}

func main() {
	// 測試範例
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(2))  // false
}
