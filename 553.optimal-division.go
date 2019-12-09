package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=553 lang=golang
 *
 * [553] Optimal Division
 */

// @lc code=start

// max true, min false
func getMaxOrMin(nums []int, mode bool) (float64, string) {
	n := len(nums)
	if n == 1 {
		return float64(nums[0]), strconv.Itoa(nums[0])
	}
	res := -1.0
	resLs := ""
	resRs := ""
	pos := -1
	for i := 1; i < n; i++ {
		left := nums[0:i]
		right := nums[i:n]
		l, ls := getMaxOrMin(left, mode)
		r, rs := getMaxOrMin(right, !mode)
		if res == -1.0 {
			res = l / r
			resLs = ls
			resRs = rs
			pos = i
		} else if (mode && l/r > res) || (!mode && l/r < res) {
			res = l / r
			resLs = ls
			resRs = rs
			pos = i
		}
	}
	if pos == n-1 {
		return res, resLs + "/" + resRs
	} else {
		return res, resLs + "/(" + resRs + ")"
	}
}

func optimalDivision(nums []int) string {
	_, s := getMaxOrMin(nums, true)
	return s
}

// @lc code=end

func main() {
	nums := []int{1000, 100, 10, 2}
	fmt.Println(optimalDivision(nums))
}
