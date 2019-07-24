package main

import (
	"math"
)

/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	n := len(nums)
	if n <= 1 || k == 0 {
		return false
	}

	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && j >= i-k; j-- {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) {
				return true
			}
		}
	}

	return false
}

// func main() {
// 	nums := []int{1, 5, 9, 1, 5, 9}
// 	k := 2
// 	t := 3
// 	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
// }
