package main

import (
	"fmt"
)

/*
 * @lc app=leetcode id=795 lang=golang
 *
 * [795] Number of Subarrays with Bounded Maximum
 */

// @lc code=start
// 0 -2 3

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func numSubarrayBoundedMax(A []int, L int, R int) int {

	idx := make([]int, 1)
	idx[0] = 0
	for i := 0; i < len(A); i++ {
		if A[i] > R {
			idx = append(idx, -i-1)
		} else if A[i] >= L {
			idx = append(idx, i+1)
		}
	}
	idx = append(idx, -len(A)-1)

	ans := 0
	lastMaxpos := 0
	lastGoodPos := 0

	for i := 0; i < len(idx); i++ {
		pos := idx[i]
		if lastGoodPos > 0 {
			ans += (lastGoodPos - lastMaxpos) * (abs(pos) - lastGoodPos - 1)
		}

		if pos <= 0 {
			lastMaxpos = -pos
			lastGoodPos = 0
		} else {
			ans += pos - lastMaxpos
			lastGoodPos = pos
		}
	}

	return ans
}

// @lc code=end
func main() {
	a := []int{45, 2, 49, 6, 55, 73, 27, 17, 4, 71}
	l := 17
	r := 33
	fmt.Println(numSubarrayBoundedMax(a, l, r))
}
