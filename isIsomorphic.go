package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	// 创建两个哈希表用于字符映射
	mapST := make(map[byte]byte)
	mapTS := make(map[byte]byte)

	// 遍历字符串
	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]

		// 检查s到t的映射
		if val, ok := mapST[charS]; ok {
			if val != charT {
				return false
			}
		} else {
			mapST[charS] = charT
		}

		// 检查t到s的映射
		if val, ok := mapTS[charT]; ok {
			if val != charS {
				return false
			}
		} else {
			mapTS[charT] = charS
		}
	}

	return true
}

func main() {
	s1 := "egg"
	t1 := "add"
	fmt.Println(isIsomorphic(s1, t1)) // 输出: true

	s2 := "foo"
	t2 := "bar"
	fmt.Println(isIsomorphic(s2, t2)) // 输出: false

	s3 := "paper"
	t3 := "title"
	fmt.Println(isIsomorphic(s3, t3)) // 输出: true
}
