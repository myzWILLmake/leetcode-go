package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */
func largestDivisibleSubset(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	sort.Ints(nums)
	f := make([][]int, n)
	f[0] = []int{nums[0]}

	res := f[0]
	for i := 1; i < n; i++ {
		max := []int{}
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if len(f[j]) > len(max) {
					max = f[j]
				}
			}
		}
		f[i] = make([]int, len(max)+1)
		copy(f[i], max)
		f[i][len(max)] = nums[i]

		if len(f[i]) > len(res) {
			res = f[i]
		}
	}
	return res
}

// func main() {
// 	nums := []int{1, 2, 3, 4, 8, 16, 9, 18}
// 	fmt.Println(largestDivisibleSubset(nums))
// }
