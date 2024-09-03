package main

import (
	"fmt"
)

func titleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		result = result*26 + int(columnTitle[i]-'A'+1)
	}
	return result
}

func main() {
	fmt.Println(titleToNumber("A"))  // Output: 1
	fmt.Println(titleToNumber("AB")) // Output: 28
	fmt.Println(titleToNumber("ZY")) // Output: 701
}
