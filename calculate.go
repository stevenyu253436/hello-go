package main

import (
	"fmt"
)

func calculate(s string) int {
	stack := []int{}
	result := 0
	sign := 1
	n := len(s)

	for i := 0; i < n; i++ {
		char := s[i]

		if char >= '0' && char <= '9' {
			num := 0
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			result += sign * num
			i-- // 調整索引以考慮外部for循環的增量
		} else if char == '+' {
			sign = 1
		} else if char == '-' {
			sign = -1
		} else if char == '(' {
			// 把當前的結果和符號推入棧中
			stack = append(stack, result)
			stack = append(stack, sign)
			// 重置結果和符號
			result = 0
			sign = 1
		} else if char == ')' {
			// 從棧中彈出符號並與當前結果相乘
			result *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 與棧中的結果相加
			result += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}

func main() {
	fmt.Println(calculate("1 + 1"))               // Output: 2
	fmt.Println(calculate(" 2-1 + 2 "))           // Output: 3
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)")) // Output: 23
}
