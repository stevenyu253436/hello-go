package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[rune]int)
    maxLength := 0
    left := 0

    for right, char := range s {
        if index, found := charIndex[char]; found && index >= left {
            // Move the left pointer to the right of the last occurrence of the current character
            left = index + 1
        }
        // Update the character's last occurrence
        charIndex[char] = right
        // Calculate the max length
        currentLength := right - left + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
    fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
    fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
}
