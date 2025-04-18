// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package palindromenumber

func isPalindrome(x int) bool {
	// Early exits for non-palindrome candidates
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	if x < 10 {
		return true // Single-digit numbers are always palindromes
	}

	reversedHalf := 0
	original := x // Preserve original for comparison

	// Reverse only half of the number
	for original > reversedHalf {
		reversedHalf = reversedHalf*10 + original%10
		original /= 10
	}

	// Compare halves (even digits) or ignore middle digit (odd digits)
	return original == reversedHalf || original == reversedHalf/10
}
