// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package reverseinteger

func reverse(x int) int {
	const (
		maxInt32 = 1<<31 - 1 // 2147483647
		minInt32 = -1 << 31  // -2147483648
	)

	result := 0
	sign := 1

	if x < 0 {
		sign = -1
		x = -x
	}

	for x > 0 {
		remainder := x % 10

		// Check for positive overflow
		if sign == 1 && (result > maxInt32/10 || (result == maxInt32/10 && remainder > 7)) {
			return 0
		}

		// Check for negative overflow
		if sign == -1 && (result > -(minInt32/10) || (result == -(minInt32/10) && remainder > 8)) {
			return 0
		}

		result = result*10 + remainder
		x /= 10
	}

	return sign * result
}
