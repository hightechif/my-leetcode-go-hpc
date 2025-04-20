// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	remaining := strs[1:]

	for i := 0; i < len(first); i++ {
		b := first[i] // byte comparison
		for _, s := range remaining {
			if i >= len(s) || s[i] != b {
				return first[:i]
			}
		}
	}
	return first
}
