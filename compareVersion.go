func compareVersion(version1 string, version2 string) int {
	// Split the version strings into slices of strings
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	// Find the maximum length between the two version slices
	maxLen := max(len(v1), len(v2))

	for i := 0; i < maxLen; i++ {
		// Convert the current revision to integer for both versions
		rev1 := 0
		rev2 := 0

		if i < len(v1) {
			rev1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			rev2, _ = strconv.Atoi(v2[i])
		}

		// Compare the revisions
		if rev1 < rev2 {
			return -1
		} else if rev1 > rev2 {
			return 1
		}
	}

	// If all revisions are equal, return 0
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}