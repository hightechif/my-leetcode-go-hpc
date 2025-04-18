// Copyright (c) 2025 Ridhan Fadhilah
// License: MIT (https://opensource.org/licenses/MIT)

package twosum

func twoSum(nums []int, target int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j-i]+nums[j] == target {
				var res = []int{j - i, j}
				return res
			}
		}
	}
	return []int{}
}
