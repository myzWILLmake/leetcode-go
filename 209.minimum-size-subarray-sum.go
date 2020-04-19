package main

import "fmt"

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	ans := 0

	sum = nums[0]
	if sum >= s {
		ans = 1
	}

	for i, j := 0, 1; i < len(nums) && j < len(nums); j++ {
		sum += nums[j]

		for i < len(nums) && sum-nums[i] >= s {
			sum -= nums[i]
			i++
		}

		if sum >= s {
			if ans == 0 || j-i+1 < ans {
				ans = j - i + 1
			}
		}

	}
	return ans
}

// @lc code=end

func main() {
	s := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
