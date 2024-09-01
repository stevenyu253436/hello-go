func partitionString(s string) int {
	seen := make(map[rune]bool)
	count := 1

	for _, char := range s {
		if seen[char] {
			count++
			seen = make(map[rune]bool)
		}
		seen[char] = true
	}

	return count
}
