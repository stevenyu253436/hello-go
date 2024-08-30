package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Step 1: Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Step 2: Split the string by spaces into words
	words := strings.Fields(s)

	// Step 3: Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Step 4: Join the words into a single string with a single space separator
	return strings.Join(words, " ")
}

func main() {
	// Test cases
	fmt.Println(reverseWords("the sky is blue"))            // Output: "blue is sky the"
	fmt.Println(reverseWords("  hello world  "))            // Output: "world hello"
	fmt.Println(reverseWords("a good   example"))           // Output: "example good a"
	fmt.Println(reverseWords("  multiple   spaces   here")) // Output: "here spaces multiple"
}
