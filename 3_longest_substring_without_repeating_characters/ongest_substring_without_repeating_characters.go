// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	// Use byte instead of rune since we only handle ASCII
	lastSeen := make([]int, 128) // ASCII covers 0-127
	for i := range lastSeen {
		lastSeen[i] = -1 // Initialize with -1 (not seen)
	}

	maxLen, start := 0, 0

	for end := 0; end < len(s); end++ {
		b := s[end] // Get ASCII byte
		if idx := lastSeen[b]; idx >= start {
			start = idx + 1
		}
		lastSeen[b] = end
		if currentLen := end - start + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
