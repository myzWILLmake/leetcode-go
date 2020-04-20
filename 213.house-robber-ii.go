package main

import "fmt"

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func big(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	r := make([]int, n)
	nr := make([]int, n)
	r[0] = nums[0]
	ans := r[0]
	for i := 1; i < n; i++ {
		r[i] = nr[i-1] + nums[i]
		nr[i] = big(r[i-1], nr[i-1])
		if r[i] > ans {
			ans = r[i]
		}
	}
	return ans
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return big(_rob(nums[1:]), _rob(nums[:n-1]))
}

// @lc code=end

func main() {
	nums := []int{2, 3, 2}
	fmt.Println(rob(nums))
}
