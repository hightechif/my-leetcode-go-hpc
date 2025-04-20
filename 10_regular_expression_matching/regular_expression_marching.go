// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package regularexpressionmatching

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string matches empty pattern
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* etc.
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc, pc := s[i-1], p[j-1]
			if pc == '*' {
				// Case 1: match zero of preceding element
				// Case 2: match one or more of preceding element
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				dp[i][j] = dp[i-1][j-1] && (sc == pc || pc == '.')
			}
		}
	}

	return dp[m][n]
}
