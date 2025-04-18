// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX < m {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY < n {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}
