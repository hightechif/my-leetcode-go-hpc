// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// Calculate current area
		currentHeight := min(height[left], height[right])
		width := right - left
		currentArea := currentHeight * width

		// Update max area
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
