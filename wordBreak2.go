func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	memo := make(map[string][]string)

	return helper(s, dict, memo)
}

func helper(s string, dict map[string]bool, memo map[string][]string) []string {
	if res, exists := memo[s]; exists {
		return res
	}

	if len(s) == 0 {
		return []string{""}
	}

	var res []string
	for i := 1; i <= len(s); i++ {
		prefix := s[:i]
		if dict[prefix] {
			suffix := s[i:]
			suffixBreaks := helper(suffix, dict, memo)
			for _, segment := range suffixBreaks {
				if segment == "" {
					res = append(res, prefix)
				} else {
					res = append(res, prefix+" "+segment)
				}
			}
		}
	}

	memo[s] = res
	return res
}
