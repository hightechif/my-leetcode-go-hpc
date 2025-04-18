// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package longestpalindromicsubstring

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1
	n := len(s)

	// Manually expand around center without function calls
	for i := 0; i < n; {
		l, r := i, i

		// Skip duplicate characters for odd length
		for r < n-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1 // Skip past duplicates

		// Expand around center
		for l > 0 && r < n-1 && s[l-1] == s[r+1] {
			l--
			r++
		}

		if length := r - l + 1; length > maxLen {
			maxLen = length
			start = l
		}
	}

	return s[start : start+maxLen]
}
