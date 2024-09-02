func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber-- // Adjust the number to 0-based
		remainder := columnNumber % 26
		result = string('A'+remainder) + result
		columnNumber /= 26
	}

	return result
}
