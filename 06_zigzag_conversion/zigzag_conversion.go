// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package zigzagconversion

func convert(s string, numRows int) string {
	n := len(s)
	if n <= 1 || numRows == 1 {
		return s
	}

	// Create rows as slices of bytes (more efficient than strings for appending)
	rows := make([][]byte, min(numRows, n))

	currentRow := 0
	goingDown := false

	for i := 0; i < n; i++ {
		rows[currentRow] = append(rows[currentRow], s[i])

		// Change direction if we've hit the top or bottom row
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}

		// Move to next row
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	// Combine all rows
	result := make([]byte, 0, n)
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}
