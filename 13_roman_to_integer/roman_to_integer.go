// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package romantointeger

func romanToInt(s string) int {
	// Direct ASCII values (no subtraction math)
	var values = [256]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := values[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if current := values[s[i]]; current < values[s[i+1]] {
			sum -= current
		} else {
			sum += current
		}
	}
	return sum
}
