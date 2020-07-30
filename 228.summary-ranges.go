package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func getRange(l, r int) string {
	if r == -99999 {
		return strconv.Itoa(l)
	}
	return strconv.Itoa(l) + "->" + strconv.Itoa(r)

}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	l := nums[0]
	r := -99999
	res := []string{}
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		if x == l+1 {
			r = x
		} else if x == r+1 {
			r = x
		} else {
			res = append(res, getRange(l, r))
			l = x
			r = -99999
		}
	}
	res = append(res, getRange(l, r))
	return res
}

// @lc code=end

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
