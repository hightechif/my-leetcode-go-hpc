// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package threesum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums) // Sort the array first
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// Skip duplicates for left and right pointers
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++ // Need larger sum
			} else {
				right-- // Need smaller sum
			}
		}
	}
	return res
}
