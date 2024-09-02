import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	// Handle the sign
	if numerator == 0 {
		return "0"
	}
	result := ""
	if (numerator < 0) != (denominator < 0) {
		result = "-"
	}

	// Work with positive values for simplicity
	n := abs(numerator)
	d := abs(denominator)

	// Append the integer part
	result += strconv.Itoa(n / d)

	// Handle the remainder part
	remainder := n % d
	if remainder == 0 {
		return result // No fractional part
	}

	result += "."

	// Map to store the positions of remainders
	remainderMap := make(map[int]int)

	// Process the fractional part
	for remainder != 0 {
		// If remainder is already in the map, we found a repeating sequence
		if pos, ok := remainderMap[remainder]; ok {
			result = result[:pos] + "(" + result[pos:] + ")"
			return result
		}

		// Store the current position of the remainder
		remainderMap[remainder] = len(result)

		remainder *= 10
		result += strconv.Itoa(remainder / d)
		remainder %= d
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
