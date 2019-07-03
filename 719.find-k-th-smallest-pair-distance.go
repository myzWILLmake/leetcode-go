package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=719 lang=golang
 *
 * [719] Find K-th Smallest Pair Distance
 */
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	n := len(nums)
	lo := 0
	hi := nums[n-1] - nums[0]

	for lo < hi {
		mi := (lo + hi) / 2
		count := 0
		left := 0
		for right := 0; right < n; right++ {
			for nums[right]-nums[left] > mi {
				left++
			}
			count += right - left
		}
		if count >= k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}

	return lo
}

// func main() {
// 	nums := []int{1, 3, 4, 5, 6, 7}
// 	k := 6
// 	fmt.Println(smallestDistancePair(nums, k))
// }
