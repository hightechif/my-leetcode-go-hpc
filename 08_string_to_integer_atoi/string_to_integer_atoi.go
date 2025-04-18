// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package stringtointegeratoi

func myAtoi(s string) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	index := 0
	n := len(s)

	// Skip leading whitespace
	for index < n && s[index] == ' ' {
		index++
	}

	// Handle empty string after whitespace
	if index == n {
		return 0
	}

	// Determine sign
	sign := 1
	if s[index] == '-' || s[index] == '+' {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	result := 0
	for index < n && isDigit(s[index]) {
		digit := int(s[index] - '0')

		// Check for overflow
		if sign == 1 && result > (MaxInt32-digit)/10 {
			return MaxInt32
		}
		if sign == -1 && -result < (MinInt32+digit)/10 {
			return MinInt32
		}
		// Special case: -2147483648
		if sign == -1 && result*10+digit == -MinInt32 {
			return MinInt32
		}

		result = result*10 + digit
		index++
	}

	return sign * result
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
